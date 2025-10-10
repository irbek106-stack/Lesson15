package main

import "fmt"



type Reader struct {
	ID   int
	FirstName string
	LastName string
	IsActive bool
}

type Book struct {
	ID       int
	Title    string
	Author   string
	Year 	int
	readerID *Reader
	IsIssued bool
}

func (r Reader) DisplayReader(){
	fmt.Printf("Читатель: %s %s (ID: %d)\n", r.FirstName, r.LastName, r.ID)
}

func (r *Reader) Deactivate(){
	r.IsActive = false
	fmt.Printf("Пользователь: %s %s деактивирован.\n", r.FirstName, r.LastName)	
}

func (r *Reader) Activate(){
	r.IsActive = true
	fmt.Printf("Пользователь: %s %s активирован.\n", r.FirstName, r.LastName)	
}

func (r Reader) String() string{
	status := ""
	if r.IsActive{
		status = "активен"

	} else {
		status = "не активен"
	}
	return fmt.Sprintf("Пользователь %s %s, ID: %d, %s", r.FirstName, r.LastName, r.ID, status)
}

func (b Book) String() string{
	return fmt.Sprintf("'%s' (%s, %d)", b.Title, b.Author, b.Year )
}

func (b *Book) IssueBook(reader *Reader){

	if b.IsIssued{
		fmt.Printf("Книга уже выдана\n")
	} 
	if !reader.IsActive{
		fmt.Printf("Пользователь %s %s неактивен\n", reader.FirstName, reader.LastName)
		return
	}
	b.IsIssued = true
	b.readerID = nil
	fmt.Printf("Книга %s была выдана\n", b.Title)
}

func (b *Book) ReturnBook(){
	if !b.IsIssued{
		fmt.Printf("Книга %s не была выдана\n", b.Title)
		return
	}
	b.IsIssued = false
	b.readerID = nil
	fmt.Printf("Книга %s возвращена\n", b.Title)
}