package main

import (
	"fmt"
	"library-app/cmd/cli"
	"library-app/library"
	"library-app/storage"
)

func main() {
	myLibrary := library.New()
	var err error
	myLibrary.Books, err = storage.LoadBooksFromCSV("books.csv")
	if err != nil {
		fmt.Println("Поизошла ошибка базы:", err)
	}

	myLibrary.Readers, err = storage.LoadReadersFromCSV("readers.csv")
	if err != nil {
		fmt.Println("Поизошла ошибка базы:", err)
	}

	cli.Run(myLibrary)

	// myLibrary.AddBook("Война и мир", "Лев Толстой", 1763)
	// myLibrary.AddBook("Преступление и наказание", "Ф.М. Достоевский", 1343)

	// // Добавляем читателей
	// myLibrary.AddReader("Иван", "Иванов")
	// myLibrary.AddReader("Валерий", "Кузьмин")
	// myLibrary.IssueBookToReader(1, 1)

	// fmt.Println("\n--- Библиотека готова к работе ---")
	// fmt.Println("Количество читателей:", len(myLibrary.Readers))
	// fmt.Println("Количество книг:", len(myLibrary.Books))

	// fmt.Println("\n---Тестируем выдачу книг---")

	// err := myLibrary.IssueBookToReader(1, 1)
	// if err != nil {
	//     fmt.Println("Ошибка выдачи:", err)
	// }

	// book, _ := myLibrary.FindBookByID(1)
	// if book != nil {
	//     fmt.Println("Статус книги после выдачи:", book)
	// }

	// err = myLibrary.IssueBookToReader(99, 1)
	// if err != nil {
	//     fmt.Println("Ожидаемая ошибка:", err)
	// }

	// fmt.Println()

	// books := myLibrary.GetAllBooks()
	// for _, book := range books{
	// 	fmt.Println(book)
	// }

	// fmt.Println("Запуск системы управления библиотекой...")

	// myLibrary := &Library{}

	// fmt.Println("\n--- Наполняем библиотеку ---")

	// myLibrary.AddReader("Агунда", "Кокойты")
	// myLibrary.AddReader("Сергей", "Меняйло")

	// myLibrary.AddBook("1984", "Джордж Оруэлл", 1949)
	// myLibrary.AddBook("Мастер и Маргарита", "Михаил Булгаков", 1967)

}
