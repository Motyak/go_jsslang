package main

import (
    "runtime"
    "os"
    "bufio"
    "bytes"
    "strings"
    "fmt"
)

func tommy_str(raw_str string) string {
    _, filename, lineno, ok := runtime.Caller(1)
    if !ok {panic("runtime.Caller() not ok")}

    indent_level := 0
    /* detect indent level where the multiline str literal
       (passed as raw_str) starts */
    {
        var file_handle *os.File
        file_handle, err = os.Open(filename)
        _ = err == nil || die("Can't open file")
        defer file_handle.Close()

        reader := bufio.NewReader(file_handle)
        currLineNo := 0

        for currLineNo != lineno - 1 {
            currLineNo++
            _, err = reader.ReadString('\n')
            _ = err == nil || die("")
        }

        spaces := 0
        for true {
            var c []byte
            c, err = reader.Peek(1)
            _ = err == nil || die("")
            if c[0] != ' ' && c[0] != '|' {
                break
            }
            reader.Discard(1)
            spaces++
            if spaces % 4 == 0 {
                indent_level++
            }
        }
        if spaces % 4 != 0 {
            die(fmt.Sprintf("tommy_str() starts on a line with non-%%4 indent at:\n  %s:%d", filename, lineno))
        }
    }

    res := ""
    /* now build res string */
    {
        var currLine string
        i := 0
        reader := bytes.NewBuffer([]byte(raw_str))
        _, err = reader.ReadString('\n') // initial newline
        _ = err == nil || die("")
        for true {
            i += 1
            currLine, _ = reader.ReadString('\n')
            if len(currLine) == indent_level * 4 {
                break
            }
            start := (indent_level + 1) * 4
            if len(strings.TrimSpace(currLine[:start])) != 0 &&
            (len(strings.TrimSpace(currLine[:start])) != 1 || currLine[start-1] != '|') {
                die(fmt.Sprintf("your tommy_str() is broken at %s:%d", filename, lineno + i))
            }
            actualContent := currLine[start:]
            actualContent = strings.ReplaceAll(actualContent, `\q`, "`")
            actualContent = strings.ReplaceAll(actualContent, `\s`, " ")
            res += actualContent
        }
    }

    if len(res) == 0 {
        return ""
    }
    return res[:len(res)-1]
}
