package ep3

import "fmt"

var GlobalProjectVar = "Hello :-)" //Visible on package import
var packageVar = "hello"           // Visible inside this package only

func main() {
	val := 10
	fmt.Println(val)
}
