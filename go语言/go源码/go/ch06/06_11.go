package main
import (
	"fmt"
	"math"
)
func main() {
	fmt.Println(math.IsNaN(3.4))      //false
	fmt.Println(math.Ceil(1.000001))  //2
	fmt.Println(math.Floor(1.999999)) //1
	fmt.Println(math.Trunc(1.999999)) //1
	fmt.Println(math.Abs(-1.3))       //1.3
	fmt.Println(math.Max(-1.3, 0))    //0
	fmt.Println(math.Min(-1.3, 0))    //-1.3
	fmt.Println(math.Dim(-12, -19))   //7
	fmt.Println(math.Dim(-12, 19))    //0
	fmt.Println(math.Mod(9, 4))       //1
	fmt.Println(math.Sqrt(9))         //3
	fmt.Println(math.Cbrt(8))         //2
	fmt.Println(math.Hypot(3, 4))     //5
	fmt.Println(math.Pow(2, 8))       //256
	fmt.Println(math.Log(1))          //0
	fmt.Println(math.Log2(16))        //4
	fmt.Println(math.Log10(1000))     //3
}
