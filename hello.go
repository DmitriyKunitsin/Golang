package main

///  с помощью оператора package указывается, к какому пакету будет принадлежать файл.
/// В данном случае это пакет main

import "fmt"

// с помощью директивы import мы подключаем этот пакет

func main() {
	// Далее идет функция main. Это главная функция любой программы на Go. По сути все, что выполняется в программе, выполняется именно функции main.
	fmt.Println("Hello Go!")
}
