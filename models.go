package main

import (
	"errors"
	"fmt"
	"strings"
)

type Book struct {
	ID       int
	Title    string
	Author   string
	Year     int
	IsIssued bool
	ReaderID *int //ID читателя, который взял книгу
}

func (b Book) String() string {
	status := "в библиотеке"
	if b.IsIssued && b.ReaderID != nil {
		status = fmt.Sprintf("на руках у читателя с ID %d", *b.ReaderID)
	}
	return fmt.Sprintf("%s (%s, %d), статус: %s", b.Title, b.Author, b.Year, status)
}

// IssueBook выдает книгу читателю. Теперь возвращает ошибку.
func (b *Book) IssueBook(reader *Reader) error {
	if b.IsIssued {
		//Теперь возвращаем ошибку, а не печатаем в консоль
		return fmt.Errorf("книга '%s' уже выдана", b.Title)
	}
	if !reader.IsActive {
		return fmt.Errorf("читатель %s %s не активен и не может получить книгу", reader.FirstName, reader.LastName)
	}
	b.IsIssued = true
	b.ReaderID = &reader.ID

	return nil //Книга успешно выдана
}

// ReturnBook возвращает книгу в библиотеку
func (b *Book) ReturnBook() error {
	if !b.IsIssued {
		return fmt.Errorf("книга '%s' и так в библиотеке", b.Title)
	}
	b.IsIssued = false
	b.ReaderID = nil
	return nil
}

type Reader struct {
	ID        int
	FirstName string
	LastName  string
	IsActive  bool
}

func (r Reader) String() string {
	status := ""
	if r.IsActive {
		status = "активен"
	} else {
		status = "не активен"
	}
	return fmt.Sprintf("Пользователь %s %s, № %d, пользователь %s", r.FirstName, r.LastName, r.ID, status)
}

// Deactivate делает пользователя неактивным
func (r *Reader) Deactivate() {
	r.IsActive = false
}

func (r *Reader) Activate() {
	r.IsActive = true
}

// Library - наша центральная структура-агрегатор
type Library struct {
	Books   []*Book
	Readers []*Reader

	//Счетчики для генерации уникальных ID
	lastBookID   int
	lastReaderID int
}

func (lib *Library) AddReader(firstName, lastName string) (*Reader, error) {
	cleanedFirstName := strings.TrimSpace(firstName)
	cleanedLastName := strings.TrimSpace(lastName)

	if cleanedFirstName == "" || cleanedLastName == "" {
		return nil, errors.New("фамилия и имя не могут быть пустыми")
	}

	lib.lastReaderID++

	//Создаем нового читателя
	newReader := &Reader{
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
func (lib *Library) AddBook(title, author string, year int) *Book {
	lib.lastBookID++

	//Создаем новую книгу
	newBook := &Book{
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
func (lib *Library) FindBookByID(id int) (*Book, error) {
	for _, book := range lib.Books {
		if book.ID == id {
			return book, nil
		}
	}

	return nil, fmt.Errorf("книга с ID %d не найдена в библиотеке", id)
}

// FindReaderByID ищет читателя по его уникальному ID
func (lib *Library) FindReaderByID(id int) (*Reader, error) {
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

func (lib *Library) ReturnBook(bookID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}
	err = book.ReturnBook()
	if err != nil {
		return err
	}
	return nil
}

// GetAllBooksПоказывает все книги в библиотеке
func (lib *Library) GetAllBooks() []*Book {
	return lib.Books
}
