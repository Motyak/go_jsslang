package main

import (
    "runtime"
    "os"
    "bufio"
    "bytes"
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
            _ = err == nil || SHOULD_NOT_HAPPEN()
        }

        spaces := 0
        for true {
            var c []byte
            c, err = reader.Peek(1)
            _ = err == nil || SHOULD_NOT_HAPPEN()
            if c[0] != ' ' && c[0] != '|' {
                break
            }
            reader.Discard(1)
            spaces++
            if spaces % 4 == 0 {
                indent_level++
            }
        }
    }

    res := ""
    /* now build res string */
    {
        var currLine string
        reader := bytes.NewBuffer([]byte(raw_str))
        _, err = reader.ReadString('\n') // initial newline
        _ = err == nil || SHOULD_NOT_HAPPEN()
        for true {
            currLine, _ = reader.ReadString('\n')
            if len(currLine) == indent_level * 4 {
                break
            }
            actualContent := currLine[(indent_level+1)*4:]
            res += actualContent
        }
    }

    if len(res) == 0 {
        return ""
    }
    return res[:len(res)-1]
}
