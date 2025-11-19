package library

import (
	"errors"
	"fmt"
	"strings"

	"github.com/BatrazG/simple-library/domain"
)

// Library - наша центральная структура-агрегатор
type Library struct {
	Books   []*domain.Book
	Readers []*domain.Reader

	//Счетчики для генерации уникальных ID
	lastBookID   int
	lastReaderID int
}

func New() *Library {
	return &Library{
		Books:   make([]*domain.Book, 0),
		Readers: make([]*domain.Reader, 0),
	}
}

func (lib *Library) AddReader(firstName, lastName string) (*domain.Reader, error) {
	cleanFirstName := strings.TrimSpace(firstName)
	cleanLastName := strings.TrimSpace(lastName)
	if cleanFirstName == "" || cleanLastName == "" {
		return nil, errors.New("фамимлия и имя не могут быть пустыми")
	}

	lib.lastReaderID++
	//Создаем нового читателя
	newReader := &domain.Reader{
		ID:        lib.lastReaderID,
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  true, //Новый читатель всегда активный
	}
	//Добавляем читателя в срез
	lib.Readers = append(lib.Readers, newReader)
	return newReader, nil
}

// AddBook добавляет новую книгу в библиотеку
func (lib *Library) AddBook(title, author string, year int) *domain.Book {
	lib.lastBookID++

	//Создаем новую книгу
	newBook := &domain.Book{
		ID:       lib.lastBookID,
		Title:    title,
		Author:   author,
		Year:     year,
		IsIssued: false, //Новая книга всегда в наличии
	}

	//Добавляем новую книгу в библиотеку
	lib.Books = append(lib.Books, newBook)
	return newBook
}

// FindBookByID ищет книгу по ее уникальному ID
func (lib *Library) FindBookByID(id int) (*domain.Book, error) {
	for _, book := range lib.Books {
		if book.ID == id {
			return book, nil
		}
	}

	return nil, fmt.Errorf("книга с ID %d не найдена в библиотеке", id)
}

func (lib *Library) FindBookByTitle(title string) ([]*domain.Book, error) {
	if lib == nil || len(lib.Books) == 0 {
		return nil, fmt.Errorf("библиотека не создана или не содержит книг")
	}

	cleanedTitle := strings.ToLower(strings.TrimSpace(title))

	books := []*domain.Book{}

	if cleanedTitle == "" {
		return books, errors.New("название книги не может быть пустым")
	}

	for _, book := range lib.Books {
		if book == nil {
			continue
		}
		cleanedBook := strings.ToLower(strings.TrimSpace(book.Title))

		if cleanedBook == cleanedTitle {
			books = append(books, book)
		}
	}
	// Также для strings.EqualFold(book.Title, title)
	//и вообще для проверки пользовательского ввода
	//-сравнение строк без учета регистра

	return books, nil
}

// FindReaderByID ищет читателя по его уникальному ID
func (lib *Library) FindReaderByID(id int) (*domain.Reader, error) {
	for _, reader := range lib.Readers {
		if reader.ID == id {
			return reader, nil
		}
	}

	return nil, fmt.Errorf("читатель с ID %d не найден", id)
}

// IssueBookToReader - основной публичный метод для выдачи книги
func (lib *Library) IssueBookToReader(bookID, readerID int) error {
	//1. Найти книгу
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}

	//2. Найти читателя
	reader, err := lib.FindReaderByID(readerID)
	if err != nil {
		return err
	}

	//Вызываем обновленный метод и ПРОВЕРЯЕМ ОШИБКУ
	err = book.IssueBook(reader)
	if err != nil {
		return err
	}
	return nil //Все 3 шага прошли успешно
}

func (lib Library) ReturnBook(bookID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}
	return book.ReturnBook()
}

func (lib *Library) GetAllBooks() []*domain.Book {
	//Просто возвращаем срез
	return lib.Books
}

// UpdateIDs находит максимальные ID среди книг и читателей
// и устанавливает внутренние счетчики библиотеки на эти значения
// Это необходимо вызывать после загрузки данных из файла
func (lib *Library) UpdateIDs() {
	maxBookID := 0
	for _, book := range lib.Books {
		if book.ID > maxBookID {
			maxBookID = book.ID
		}
	}
	lib.lastBookID = maxBookID

	maxReaderID := 0
	for _, reader := range lib.Readers {
		if reader.ID > maxReaderID {
			maxReaderID = reader.ID
		}
	}
	lib.lastReaderID = maxReaderID
}
