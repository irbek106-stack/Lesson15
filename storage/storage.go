package storage

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/BatrazG/simple-library/domain"
	"github.com/BatrazG/simple-library/library"
)

type Storable interface {
	Save() error
	Load() error
}

// Сохраняет срез книг в csv-файл
func SaveBooksToCSV(filename string, books []*domain.Book) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("не удалось создать файл %s: %w", filename, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	//Записываем заголовок
	headers := []string{"ID", "Название", "Автор", "Год", "Выдана", "ID читателя"}
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("не удалось записать заголовок: %w", err)
	}

	//Записываем данные книг
	for _, book := range books {
		var readerID string
		if book.ReaderID != nil {
			readerID = strconv.Itoa(*book.ReaderID)
		}
		record := []string{
			strconv.Itoa(book.ID),
			book.Title,
			book.Author,
			strconv.Itoa(book.Year),
			strconv.FormatBool(book.IsIssued),
			readerID,
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("не удалось записать список книги с ID %d: %w", book.ID, err)
		}
	}
	return nil
}

// LoadBooksFromCSV загружает список книг из csv
func LoadBooksFromCSV(filename string) ([]*domain.Book, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл %s: %w", filename, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	//Ожидаемое количество столбцов
	const expectedColumns = 6

	//Читаем заголовок отдельно, чтобы пропустить его
	//Заодно убедимся, что файл не пустой
	if _, err := reader.Read(); err != nil {
		if errors.Is(err, io.EOF) {
			//Файл пустой или содержит только заголовок
			//Это не ошибка
			return []*domain.Book{}, nil
		}
		return nil, fmt.Errorf("не удалось прочитать заголовок: %w", err)
	}

	var books []*domain.Book
	//Добавляем счетчик строк для более информативных логов
	var lineNum int

	//Читаем построчно
	//База данных может сильно разрастись
	for {
		lineNum++
		record, err := reader.Read()
		if err != nil {
			//Если мы достигла конца файла
			//-это нормальное завершение цикла
			if errors.Is(err, io.EOF) {
				break
			}
			//Любая другая ошибка при чтении является критической
			return nil, fmt.Errorf("ошибка чтения файла на строке %d: %w", lineNum, err)
		}

		//Проверяем на точное соответствие количества колонок
		if len(record) != expectedColumns {
			log.Printf("ПРЕДУПРЕЖДЕНИЕ: Пропущена строка %d, неверное количество колонок (ожидалось %d, получено %d)", lineNum, expectedColumns, len(record))
			continue
		}

		//Ошибки будем логировать
		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("ПРЕДУПРЕЖДЕНИЕ: Пропускаем строку %d, неверный формат ID: %v", lineNum, err)
			continue //Неверный формат ID, пропускаем
		}

		year, err := strconv.Atoi(record[3])
		if err != nil {
			log.Printf("ПРЕДУПРЕЖДЕНИЕ: Пропускаем строку %d, неверный формат года: %v", lineNum, err)
			continue
		}

		isIssued, err := strconv.ParseBool(record[4])
		if err != nil {
			log.Printf("ПРЕДУПРЕЖДЕНИЕ: Пропускаем строку %d, неверный формат поля 'Выдана': %v", lineNum, err)
			continue
		}

		// Если поле пустое, указатель должен быть nil. Иначе - парсим значение.
		var readerIDPtr *int
		if record[5] != "" {
			readerID, err := strconv.Atoi(record[5])
			if err != nil {
				log.Printf("ПРЕДУПРЕЖДЕНИЕ: Пропускаем строку %d, неверный формат ID читателя: %v", lineNum, err)
				continue
			}
			readerIDPtr = &readerID
		}

		book := domain.Book{
			ID:       id,
			Title:    record[1],
			Author:   record[2],
			Year:     year,
			IsIssued: isIssued,
			ReaderID: readerIDPtr, // Используем созданный указатель.
		}

		books = append(books, &book)
	}
	return books, nil
}

func SaveReaderToCSV(filename string, readers []*domain.Reader) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("ошибка создания файла %s: %w", filename, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	//Записываем заголовок
	headers := []string{"ID", "Имя", "Фамилия", "Активен"}
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("не удалось записать заголовок: %w", err)
	}

	for _, reader := range readers {
		id := strconv.Itoa(reader.ID)
		status := strconv.FormatBool(reader.IsActive)

		record := []string{
			id,
			reader.FirstName,
			reader.LastName,
			status,
		}

		if err := writer.Write(record); err != nil {
			return fmt.Errorf("не удалось саписать книгу с ID %d: %s", reader.ID, err)
		}
	}

	return nil
}

func LoadReadersFromCSV(filename string) ([]*domain.Reader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("произошла ошибка открытия файла %s", filename)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	//Ожидаемое количество столбцов
	const expectedColumns = 4

	//Читаем заголовок отдельно, чтобы пропустить его
	//Заодно убедимся, что файл не пустой
	if _, err := reader.Read(); err != nil {
		if errors.Is(err, io.EOF) {
			//Файл пустой или содержит только заголовок
			//Это не ошибка
			return []*domain.Reader{}, nil
		}
		return nil, fmt.Errorf("не удалось прочитать заголовок: %w", err)
	}

	var readers []*domain.Reader
	//Добавляем счетчик для более информативных логов
	var lineNum int

	// Читаем построчно
	for {
		lineNum++
		record, err := reader.Read()
		if err != nil {
			//Если мы достигла конца файла
			//-это нормальное завершение цикла
			if errors.Is(err, io.EOF) {
				break
			}
			//Любая другая ошибка при чтении является критической
			return nil, fmt.Errorf("ошибка чтения файла на строке %d: %w", lineNum, err)
		}

		//Проверяем точное соответствие количества колонок
		if len(record) != expectedColumns {
			log.Printf("ПРЕДУПРЕЖДЕНИЕ: Пропущена строка %d, неверное количество колонок (%d вместо %d)", lineNum, expectedColumns, len(record))
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("ПРЕДУПРЕЖДЕНИЕ: Пропускаем строку %d, неверный формат ID: %v", lineNum, err)
		}

		status, err := strconv.ParseBool(record[3])
		if err != nil {
			log.Printf("ПРЕДУПРЕЖДЕНИЕ: Пропускаем строку %d, неверный формат поля \"Активен\"", lineNum)
			continue
		}

		reader := domain.Reader{
			ID:        id,
			FirstName: record[1],
			LastName:  record[2],
			IsActive:  status,
		}

		readers = append(readers, &reader)
	}

	return readers, nil
}

//Работа с JSON

// storageData - это вспомогательная структура для упаковки всех данных
// библиотеки в один объект для удобной сериализации в JSON.
type storageData struct {
	Books   []*domain.Book   `json:"books"`
	Readers []*domain.Reader `json:"readers"`
}

// Сохраняет все данные библиотеки в JSON файл
func SaveLibraryToJSON(filePath string, lib *library.Library) error {
	//1. Создаем экземпляр нашей структуры-контейнера
	data := storageData{
		Books:   lib.Books,
		Readers: lib.Readers,
	}

	//2. Сериализируем данные в JSON с отступами для читаемости.
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	//3. Записываем JSON-данные в файл
	// 0644 - это стандартные права доступа к файлу (чтение/запись для владельца, чтение для остальных).
	return os.WriteFile(filePath, jsonData, 0644)
}

func LoadLibraryFromJSON(filePath string) (*library.Library, error) {
	//1. Читаем содержимое файла
	jsonData, err := os.ReadFile(filePath) //Если файл не открывается, возвращаем ошибку
	if err != nil {
		return nil, err
	}

	//2. Создаем переменную-приемник
	var data storageData

	//3. Десериализируем JSON в нашу переменную - приемник
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, err
	}

	//4. Создаем новый экземпляр библиотеки
	lib := library.New()
	lib.Books = data.Books
	lib.Readers = data.Readers

	//Важно!!!
	//Нужно обновить внутренние счетчики ID в библиотеке
	lib.UpdateIDs()

	return lib, nil
}
