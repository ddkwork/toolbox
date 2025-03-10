// Copyright ©2016-2023 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Package web provides a web server with some standardized logging and
// handler wrapping.
package web

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/ddkwork/toolbox/atexit"
	"github.com/ddkwork/toolbox/errs"
	"github.com/ddkwork/toolbox/xio/network"
	"github.com/ddkwork/toolbox/xio/network/xhttp"
)

// Constants for protocols the server can provide.
const (
	ProtocolHTTP  = "http"
	ProtocolHTTPS = "https"
)

// Server holds the data necessary for the server.
type Server struct {
	CertFile            string
	KeyFile             string
	ShutdownGracePeriod time.Duration
	Logger              *slog.Logger
	WebServer           *http.Server
	Ports               []int
	ShutdownCallback    func()
	StartedChan         chan any // If not nil, will be closed once the server is ready to accept connections
	addresses           []string
	port                int
}

// Protocol returns the protocol this server is handling.
func (s *Server) Protocol() string {
	if s.CertFile != "" && s.KeyFile != "" {
		return ProtocolHTTPS
	}
	return ProtocolHTTP
}

// Addresses returns the host addresses being listened to.
func (s *Server) Addresses() []string {
	return s.addresses
}

// Port returns the port being listened to.
func (s *Server) Port() int {
	return s.port
}

// LocalBaseURL returns the local base URL that will reach the server.
func (s *Server) LocalBaseURL() string {
	return fmt.Sprintf("%s://%s:%d", s.Protocol(), network.IPv4LoopbackAddress, s.port)
}

func (s *Server) String() string {
	var buffer strings.Builder
	buffer.WriteString(s.Protocol())
	buffer.WriteString(" on ")
	for i, addr := range s.addresses {
		if i != 0 {
			buffer.WriteString(", ")
		}
		fmt.Fprintf(&buffer, "%s:%d", addr, s.port)
	}
	return buffer.String()
}

// Run the server. Does not return until the server is shutdown.
func (s *Server) Run() error {
	atexit.Register(s.Shutdown)
	if s.Logger == nil {
		s.Logger = slog.Default()
	}
	handler := s.WebServer.Handler
	s.WebServer.Handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		started := time.Now()
		req.URL.Path = path.Clean(req.URL.Path)
		req = req.WithContext(context.WithValue(req.Context(), routeKey, &route{path: req.URL.Path}))
		sw := &xhttp.StatusResponseWriter{
			Original: w,
			Head:     req.Method == http.MethodHead,
		}
		defer func() {
			if recovered := recover(); recovered != nil {
				err, ok := recovered.(error)
				if !ok {
					err = errs.Newf("%+v", recovered)
				}
				errs.LogTo(s.Logger, errs.NewWithCause("recovered from panic in handler", err))
				sw.WriteHeader(http.StatusInternalServerError)
			}
			since := time.Since(started)
			millis := int64(since / time.Millisecond)
			micros := int64(since/time.Microsecond) - millis*1000
			written := sw.BytesWritten()
			s.Logger.Info("web", "status", sw.Status(), "elapsed", fmt.Sprintf("%d.%03dms", millis, micros),
				"bytes", written, "method", req.Method, "url", req.URL)
		}()
		handler.ServeHTTP(sw, req)
	})
	var ln net.Listener
	host, _, err := net.SplitHostPort(s.WebServer.Addr)
	if err == nil {
		ln, err = net.Listen("tcp", s.WebServer.Addr)
	} else {
		ports := s.Ports
		if len(ports) == 0 {
			ports = []int{0}
		}
		for _, one := range ports {
			if ln, err = net.Listen("tcp", net.JoinHostPort(s.WebServer.Addr, strconv.Itoa(one))); err == nil {
				break
			}
		}
	}
	if err != nil {
		return errs.Wrap(err)
	}
	listener := network.TCPKeepAliveListener{TCPListener: ln.(*net.TCPListener)}
	var portStr string
	if _, portStr, err = net.SplitHostPort(ln.Addr().String()); err != nil {
		return errs.Wrap(err)
	}
	if s.port, err = strconv.Atoi(portStr); err != nil {
		return errs.Wrap(err)
	}
	s.addresses = network.AddressesForHost(host)
	s.Logger.Info("listening", "protocol", s.Protocol(), "addresses", s.addresses, "port", s.port)
	go func() {
		if s.StartedChan != nil {
			close(s.StartedChan)
		}
	}()
	if s.Protocol() == ProtocolHTTPS {
		err = s.WebServer.ServeTLS(listener, s.CertFile, s.KeyFile)
	} else {
		err = s.WebServer.Serve(listener)
	}
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return errs.Wrap(err)
	}
	return nil
}

// Shutdown the server gracefully.
func (s *Server) Shutdown() {
	startedAt := time.Now()
	slog.Info("starting shutdown", "protocol", s.Protocol(), "addresses", s.addresses, "port", s.port)
	defer func() {
		slog.Info("finished shutdown", "protocol", s.Protocol(), "addresses", s.addresses, "port", s.port, "elapsed",
			time.Since(startedAt))
	}()
	gracePeriod := s.ShutdownGracePeriod
	if gracePeriod <= 0 {
		gracePeriod = time.Minute
	}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(gracePeriod))
	defer cancel()
	if err := s.WebServer.Shutdown(ctx); err != nil {
		errs.LogTo(s.Logger, errs.NewWithCause("unable to shutdown gracefully", err), "protocol", s.Protocol())
	}
	if s.ShutdownCallback != nil {
		s.ShutdownCallback()
	}
}
