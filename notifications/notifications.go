package notifications

import "fmt"

type Notifer interface {
	Notify(message string)
}

// EmailNotifier - конкретная реализация уведомителя через Email.
type EmailNotifer struct {
	EmailAdress string
}

// SMSNotifier - конкретная реализация уведомителя через SMS.
type SMSNotifer struct {
	PhoneNumber string
}

func (e EmailNotifer) Notify(message string) {
	fmt.Printf("Отправляю email на адрес %s: %s\n", e.EmailAdress, message)
}

func (s SMSNotifer) Notify(message string) {
	fmt.Printf("Отправляю СМС на номер %s: %s\n", s.PhoneNumber, message)
}
