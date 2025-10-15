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

	err := myLibrary.IssueBookToReader(1, 1)
	if err != nil {
		fmt.Println("Ошибка выдачи", err)
	}

	book, _ := myLibrary.FindBookByID(1)
	if book != nil {
		fmt.Println("Статус книги после выдачи:", book)
	}

	err = myLibrary.IssueBookToReader(99, 1)
	if err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	}
	
	


}