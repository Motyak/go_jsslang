//go:build variant_main
package main

import (
    "fmt"
)

func main() {
    // var val = value_t{float64(123)}
    // var val = value_t{bool(true)}
    // var val = value_t{[]value_t{value_t{1}, value_t{2}, value_t{3}}}
    // var val = value_t{map[value_t]value_t{
    //     value_t{"a"}: value_t{1},
    // }}
    var val = value_t{}
    doit(val)
    
    if val.is_nil() {
        return
    }
    if val.holds(Bool) {
        var bool_ = val.variant.(value_t__Bool)
        fmt.Printf("bool: %t\n", bool_)
    }
}
