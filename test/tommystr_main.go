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

    // str := tommy_str(`
    //     \s\s
    // `)

    str := tommy_str(`
       |-> SquareBracketsGroup
       |  -> Term #1
       |    -> Word: Atom: \qfds\q
       |  -> Term #2
       |    -> Word: Atom: \qsdf\q
    `)

    println("`" + str + "`")
}
