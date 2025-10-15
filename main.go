package main

import "fmt"

func main() {
	fmt.Println("������ ������� ���������� �����������")

	//1. ������� ��������� ����������
	myLibrary := &Library{} //������ ���������� ������ � ������

	fmt.Println("��������� ����������")
	//2. ��������� ���������
	myLibrary.AddReader("������", "�������")
	myLibrary.AddReader("������", "�������")

	//3. ��������� �����
	myLibrary.AddBook("1984", "������ ������", 1949)
	myLibrary.AddBook("������ � ���������", "������ ��������", 1967)

	fmt.Println("\n---���������� ������ � ������---")
	fmt.Println("���������� ���������:", len(myLibrary.Readers))
	fmt.Println("���������� ����:", len(myLibrary.Books))

	//������ 16. ���������
	fmt.Println("---��������� ������ ����---")
	//������ ����� 1 �������� 1
	err := myLibrary.IssueBookToReader(1, 1)
	if err != nil {
		fmt.Println("������ ������", err)
	}

	//��������� ������ ����� ����� ������
	book, _ := myLibrary.FindBookByID(1)
	if book != nil {
		fmt.Println("������ ����� ����� ������:", book)
	}

	//������� ������ �������������� �����
	err = myLibrary.IssueBookToReader(99, 1)
	if err != nil {
		fmt.Println("��������� ������:", err)
	}

}






























































































































































