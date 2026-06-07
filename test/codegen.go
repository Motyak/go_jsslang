package main

import (
    "fmt"
    "io"
    "bytes"
)

func codegen_(val value_t) string {
    oss := bytes.NewBufferString("")
    codegen(oss, val)
    return oss.String()
}

func codegen(writer io.Writer, val value_t) {
    visit(val, value_t__Visitor{
        Nil: func() {
            fmt.Fprintf(writer, "processing value_t__Nil\n")
        },

        Bool: func(bool_ value_t__Bool) {
            fmt.Fprintf(writer, "processing value_t__Bool: %t\n", bool_)
        },

        Float: func(float_ value_t__Float) {
            fmt.Fprintf(writer, "processing value_t__Float: %f\n", float_)
        },

        Str: func(str_ value_t__Str) {
            fmt.Fprintf(writer, "processing value_t__Str: %v\n", str_)
        },

        List: func(list value_t__List) {
            fmt.Fprintf(writer, "processing value_t__List: %s\n", list)
        },

        Map: func(map_ value_t__Map) {
            fmt.Fprintf(writer, "processing value_t__Map: %s\n", map_)
        },
    })
}
