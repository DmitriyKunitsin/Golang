/*
	Срезы (Slice) последовательность элементов одного типа
	В отличие от массивов длина (размер) не фиксирован
	и динамически может меняться
*/
package main

import "fmt"

func main() {
	var users = []string{"Dima", "Andrey", "Vova"}
	fmt.Println(users[2])
	fmt.Println()
	for _, value := range users {
		fmt.Println(value)
	}

	var makeString []string = make([]string, 3)
	for i, value := range users {
		makeString[i] = value
		fmt.Println(makeString[i])
	}
	for _, value := range makeString {
		users = append(users, value)
	}
	fmt.Printf("\n\n")
	for _, value := range users[:3] {
		fmt.Println(value)
	}

	user1 := users[2:6] // с 3го по 6й
	user2 := users[3:] // с 4 до конца
	user3 := users[:4] // с 1 по 4
	fmt.Println()
	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(user3)
}