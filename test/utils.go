package main

import (
    "fmt"
    "io"
    "os"
    "runtime"
    "strings"
    "bufio"
)

var err error

func die(msg any) bool {
    fmt.Fprintf(os.Stderr, "die(): %s\n", msg)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
    fmt.Fprintln(os.Stderr)

    _, _, _, ok := runtime.Caller(1)
    if !ok { // fallback
        panic(msg)
    }

    buf := make([]byte, 4096)
    n := runtime.Stack(buf, false)
    str := string(buf[:n])
    str = strings.ReplaceAll(str, "\t", "  ")
    lines := strings.Split(str, "\n")
    if len(lines) > 3 {
        str = strings.Join(append(lines[0:1], lines[3:]...), "\n")
    }
    fmt.Fprintln(os.Stderr, str)
    os.Exit(2)
    return false
}

func ASSERT(test bool) {
    if test == true {
        return
    }

    _, filename, lineno, ok := runtime.Caller(1)

    if !ok { // fallback
        panic("Assertion failed")
    }

    fmt.Fprintf(os.Stderr, "Assertion failed at:\n  %s:%d\n\n", filename, lineno)
    buf := make([]byte, 4096)
    n := runtime.Stack(buf, false)
    str := string(buf[:n])
    str = strings.ReplaceAll(str, "\t", "  ")
    lines := strings.Split(str, "\n")
    if len(lines) > 3 {
        str = strings.Join(append(lines[0:1], lines[3:]...), "\n")
    }
    fmt.Fprintln(os.Stderr, str)
    os.Exit(2)
}

func SHOULD_NOT_HAPPEN() {
    _, filename, lineno, ok := runtime.Caller(1)

    if !ok { // fallback
        panic("Should not happen")
    }

    fmt.Fprintf(os.Stderr, "`Should not happen` raised at:\n  %s:%d\n\n", filename, lineno)
    buf := make([]byte, 4096)
    n := runtime.Stack(buf, false)
    str := string(buf[:n])
    str = strings.ReplaceAll(str, "\t", "  ")
    lines := strings.Split(str, "\n")
    if len(lines) > 3 {
        str = strings.Join(append(lines[0:1], lines[3:]...), "\n")
    }
    fmt.Fprintln(os.Stderr, str)
    os.Exit(2)
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

func peekStr(reader *bufio.Reader, str string) bool {
    var peeked []byte
    peeked, err = reader.Peek(len(str))
    if err == io.EOF {
        return false
    }
    _ = err == nil || die("")
    match := true
    for i, c := range peeked {
        if c != str[i] {
            match = false
            break
        }
    }
    return match
}
