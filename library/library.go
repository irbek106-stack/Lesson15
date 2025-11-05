package library

import(
	"library-app/domain"
	"strings"
	"fmt"
	"errors"
	// "strings"
)


// Library - наша центральная структура-агрегатор
type Library struct {
	Books   []*domain.Book
	Readers []*domain.Reader

	//Счетчики для генерации уникальных ID
	lastBookID   int
	lastReaderID int
}


func (lib *Library) AddReader(firstName, lastName string) (*domain.Reader,error) {
	cleanedFirstName := strings.TrimSpace(firstName)
	cleanedLastName := strings.TrimSpace(lastName)

	if cleanedFirstName == "" || cleanedLastName == ""{
		return nil, errors.New("фамилия и имя не могут быть пустыми")
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
func (lib *Library) AddBook(title, author string, year int) (*domain.Book, error) {
	// Проверка дубля
    for _, b := range lib.Books {
        if b.Title == title && b.Author == author {
            return nil, fmt.Errorf("книга '%s' авторства '%s' уже существует", title, author)
        }
    }
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


	return newBook, nil
}



// FindBookByID ищет книгу по ее уникальному ID
func (lib *Library) FindBookByID(id int) (*domain.Book, error) {
	for _, book := range lib.Books {
		if book.ID == id {
			return book, nil
		}
	}
	return nil, fmt.Errorf("книга с ID %d не найдена", id)
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

func (lib *Library) IssueBookToReader(bookID, readerID int) error {

	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err 
	}
	
	if book.IsIssued{
		return fmt.Errorf("Книга '%s' уже выдана", book.Title)
	}

	readerExists := false
	for _, reader := range lib.Readers{		
		if reader.ID == readerID{
			readerExists = true
			break
		}
	}
	if !readerExists{
		return fmt.Errorf("Читатель с ID %d не найден", readerID)
	}

	book.IsIssued = true
	book.ReaderID = &readerID
	return nil
}


func (lib *Library) ReturnBook(bookID int) error{
	book, err := lib.FindBookByID(bookID)
	if err!= nil{
		return err
	}
	err = book.ReturnBook()
	if err!= nil{
		return err
	}
	return book.ReturnBook()
}


func (lib *Library) GetAllBooks() []*domain.Book{
	return lib.Books

}


func New() *Library {
	return &Library{
		Books:   []*domain.Book{},
		Readers: []*domain.Reader{},
	}
}