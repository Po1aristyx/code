package main
import (
    "encoding/json"
    "fmt"
)
type DebugInfo struct {
    Level  string `json:"level"`   // level 解码为 Level
    Msg    string `json:"message"` // message 解码为 Msg
    Author string `json:"-"`       // 忽略Author
}
func (dbgInfo DebugInfo) String() string {
    return fmt.Sprintf("{Level: %s, Msg: %s}", dbgInfo.Level, dbgInfo.Msg)
}
func main() {
    data := `[{"level":"debug","message":"File Not Found","author":"Cynhard"},` +
        `{"level":"","message":"Logic error","author":"Gopher"}]`
    var dbgInfos []DebugInfo
    json.Unmarshal([]byte(data), &dbgInfos)
    fmt.Println(dbgInfos)
}


