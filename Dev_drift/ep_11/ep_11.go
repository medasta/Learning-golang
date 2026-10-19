package ep11

import "fmt"

func main() {
	fmt.Println("if struct")
	a := 1
	b := 6
	if a < b {
		fmt.Printf(`a < b [%v]\n`, a < b)
	}

	fmt.Println("if else struct")
	if a < b {
		fmt.Printf(`a < b [%v]\n`, a < b)
	} else {
		fmt.Printf(`a > b [%v]\n`, a > b)
	}

	fmt.Println("if/elif/else struct")
	if a > b {
		fmt.Printf(`a > b [%v]\n`, a > b)
	} else if a < b {
		fmt.Printf(`a < b [%v]\n`, a < b)
	} else {
		fmt.Printf(`a == b [%v]\n`, a == b)
	}

	fmt.Printf("case struct")
	switch a == b {
	case true:
		fmt.Println("a == b")
	default:
		fmt.Println("a != b")
	}
}
