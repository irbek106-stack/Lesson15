package main

import "fmt"

type Notifer interface {
	Notify(message string)
}

type EmailNotifer struct {
	EmailAdress string
}

type SMSNotifer struct {
	PhoneNumber string
}

func (e EmailNotifer) Notify(message string) {
	fmt.Printf("Сообщение придет на эл. почту %s: %s\n", e.EmailAdress, message)
}

func (s SMSNotifer) Notify(message string) {
	fmt.Printf("Сообщение придет на личный номер %s: %s\n", s.PhoneNumber, message)
}