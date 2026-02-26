package main
import (
	"encoding/json"
	"fmt"
)
type DebugInfo struct {
	Level  string
	Msg    string
	author string // 未导出字段不会被json解析(首字母小写)
}
func main() {
	dbgInfs := []DebugInfo{
		DebugInfo{"debug", `File: "test.txt" Not Found`, "Cynhard"},
		DebugInfo{"", "Logic error", "Gopher"},
	}
	if data, err := json.Marshal(dbgInfs); err == nil {
		fmt.Printf("%s\n", data)
	}
}


