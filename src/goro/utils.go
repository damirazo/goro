package goro

import "fmt"

// Функция обработки ошибок
func HandleError(err error) {
	if err != nil {
		fmt.Println("Error: ", err.Error())
		panic(err)
	}
}
