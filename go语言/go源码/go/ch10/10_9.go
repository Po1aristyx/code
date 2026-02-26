package main
import (
	"encoding/json"
	"fmt"
)
func main() {
    m := map[string][]string{
        "level":   {"debug"},
        "message": {"File not found", "Stack overflow"},
    }
    if data, err := json.MarshalIndent(m, "", " "); err == nil {
        fmt.Printf("%s\n", data)
    }
}

