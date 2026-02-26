package main
import "fmt"
func main() {
	personSalary := map[string]int{
		"Steven": 18000,
		"Daniel": 5000,
		"Josh":   20000,
	}
	fmt.Println("原始薪资：", personSalary)
	newPersonSalary := personSalary
	newPersonSalary["Daniel"] = 8000
	fmt.Println("修改后newPersonSalary：", newPersonSalary)
	fmt.Println("personSalary受影响情况：", personSalary)
}


