package main

import "fmt"

// Notifier - ��� ��� ��������. ����� ���������, ������� �����
// ���� "������������", ������ ����� ����� Notify.
type Notifer interface {
	Notify(message string)
}

// EmailNotifier - ���������� ���������� ����������� ����� Email.
type EmailNotifer struct {
	EmailAdress string
}

// SMSNotifier - ���������� ���������� ����������� ����� SMS.
type SMSNotifer struct {
	PhoneNumber string
}

// ��������� ��������� Notifier ��� EmailNotifier.
// ������ EmailNotifier ������ �������� Notifier'��.
func (e EmailNotifer) Notify(message string) {
	fmt.Printf("��������� email �� ����� %s: %s\n", e.EmailAdress, message)
}

// ��������� ��� �� ��������� Notifier ��� SMSNotifier.
func (s SMSNotifer) Notify(message string) {
	fmt.Printf("��������� ��� �� ����� %s: %s\n", s.PhoneNumber, message)
}