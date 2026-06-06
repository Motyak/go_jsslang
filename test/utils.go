package main

import "fmt"
import "io"
import "os"

var err error

func die(msg any) bool {
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
    panic(msg)
}

func ASSERT(test bool) {
    if test == false {
        panic("Assertion failed")
    }
}

func SHOULD_NOT_HAPPEN() bool {
    panic("Should not happen")
}

func stdin() []byte {
    var stdin []byte
    stdin, err = io.ReadAll(os.Stdin)
    _ = err == nil || die("failed to read STDIN")
    return stdin
}

func Byte(c string) byte {
    _ = len(c) == 1 || die("Byte() arg should be a 1-char string")
    return c[0]
}
