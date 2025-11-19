package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/BatrazG/simple-library/library"
	"github.com/BatrazG/simple-library/storage"
)

// Run запускает главный цикл консольного приложения.
// Он принимает сервис библиотеки как зависимость.
func Run(lib *library.Library, dbPath string) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printMenu()

		scanner.Scan()
		inputText := scanner.Text()

		choice, err := strconv.Atoi(strings.TrimSpace(inputText))
		if err != nil {
			fmt.Println("Ошибка: Пожалуйста, введите число.")
			continue
		}

		handleChoice(choice, lib, scanner) // Передаем сервис и сканер в обработчик

		if choice == 0 {
			fmt.Println("Сохранение данных и выход.")
			if err := storage.SaveLibraryToJSON("books.json", lib); err != nil {
				fmt.Println("Произошла ошибка сохранения списка книг:", err)
			}
			return // Выходим из цикла, если выбрали выход
		}
	}
}

// printMenu отвечает за вывод в консоль пользовательского меню
func printMenu() {
	//Вывод меню
	fmt.Println("")
	fmt.Println("----------------")
	fmt.Println("Главное меню:")
	fmt.Println("1. Поиск книги по названию")
	fmt.Println("2. Поиск книги по номеру")
	fmt.Println("3. Выдать книгу")
	fmt.Println("4. Вернуть книгу")
	fmt.Println("5. Поиск читателя по номеру")
	fmt.Println("6. Показать список книг")
	fmt.Println("7. Экспорт списка книг")
	fmt.Println("8. Импорт списка книг")
	fmt.Println("9. Добавление новой книги")
	fmt.Println("10 Добавление нового читателя")
	fmt.Println("11 Экспорт списка читателей")
	fmt.Println("12 Импорт списка читателей")
	fmt.Println("0. Выход")
	fmt.Println("Выберите пункт меню:")
}

// handlerChoice обрабатывает выбор пользователем пункта меню
func handleChoice(choice int, lib *library.Library, scanner *bufio.Scanner) {
	switch choice {
	case 1: //поиск книги по названию
		fmt.Println("---Введите название книги:---")
		scanner.Scan()
		title := scanner.Text()
		foundBooks, err := lib.FindBookByTitle(title)
		if err != nil {
			fmt.Println("Произошла ошибка поиска: ", err)
			return
		}

		if len(foundBooks) == 0 {
			fmt.Printf("Совпадений с названием %s не найдено\n", title)
		} else {
			for _, book := range foundBooks {
				fmt.Println(book)
			}
		}
	case 2: //Поиск книги по ID
		fmt.Println("---Введите номер книги:---")
		scanner.Scan()
		bookID, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Номер книги должен быть числом")
			return
		}
		foundBook, err := lib.FindBookByID(bookID)
		if err != nil {
			fmt.Println("Произошла ошибка поиска: ", err)
		} else {
			fmt.Printf("Книга с номером %d %s\n:", bookID, foundBook)
		}
	case 3: //Выдача книги читателю
		fmt.Println("---Введите ID книги:---")
		scanner.Scan()
		bookID, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Номер книги должен быть числом")
			return
		}
		fmt.Println("---Введите ID читателя---")
		scanner.Scan()
		readerID, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Номер читателя должен быть числом")
			return
		}
		//Выдыча книги читателю
		err = lib.IssueBookToReader(bookID, readerID)
		if err != nil {
			fmt.Println("Произошла ошибка выдачи книги:", err)
			return
		}
		fmt.Printf("Книга с ID %d успешно выдана читателю с ID %d", bookID, readerID)
	case 4: //Возврат книги
		fmt.Println("---Введите ID книги:---")
		scanner.Scan()
		bookID, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Номер книги должен быть числом")
			return
		}
		err = lib.ReturnBook(bookID)
		if err != nil {
			fmt.Println("Произошла при возврате книги:", err)
			return
		}
		fmt.Printf("Книга с ID %d успешно возвращена в библиотеку", bookID)
	case 5: //Поиск читателя по ID
		fmt.Println("---Посик читателя по номеру---")
		fmt.Println("Введите номер читателя")
		scanner.Scan()
		readerID, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Номер читателя должен быть числом")
			return
		}
		reader, err := lib.FindReaderByID(readerID)
		if err != nil {
			fmt.Println("Произошла ошибка поиска читателя:", err)
			return
		}
		fmt.Println(reader)
	case 6: //Вывод на экран списка всех книг
		fmt.Println("\n---Список всех книг библиотеки---")
		books := lib.GetAllBooks()
		if len(books) == 0 {
			fmt.Println("Библиотечный фонд пуст")
			return
		}
		for i, book := range books {
			fmt.Println(i+1, "\t", book)
		}
	case 7: //Экспорт списка книг в csv
		fmt.Println("Введите название файла для экспорта в формате <название.csv>:")
		scanner.Scan()
		filename := scanner.Text()
		if err := storage.SaveBooksToCSV(filename, lib.Books); err != nil {
			fmt.Println("Произошла ошибка экспорта:", err)
			return
		}
		fmt.Printf("Список книг успешно выгружен в файл %s", filename)
	case 8: //Импорт книг из csv
		fmt.Println("Введите название файла для импорта в формате <название.csv>:")
		scanner.Scan()
		filename := scanner.Text()
		loadedBooks, err := storage.LoadBooksFromCSV(filename)
		if err != nil {
			fmt.Println("Ошибка импорта:", err)
			return
		}
		lib.Books = loadedBooks
		fmt.Printf("Список книгу успешно импортирован из файла %s\n", filename)
	case 9: //Новая книга
		fmt.Println("Введите название книги")
		scanner.Scan()
		title := scanner.Text()
		fmt.Println("Введите автора")
		scanner.Scan()
		author := scanner.Text
		fmt.Println("Введите год издания")
		scanner.Scan()
		year, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Год должен состояь из цифр.")
			return
		}
		lib.AddBook(title, author(), year)
	case 10: //Новый читатель
		fmt.Println("Введите имя")
		scanner.Scan()
		firstName := scanner.Text()
		fmt.Println("Введите фамилию")
		scanner.Scan()
		lastName := scanner.Text()

		if _, err := lib.AddReader(firstName, lastName); err != nil {
			fmt.Printf("Ошибка регистрации: %v", err)
			return
		}
	case 11: //Экспорт списка читателей
		fmt.Println("Введите название файла для экспорта в формате <название>.csv")
		scanner.Scan()
		filename := scanner.Text()
		if err := storage.SaveReaderToCSV(filename, lib.Readers); err != nil {
			fmt.Printf("Произошла ошибка экспорта: %v", err)
			return
		}
		fmt.Printf("Список читателей успешно выгружен в файл %s", filename)
	case 12: //Экспорт списка читателей
		fmt.Println("Введите название файла для экспорта читателей в формате <название>.csv")
		scanner.Scan()
		filename := scanner.Text()
		loadedReaders, err := storage.LoadReadersFromCSV(filename)
		if err != nil {
			fmt.Println("Ошибка импорта: ", err)
			return
		}
		lib.Readers = loadedReaders
		fmt.Printf("Список читателей успешно импортирован из файла %s\n", filename)
	case 0:
		fmt.Println("Всего доброго!")
	} //switch
}

/*
//Создаем меню консольного приложения
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Добро пожаловать в")
		fmt.Println("\033[1m" + "Simple library" + "\033[0m") //С помощью управляющих символов делаем строку жирной
		fmt.Println()



		//Считываем ввод пользователя
		scanner.Scan()
		inputText := scanner.Text()

		//Преобразуем строку в число
		choice, err := strconv.Atoi(inputText)

		//Проверяем на ошибку(если ввели не число)
		if err != nil {
			fmt.Println("Ошибка: пожалуйста, введите число от 1 до 8")
			continue
		}

		//Выбираем действие
		switch choice {
		case 1: //поиск книги по названию
			fmt.Println("Введите название книги:")
			scanner.Scan()
			title := scanner.Text()
			foundBooks, err := myLibrary.FindBookByTitle(title)
			if err != nil {
				fmt.Println("Произошла ошибка: ", err)
			} else if len(foundBooks) == 0 {
				fmt.Printf("Совпадений с названием %s не найдено\n", title)
			} else {
				for _, book := range foundBooks {
					fmt.Println(book)
				}
			}
		case 2:
			fmt.Println("Введите номер книги:")
			scanner.Scan()
			bookID, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Номер книги должен быть числом")
				continue
			}
			foundBook, err := myLibrary.FindBookByID(bookID)
			if err != nil {
				fmt.Println("Произошла ошибка: ", err)
			} else {
				fmt.Printf("Книга с номером %d %s\n:", bookID, foundBook)
			}
		}
	}
*/
