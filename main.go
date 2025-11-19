package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BatrazG/simple-library/cmd/cli"
	"github.com/BatrazG/simple-library/library"
	"github.com/BatrazG/simple-library/storage"
)

func main() {

	/*myLibrary := library.New()

	myLibrary.Books, _ = storage.LoadBooksFromCSV("books.csv")
	myLibrary.Readers, _ = storage.LoadReadersFromCSV("readers.csv")*/

	//Выносим имя файла в константу
	const dbFile = "books.json"

	//Пытаемся загрузить библиотеку из файла
	myLibrary, err := storage.LoadLibraryFromJSON(dbFile)
	if err != nil {
		//Если файл не найден - это не ошибка
		//Создаем новую пустую библиотеку
		if os.IsNotExist(err) {
			fmt.Println("Файл данных не найден, создана новая библиотека")
			myLibrary = library.New()
		} else {
			//Если произошла другая ошибка, например JSON файл не корректен
			//завершаем работу
			log.Fatalf("Ошибка при загрузке библиотеки: %v", err)
		}
	} else {
		fmt.Println("Библиотека успешно загружена из файла")
	}

	cli.Run(myLibrary, dbFile)

}
