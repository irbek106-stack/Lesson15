package library

import (
	"errors"
	"fmt"
	"libary-app/domain"
	"strings"
)


func New() *Library{


}   


func (b *domain.Book) IssueBook(reader *domain.Reader) error {
	if b.IsIssued {
		return fmt.Errorf("книга '%s' уже выдана", b.Title)
	}
	if !reader.IsActive {
		return fmt.Errorf("читатель %s %s не активен", reader.FirstName, reader.LastName)
	}
	b.IsIssued = true
	b.ReaderID = &reader.ID
	return nil
}

func (b *domain.Book) ReturnBook() error{
	if !b.IsIssued {
		return fmt.Errorf("книга '%s' и так в библиотеке", b.Title)
	}
	b.IsIssued = false
	b.ReaderID = nil
	return nil
}


func (lib *Library) AddReader(firstName, lastName string) (*domain.Reader, error) {
	cleanedFirstName := strings.TrimSpace(firstName)
	cleanedLastName := strings.TrimSpace(lastName)

	if cleanedFirstName == "" || cleanedLastName == "" {
		return nil, errors.New("фамилия и имя не могут быть пустыми")
	}
	
	newReader := &domain.Reader{
		ID:        lib.lastReaderID,
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  true,
	}

	lib.Readers = append(lib.Readers, newReader)

	fmt.Printf("Зарегистрирован новый читатель: %s %s \n", firstName, lastName)
	return newReader, nil
}

func (lib *Library) AddBook(title, author string, year int) *domain.Book {
	lib.lastBookID++

	newBook := &domain.Book{
		ID:       lib.lastBookID,
		Title:    title,
		Author:   author,
		Year:     year,
		IsIssued: false,
	}

	lib.Books = append(lib.Books, newBook)

	fmt.Printf("Добавлена новая книга: %s\n", newBook)
	return newBook
}

func (lib *Library) FindBookByID(id int) (*domain.Book, error) {
	for _, book := range lib.Books {
		if book.ID == id {
			return book, nil
		}
	}
	return nil, fmt.Errorf("книга с ID %d не найдена", id)
}

func (lib *Library) FindReaderByID(id int) (*domain.Reader, error) {
	for _, reader := range lib.Readers {
		if reader.ID == id {
			return reader, nil
		}
	}

	return nil, fmt.Errorf("читатель с ID %d не найден", id)
}

func (lib *Library) IssueBookToReader(bookID, readerID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}

	reader, err := lib.FindReaderByID(readerID)
	if err != nil {
		return err
	}

	err = book.IssueBook(reader)
	if err != nil {
		return err
	}

	book.IssueBook(reader)
	return nil
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
	return err
}


func (lib *Library) GetAllBooks() []*domain.Book {
	return lib.Books
}
