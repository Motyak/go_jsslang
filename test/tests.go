//go:build tests
package main

import (

    "fmt"
    "runtime"
    "os"
    "strings"
)

var _ =
TEST_CASE("test name", "[test-1111][err]", func(){
    println("this is test 1")
})

var _ =
TEST_CASE("test name 2", "[test-1112][err]", func(){
    println("this is test 2")
    REQUIRE (false)
    // ASSERT (false)
})

////////////////////////////////////////////////////////////
// test execution logic
////////////////////////////////////////////////////////////

func REQUIRE(test bool) {
    if test == true {
        return
    }

    _, filename, lineno, ok := runtime.Caller(1)

    if !ok { // fallback
        panic("Requirement unmet")
    }

    fmt.Fprintf(os.Stderr, "Requirement unmet at:\n  %s:%d\n\n", filename, lineno)
    buf := make([]byte, 4096)
    n := runtime.Stack(buf, false)
    str := string(buf[:n])
    str = strings.ReplaceAll(str, "\t", "  ")
    lines := strings.Split(str, "\n")
    if len(lines) > 3 {
        str = strings.Join(append(lines[0:1], lines[3:]...), "\n")
    }
    fmt.Fprintln(os.Stderr, str)
    os.Exit(1)
}

type dummy struct {}

type TestCase struct {
	name string
	tags string
	test  func()
}

var test_cases []TestCase

func TEST_CASE(name string, tags string, test func()) dummy {
    test_cases = append(test_cases, TestCase{name, tags, test})
    return dummy{}
}

func main() {
    /* filter test_cases based on tag args */
    {
        filtered := test_cases[:0]
        for _, test_case := range test_cases {
            should_append := true
            for _, tag_arg := range os.Args[1:] {
                if !strings.Contains(test_case.tags, "[" + tag_arg + "]") {
                    should_append = false
                    break
                }
            }
            if should_append {
                filtered = append(filtered, test_case)
            }
        }
        test_cases = filtered
    }

    total_tests := len(test_cases)
    for i, test_case := range test_cases {
        fmt.Printf("%d/%d passed tests\n", i, total_tests)
        println("=== " + test_case.name + " " + test_case.tags + " ===")
        test_case.test()
    }
    fmt.Printf("OK %d/%d passed tests\n", total_tests, total_tests)
}
