package main
import "fmt"
func main(){
	var temp string
	temp =` 		
   		 x := 10
   		 y := 20
   		 z := 30
   		 fmt.Println(x, "	", y, "	", z)
   		 x, y, z = y, z, x
   		 fmt.Println(x, "	", y, "	", z)
`
	fmt.Println(temp)
}