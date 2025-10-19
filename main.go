package main

import "fmt"

func main() {
	fmt.Println("Запуск системы управления библиотекой")

	myLibrary := &Library{}

	fmt.Println("Наполняем библиотеку")

	myLibrary.AddReader("Агунда", "Кокойти")
	myLibrary.AddReader("Сергей", "Меняйло")

	myLibrary.AddBook("1984", "Джордж Оруэлл", 1949)
	myLibrary.AddBook("Мастер и Маргарита", "Михаил Булгаков", 1967)

	fmt.Println("\n---Библиотека готова к работе---")
	fmt.Println("Количество читателей:", len(myLibrary.Readers))
	fmt.Println("Количество книг:", len(myLibrary.Books))

	fmt.Println("---Тестируем выдачу книг---")

	fmt.Println("-----------------------------")
	err := myLibrary.IssueBookToReader(1, 1)
	if err != nil {
		fmt.Println("ошибка выдачи: ", err)
	} else {
		fmt.Println("Книга успешно выдана")
	}

	err = myLibrary.ReturnBook(1)
	if err != nil {
		fmt.Println("", err)
	} else {
		fmt.Println("Книга успешно выдана")
	}
	
	err = myLibrary.ReturnBook(1)
	if err != nil {
		fmt.Println("", err)
	} else {
		fmt.Println("Книга успешно выдана")
	}
}
