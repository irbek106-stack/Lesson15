package notifications

import "fmt"


// Notifier - это наш контракт. Любая структура, которая хочет
// быть "уведомителем", должна иметь метод Notify.
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

// Реализуем интерфейс Notifier для EmailNotifier.
// Теперь EmailNotifier неявно является Notifier'ом.
func (e EmailNotifer) Notify(message string) {
	fmt.Printf("Отправляю email на адрес %s: %s\n", e.EmailAdress, message)
}

// Реализуем тот же интерфейс Notifier для SMSNotifier.
func (s SMSNotifer) Notify(message string) {
	fmt.Printf("Отправляю СМС на номер %s: %s\n", s.PhoneNumber, message)
}

