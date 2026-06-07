//go:build tests
package main

import (

    "fmt"
    "runtime"
    "os"
    "strings"
)

var _ =
TEST_CASE("$nil", "[test-1111]", func(){
    input := "$nil"
    expect := "null"

    input_iss := strings.NewReader(input)
    output := parse(input_iss)
    output_str := codegen_(output)
    REQUIRE_STR_EQ (output_str, expect)
})

var _ =
TEST_CASE("test name 2", "[test-1112]", func(){
    println("this is test 2")
    // REQUIRE (false)
    // ASSERT (false)
})

////////////////////////////////////////////////////////////
// test execution logic
////////////////////////////////////////////////////////////

// func REQUIRE(test bool) {
//     if test == true {
//         return
//     }

//     _, filename, lineno, ok := runtime.Caller(1)

//     if !ok { // fallback
//         panic("Requirement unmet")
//     }

//     fmt.Fprintf(os.Stderr, "Requirement unmet at:\n  %s:%d\n\n", filename, lineno)
//     buf := make([]byte, 4096)
//     n := runtime.Stack(buf, false)
//     str := string(buf[:n])
//     str = strings.ReplaceAll(str, "\t", "  ")
//     lines := strings.Split(str, "\n")
//     if len(lines) > 3 {
//         str = strings.Join(append(lines[0:1], lines[3:]...), "\n")
//     }
//     fmt.Fprintln(os.Stderr, str)
//     os.Exit(1)
// }

func REQUIRE_STR_EQ(actual string, expected string) {
	if actual == expected {
		return
	}

	_, filename, lineno, ok := runtime.Caller(1)
	if !ok {
		panic("Requirement unmet")
	}

	// 1. Generate a character-by-character visual diff
	var diff strings.Builder
	runesActual := []rune(actual)
	runesExpected := []rune(expected)
	maxLen := len(runesActual)
	if len(runesExpected) > maxLen {
		maxLen = len(runesExpected)
	}

	diff.WriteString("--- String Mismatch Details ---\n")
	diff.WriteString(fmt.Sprintf("Actual Length:   %d\nExpected Length: %d\n\n", len(actual), len(expected)))
	diff.WriteString("Index | Actual Char | Expected Char | Match?\n")
	diff.WriteString("--------------------------------------------\n")

	for i := 0; i < maxLen; i++ {
		actChar := "<EOF>"
		expChar := "<EOF>"
		matchMarker := "❌"

		if i < len(runesActual) {
			actChar = fmt.Sprintf("%q", string(runesActual[i]))
		}
		if i < len(runesExpected) {
			expChar = fmt.Sprintf("%q", string(runesExpected[i]))
		}
		if i < len(runesActual) && i < len(runesExpected) && runesActual[i] == runesExpected[i] {
			matchMarker = "✓"
		}

		diff.WriteString(fmt.Sprintf("%5d | %11s | %13s | %s\n", i, actChar, expChar, matchMarker))
	}
	diff.WriteString("--------------------------------------------\n")

	// 2. Print metadata, error diff, and cleaned stack trace
	fmt.Fprintf(os.Stderr, "Requirement unmet at:\n  %s:%d\n\n", filename, lineno)
	fmt.Fprintln(os.Stderr, diff.String())

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
