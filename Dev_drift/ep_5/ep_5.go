package ep5

import "fmt"

var arr = []int{1, 2, 3, 4, 5}
var arr2 = [10]int{1, 2, 3} // {1,2,3,0...0}

func main() {
	arr = append(arr, 200) // добавление элемента в конец массива
	fmt.Printf(`Array:%v`, arr[0])
}
