package main
import (
    "encoding/json"
    "fmt"
)
type Point struct{ X, Y int }
type Circle struct {
    Point
    Radius int
}
func main() {
    data := `{"X":80,"Y":80,"Radius":40}`
    var c Circle
    json.Unmarshal([]byte(data), &c)
    fmt.Println(c)
}

