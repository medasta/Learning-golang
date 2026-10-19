package main

import "fmt"

var users = map[string]int{
	`Admin`:   1,
	`Manager`: 2,
}

func main() {
	// добавление нового элемента(В конец словаря)
	users["User-1"] = 100

	fmt.Printf(`%v`, users)
	fmt.Printf(`Admin = %v\n`, users[`Admin`])
	fmt.Printf(`Manager = %v\n`, users[`Manager`])
	users[`Manager`] = 10 // Переопределение хранящегося значения по ключу Manager
	fmt.Printf(`Manager = %v\n`, users[`Manager`])

	// Перебор всех ключ:значений в словаре users
	for key, value := range users {
		fmt.Printf("Key = %s Value = %d\n", key, value)
	}
	//delete(словарь, ключ) - удаление пары ключ:значение в заданном словаре
	delete(users, "User-1")
	for key, value := range users {
		fmt.Printf("Key = %s Value = %d \n", key, value)
	}
}
