package main
import (
    "encoding/json"
    "fmt"
)
type DebugInfo struct {
    Level string
    Msg string
    author string  // 未导出字段不会被json解析
}
func (dbgInfo DebugInfo) String() string {
    return fmt.Sprintf("{Level: %s, Msg: %s}", dbgInfo.Level, dbgInfo.Msg)
}
func main() {
    data := `[{"level":"debug","msg":"File Not Found","author":"Cynhard"},` +
        `{"level":"","msg":"Logic error","author":"Gopher"}]`
    var dbgInfos []DebugInfo
    json.Unmarshal([]byte(data), &dbgInfos)
    fmt.Println(dbgInfos)
}

