//go:build codegen_main
package main

import (
    // "bytes"
    // "fmt"
    "os"
)

func main() {
    {
        oss := bytes.NewBufferString("")
        val := value_t{float64(123)}
        codegen(oss, val)
        print(oss.String())
    }

    {
        val := value_t{float64(123)}
        codegen(os.Stdout, val)
    }
}
