package main
import (
	"encoding/json"
	"fmt"
)
func main() {
	data := `[{"Level":"debug","Msg":"File: \"test.txt\" Not Found"},` +
		`{"Level":"","Msg":"Logic error"}]`
	var dbgInfos []map[string]string
	json.Unmarshal([]byte(data), &dbgInfos)
	fmt.Println(dbgInfos)
}


