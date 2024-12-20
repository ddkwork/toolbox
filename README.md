[![Go Reference](https://pkg.go.dev/badge/github.com/ddkwork/toolbox.svg)](https://pkg.go.dev/github.com/ddkwork/toolbox)
[![Go Report Card](https://goreportcard.com/badge/github.com/ddkwork/toolbox)](https://goreportcard.com/report/github.com/ddkwork/toolbox)

# toolbox
Toolbox for Go.

To install this package and the tools it provides:
```
./build.sh
```

> NOTE: This library already had a v1.x.y when Go modules were first introduced. Due to this, it doesn't follow the
>       normal convention and instead treats its releases as if they are of the v0.x.y variety (i.e. it could introduce
>       breaking API changes). Keep this in mind when deciding whether or not to use it.

## Package summaries

### toolbox
Utilities that didn't have a home elsewhere.

### toolbox/atexit
Provides functionality similar to the C standard library's atexit() call. To function properly, use
`atexit.Exit(result)` rather than `os.Exit(result)`.

```go
package main

import (
    "fmt"

    "github.com/ddkwork/atexit"
)

func main() {
    atexit.Register(func() { fmt.Println("Goodbye!") })
    atexit.Register(func() { fmt.Println("Preparing to exit.") })
    doSomeStuff()
    // Always end your program with a call to atexit.Exit().
    // Because we know doSomeStuff() always exits in this example, this isn't actually needed here.
    atexit.Exit(0)
}

func doSomeStuff() {
    fmt.Println("in doSomeStuff()")
    fmt.Println("going to do the equivalent of os.Exit(1)")
    atexit.Exit(1)
}
```
Output:
```
in doSomeStuff()
going to do the equivalent of os.Exit(1)
Preparing to exit.
Goodbye!
```

`atexit.Exit()` runs any registered exit functions in the inverse order they were registered and then exits with the
specified status. If a previous call to `atexit.Exit()` is already being handled, the call does nothing but does not
return. Recursive calls to `atexit.Exit()` will trigger a panic, which the exit handling will catch and report, but
will then proceed with exit as normal. Note that once `atexit.Exit()` is called, no subsequent changes to the
registered list of functions will have an effect (i.e. you cannot `atexit.Unregister()` a function inside an exit
handler to prevent its execution, nor can you `atexit.Register()` a new function).

### toolbox/cmdline
Command line handling. Provides the tool `genversion` for generating version numbers with an embedded date.

### toolbox/collection
Provides type-safe sets for the various primitive types.

### toolbox/collection/dict
Provides some useful `map` functions that were inexplicably left out of the `maps` package introduced in Go 1.21.

### toolbox/collection/quadtree
Provides an implementation of a [Quadtree](https://en.wikipedia.org/wiki/Quadtree).

### toolbox/collection/slice
Provides some useful `slice` functions that were inexplicably left out of the `slices` package introduced in Go 1.21.

### toolbox/desktop
Desktop integration utilities.

### toolbox/errs
Provides an error that contains a stack trace with source locations, along with nested causes, if any, and is also
capable of containing multiple top-level errors.

Example output of a single error:
```
example message with a stack trace
    [play/example.MyOtherFunction] other.go:10
    [play/example.MyFunction] example.go:4
    [main.main] main.go:8
```

Example output of an error that contains multiple errors:
```
Multiple (2) errors occurred:
- first error
- second error
    [play/example.MultipleErrors] example.go:14
    [main.main] main.go:9
```

#### Panic recovery helper
An easy way to run code that may panic:
```go
func runSomeCode() {
    defer errs.Recovery(errs.Log)
    // ... run the code here ...
}
```

### toolbox/eval
Dynamically evaluate expressions.

### toolbox/formats/icon
Provides image scaling and stacking utilities.

### toolbox/formats/icon/icns
Provides an encoder for [Apple Icon Image](https://en.wikipedia.org/wiki/Apple_Icon_Image_format) files.

### toolbox/formats/icon/ico
Provides an encoder for [Windows Icon Image](https://en.wikipedia.org/wiki/ICO_(file_format)) files.

### toolbox/formats/json
Manipulation of JSON data.

### toolbox/formats/xlsx
Extract text from Excel spreadsheets.

### toolbox/i18n
Internationalization support for applications. Provides the tool `go-i18n` for generating a template for a localization
file from source code.

### toolbox/log/rotation
Provides file rotation when files hit a given size.

### toolbox/notifier
Provides a mechanism for tracking targets of notifications and methods for notifying them.

### toolbox/rate
Rate limiting which supports a hierarchy of limiters, each capped by their parent.

### toolbox/softref
Soft references.

### toolbox/taskqueue
Provides a simple asynchronous task queue.

### toolbox/txt
Text utilities.

### toolbox/vcs/git
git repository access.

### toolbox/xcrypto
Provides convenience utilities for encrypting and decrypting streams of data with public & private keys.

### toolbox/xio
io utilities.

### toolbox/xio/fs
Filesystem utilities.

### toolbox/xio/fs/paths
Platform-specific standard paths.

### toolbox/xio/fs/safe
Safe, atomic saving of files.

### toolbox/xio/fs/tar
Provides tar file extraction utilities.

### toolbox/xio/fs/zip
Provides zip file extraction utilities.

### toolbox/xio/network
Network-related utilities.

### toolbox/xio/network/natpmp
Implementation of [NAT-PMP](https://tools.ietf.org/html/rfc6886).

### toolbox/xio/network/xhttp
HTTP-related utilities.

### toolbox/xio/network/xhttp/web
Web server with some standardized logging and handler wrapping.

### toolbox/xio/term
Terminal utilities.

### toolbox/xmath
Math utilities.

### toolbox/xmath/crc
Provides CRC helpers.

### toolbox/xmath/fixed
Fixed-point types with a configurable number of decimal places. These types implement the marshal/unmarshal interfaces
for JSON and YAML.

### toolbox/xmath/geom
Geometry primitives.

### toolbox/xmath/geom/poly
Provides polygon boolean operations. These are not as robust as I'd like.

### toolbox/xmath/geom/visibility
Calculates a visibility polygon from a given point in the presence of a set of obstructions, also known as an [Isovist](https://en.wikipedia.org/wiki/Isovist).

### toolbox/xmath/num
128-bit int and uint types. These types implement the marshal/unmarshal interfaces for JSON and YAML.

### toolbox/xmath/rand
Randomizer based upon the crypto/rand package.
