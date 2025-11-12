//сохраняет срез книг в csv-файл
//create (file,err) defer
//newwriter defer flush
//headers err writer.write
//books' data

package storage

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"library-app/domain"
)

type Storable interface {
	Save() error
	Load() error
}

// Экспорт списка книг в CSV
func SaveBooksToCSV(filename string, books []*domain.Book) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("не удалось создать файл %s: %w", filename, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	//Записываем заголовок
	headers := []string{"ID", "Название", "Автор", "Год", "Статус", "ID читателя"}
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("не удалось записать заголовок: %w", err)
	}

	//Записываем данные книг
	for _, book := range books {
		var status string = ""
		if book.IsIssued {
			status = "1"
		} else {
			status = "0"
		}
		var readerID string
		if book.ReaderID == nil {
			readerID = ""
		} else {
			readerID = strconv.Itoa(*book.ReaderID)
		}
		record := []string{
			strconv.Itoa(book.ID),
			book.Title,
			book.Author,
			strconv.Itoa(book.Year),
			status,
			readerID,
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("не удалось записать список книги с ID %d: %w", book.ID, err)
		}
	}
	return nil
}

// Импорт списка книг из CSV
func LoadBooksFromCSV(filename string) ([]*domain.Book, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл %s", filename)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать данные из файла: %w", err)
	}

	if len(records) < 2 {
		return []*domain.Book{}, nil
	}

	var books []*domain.Book
	for _, record := range records[1:] {
		if len(record) < 5 {
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			continue
		}

		year, err := strconv.Atoi(record[3])
		if err != nil {
			continue
		}

		book := domain.Book{
			ID:     id,
			Title:  record[1],
			Author: record[2],
			Year:   year,
		}

		books = append(books, &book)
	}

	return books, nil
}

// Экспорт читателей
func SaveReadersToCSV(filename string, readers []*domain.Reader) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("не удалось создать файл %s: %w", filename, err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	//Записываем заголовок
	headers := []string{"ID", "Имя", "Фамилия", "Статус"}
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("не удалось записать заголовок: %w", err)
	}

	//Записываем данные книг
	for _, reader := range readers {
		var status string = ""
		if reader.IsActive {
			status = "1"
		} else {
			status = "0"
		}
		record := []string{
			strconv.Itoa(reader.ID),
			reader.FirstName,
			reader.LastName,
			status,
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("не удалось записать список читателей с ID %d: %w", reader.ID, err)
		}
	}
	return nil
}

// Импорт списка читателей из CSV

func LoadReadersFromCSV(filename string) ([]*domain.Reader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл %s", filename)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать данные из файла: %w", err)
	}

	if len(records) < 2 {
		return []*domain.Reader{}, nil
	}

	var readers []*domain.Reader
	for _, record := range records[1:] {
		if len(record) < 4 {
			continue
		}
		id, err := strconv.Atoi(record[0])
		if err != nil {
			continue
		}

		status, err := strconv.ParseBool(record[3])
		if err != nil {
			continue
		}

		read := domain.Reader{
			ID:        id,
			FirstName: record[1],
			LastName:  record[2],
			IsActive:  status,
		}

		readers = append(readers, &read)
	}

	fmt.Println("Список читателей: ")
	for _, read := range readers {
		fmt.Println(read)
	}
	return readers, nil

}
