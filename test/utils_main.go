//go:build utils_main
package main

import (
    "bufio"
    // "os"
    "strings"
)

func main() {
    // input := stdin()
    fds := Byte("{")
    println(fds)

    {
        // reader := bufio.NewReader(os.Stdin)
        reader := bufio.NewReader(strings.NewReader("+-fds"))

        if peekStr(reader, "+-") {
            println("wouhoooouu")
        }
    }
}
