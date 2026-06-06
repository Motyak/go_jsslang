//go:build tommystr_main
package main

func main() {
    // str := tommy_str(`
    // `)

    // str := tommy_str(`
    //     f
    //         d
    //             s
    // `)

    str := tommy_str(`
       |-> SquareBracketsGroup
       |  -> Term #1
       |    -> Word: Atom: `+"`"+`fds`+"`"+`
       |  -> Term #2
       |    -> Word: Atom: `+"`"+`sdf`+"`"+`
    `)

    println("`" + str + "`")
}
