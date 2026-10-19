package ep8

import "fmt"

// константы - не изменяемый тип данных

var (
	str = "Строка-1"
)

const NotChange = `Константа`

func main() {
	str = `Изменить состояние`
	// NotChange = "str"
	fmt.Printf("%s", str)
	fmt.Printf("%s", NotChange)
}
