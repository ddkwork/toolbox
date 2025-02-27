// Copyright ©2016-2022 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package fs

import (
	"bufio"
	"io"
	"io/fs"
	"os"

	"github.com/ddkwork/toolbox/errs"
	"github.com/ddkwork/toolbox/xio"
	"github.com/ddkwork/toolbox/xio/fs/safe"

	"gopkg.in/yaml.v3"
)

// LoadYAML data from the specified path.
func LoadYAML(path string, data any) error {
	f, err := os.Open(path)
	if err != nil {
		return errs.NewWithCause(path, err)
	}
	return loadYAML(f, path, data)
}

// LoadYAMLFromFS data from the specified filesystem path.
func LoadYAMLFromFS(fsys fs.FS, path string, data any) error {
	f, err := fsys.Open(path)
	if err != nil {
		return errs.NewWithCause(path, err)
	}
	return loadYAML(f, path, data)
}

func loadYAML(r io.ReadCloser, path string, data any) error {
	defer xio.CloseIgnoringErrors(r)
	if err := yaml.NewDecoder(bufio.NewReader(r)).Decode(data); err != nil {
		return errs.NewWithCause(path, err)
	}
	return nil
}

// SaveYAML data to the specified path.
func SaveYAML(path string, data any) error {
	return SaveYAMLWithMode(path, data, 0o644)
}

// SaveYAMLWithMode data to the specified path.
func SaveYAMLWithMode(path string, data any, mode os.FileMode) error {
	if err := safe.WriteFileWithMode(path, func(w io.Writer) error {
		encoder := yaml.NewEncoder(w)
		encoder.SetIndent(2)
		if err := encoder.Encode(data); err != nil {
			return errs.Wrap(err)
		}
		return errs.Wrap(encoder.Close())
	}, mode); err != nil {
		return errs.NewWithCause(path, err)
	}
	return nil
}
