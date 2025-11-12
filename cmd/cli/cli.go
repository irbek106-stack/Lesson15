package cli

import (
	"bufio"
	"fmt"
	"library-app/library"
	"library-app/storage"
	"os"
	"strconv"
)

func printMenu() {

	fmt.Println("--SIMPLE-LIBRARY--")
	fmt.Println("----------------------------------------")
	fmt.Println("-MENU-")
	fmt.Println("[0] Выход")
	fmt.Println("[1] Добавление книги")
	fmt.Println("[2] Выдача книги")
	fmt.Println("[3] Возврат книги")
	fmt.Println("[4] Поиск книги")
	fmt.Println("[5] Импорт книги")
	fmt.Println("[6] Экспорт книги")
	fmt.Println("[7] Вывод всех книг")
	fmt.Println("[8] Добавление читателя")
	fmt.Println("[9] Поиск читателя")
	fmt.Println("[10] Импорт читателя")
	fmt.Println("[11] Экспорт читателя")
	fmt.Println("[12] Вывод всех читателей")

}

func handlerChoice(choice int, scanner *bufio.Scanner, library *library.Library) {
	switch choice {
	//выход
	case 0:
		fmt.Println("Bye")
		//добавление книг
	case 1:
		fmt.Println("Введите название книги: ")
		scanner.Scan()
		title := scanner.Text()

		fmt.Println("Введите автора книги: ")
		scanner.Scan()
		author := scanner.Text()

		fmt.Println("Введите год книги: ")
		scanner.Scan()
		year, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Год должен состоять из цифр")
			return
		}

		if _, err := library.AddBook(title, author, year); err != nil {
			fmt.Printf("Произошла ошибка при добавление книги:%v", err)
			return
		}

		fmt.Printf("Книга %s успешно добавлена\n", title)

		// выдача книг читателю
	case 2:
		fmt.Println("Введите номер книги: ")
		scanner.Scan()
		idBook, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Номер должен состоять из цифр!")
			return
		}

		fmt.Println("Введите номер Читателя: ")
		scanner.Scan()
		idUser, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Номер должен состоять из цифр!")
			return
		}
		if err := library.IssueBookToReader(idBook, idUser); err != nil {
			fmt.Printf("Произошла ошибка при выдаче книги:%v", err)
			return
		}
		fmt.Printf("Книга %d выдана пользователю %d\n", idBook, idUser)
		//возврат книги
	case 3:
		fmt.Println("Введите номер книги: ")
		scanner.Scan()
		idBook, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Номер должен состоять из цифр!")
			return
		}
		err = library.ReturnBook(idBook)
		if err != nil {
			fmt.Println("Ошибка возврата книги:", err)
			return
		}
		fmt.Printf("Книга %d возвращена\n", idBook)
	// поиск книги
	case 4:
		fmt.Println("Введите номер книги: ")
		scanner.Scan()
		idBook, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Номер должен состоять из цифр!")
			return
		}
		book, err := library.FindBookByID(idBook)
		if err != nil {
			fmt.Println("Ошибка при поиске ", err)
			return
		}
		fmt.Println(book)
	//Импорт списка книг
	case 5:
		fmt.Println("Введите название файла: ")
		scanner.Scan()
		filename := scanner.Text()
		books, err := storage.LoadBooksFromCSV(filename)
		if err != nil {
			fmt.Println("Ошибка при импорте списка", err)
			return
		}
		library.Books = books
		fmt.Println("Список книг успешно импортирован!")

	//экспорт книг
	case 6:
		fmt.Println("Введите название файла: ")
		scanner.Scan()
		filename := scanner.Text()
		err := storage.SaveBooksToCSV(filename, library.Books)
		if err != nil {
			fmt.Println("Ошибка экспорта: ", err)
			return
		}
		fmt.Println("Книги успешно экспортированы")

		//Вывод всех книг
	case 7:
		books := library.GetAllBooks()
		fmt.Println("Список книг: ")
		for _, book := range books {
			fmt.Println(book)
		}

		//Добавление читателя
	case 8:
		fmt.Println("Введите имя читателя")
		scanner.Scan()
		firstName := scanner.Text()
		fmt.Println("Введите фамилию читателя")
		scanner.Scan()
		lastName := scanner.Text()
		_, err := library.AddReader(firstName, lastName)
		if err != nil {
			fmt.Println("Произошла ошибка при создании читателя:", err)
			return
		}
		fmt.Println("Создан читатель")

		//Поиск читателя
	case 9:
		fmt.Println("Введите номер Читателя: ")
		scanner.Scan()
		idUser, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Номер должен состоять из цифр!")
			return
		}
		reader, err := library.FindReaderByID(idUser)
		if err != nil {
			fmt.Println("Произошла ошибка при поиске читателя:", err)
			return
		}
		fmt.Println(reader)

		//Импорт читателей
	case 10:
		fmt.Println("Введите название файла: ")
		scanner.Scan()
		filename := scanner.Text()
		readers, err := storage.LoadReadersFromCSV(filename)
		if err != nil {
			fmt.Println("Ошибка при импорте списка", err)
			return
		}
		library.Readers = readers
		fmt.Println("Список читателей успешно импортирован!")

	//экспорт читателей
	case 11:
		fmt.Println("Введите название файла: ")
		scanner.Scan()
		filename := scanner.Text()
		err := storage.SaveReadersToCSV(filename, library.Readers)
		if err != nil {
			fmt.Println("Ошибка экспорта: ", err)
			return
		}
		fmt.Println("Читатели успешно экспортированы")

		//Вывод всех читателей
	case 12:
		readers := library.GetAllReaders()
		fmt.Println("Список : ")
		for _, reader := range readers {
			fmt.Println(reader)
		}
	}

}

func Run(lib *library.Library) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		printMenu()
		scanner.Scan()
		choice, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Нужно ввести число")
			continue
		}
		handlerChoice(choice, scanner, lib)
		if choice == 0 {
			err = storage.SaveBooksToCSV("books.csv", lib.Books)
			if err != nil {
				fmt.Println("Произошла ошибка экспорта базы:", err)
				return
			}
			break
		}
	}
}
