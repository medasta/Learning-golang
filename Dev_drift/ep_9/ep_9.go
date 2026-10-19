package ep9

import "fmt"

func main() {
	var one = 1
	var two = 2

	var total = one + two
	fmt.Printf(`Total: [%d]`, total)
	fmt.Printf(`[+] %d`, one+two)
	fmt.Printf(`[-] %d`, one-two)
	fmt.Printf(`[*] %d`, one*two)
	fmt.Printf(`[/] %d`, one/two)
	fmt.Printf(`[mod] %d`, one%two)
	one++
	two++
	one--
	two--
}
