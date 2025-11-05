package cli

import (
	"bufio"
	"fmt"
	"library-app/library"
	"strconv"
)

func printMenu() {

	fmt.Println("--SIMPLE-LIBRARY--")
	fmt.Println("----------------------------------------")
	fmt.Println("-MENU-")
	fmt.Println("[1] Добавление книги")
	fmt.Println("[2] Выдача книги")
	fmt.Println("[3] Возврат книги")
	fmt.Println("[4] Поиск книги")
	fmt.Println("[5] Импорт книги")
	fmt.Println("[6] Экспорт книги")
	fmt.Println("[7] Экспорт книги")
	fmt.Println("[8] Добавление читателя")
	fmt.Println("[9] Поиск читателя")
	fmt.Println("[10] Импорт читателя")
	fmt.Println("[11] Экспорт читателя")

}

func handlerChoice(choice int, scanner *bufio.Scanner, library library.Library) {
	switch choice {
	case 1:
		fmt.Println("Введите название книги: ")
		scanner.Scan()
		title := scanner.Text()

		fmt.Println("Введите автора книги: ")
		scanner.Scan()
		author := scanner.Text()

		fmt.Println("Введите название книги: ")
		scanner.Scan()
		year, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Год должен состоять из цифр")
			return
		}

		if _, err := library.AddBook(title, author, year); err != nil {
			fmt.Printf("Произошла ошибка при добавление книги:%v", err)
		}

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
		}

	}

}
