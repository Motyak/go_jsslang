package main

import (
    "bufio"
    "io"
)

func parse(reader io.Reader) value_t {
    // reader_ := bufio.NewReader(reader)
    _ = bufio.NewReader(reader)

    return value_t{}
}
