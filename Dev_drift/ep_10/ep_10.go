package ep10

import "fmt"

func main() {
	var a = 5
	var b = 2

	// Операции отношения
	fmt.Printf("a == b: [%v]\n", a == b)
	fmt.Printf("a > b: [%v]\n", a > b)
	fmt.Printf("a < b: [%v]\n", a < b)
	fmt.Printf("a <= b: [%v]\n", a <= b)
	fmt.Printf("a >= b: [%v]\n", a >= b)
	fmt.Printf("a != b: [%v]\n", a != b)
	// Логические операции
	var aTrue = true
	var bFalse = false
	fmt.Printf("!a: [%v]", !aTrue)
	fmt.Printf("a&&b: [%v]", aTrue && bFalse)
	fmt.Printf("a||b: [%v]", aTrue || bFalse)

}
