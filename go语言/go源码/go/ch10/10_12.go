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
    if data, err := json.Marshal(Circle{Point{50, 50}, 25}); err == nil {
        fmt.Printf("%s\n", data)
    }
}


