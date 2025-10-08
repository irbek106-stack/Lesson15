package main


type Book struct {
	ID int
	Title string
	Author string
	IsIssued bool
}

type Reader struct {
	ID int
	Name string
}