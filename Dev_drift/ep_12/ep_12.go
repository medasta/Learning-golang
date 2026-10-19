package ep12

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Value: [%v]\n", i)
	}

	j := 0
	for ; j < 10; j++ {

	}

	k := 0
	for k < 10 {
		k += 1
	}

	for n := 0; n < 10; n++ {
		for m := 0; m < 10; m++ {
			fmt.Printf("Value n: [%d], Value m: [%d]\n", n, m)
		}
	}

	var arr = [5]int{1, 2, 3, 4, 5}

	for index, value := range arr {
		fmt.Printf("index [%v], value [%v]\n", index, value)
	}

	for _, value := range arr {
		fmt.Printf("value [%v]\n", value)
	}
	// break - экстренный вызов из цикла
	// continue - пропуск итерации
}
