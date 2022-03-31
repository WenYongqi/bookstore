package dao

import (
	"bookstore/model"
	"fmt"
	"testing"
)

func TestBook(t *testing.T) {
	t.Run("", testGetBooks)
	t.Run("", testAddBook)
	t.Run("", TestDeleteBook)
	t.Run("", TestGetBookByID)
	t.Run("", TestUpdateBook)
	t.Run("", TestGetPageBooks)
	t.Run("", TestGetPageBooksByPrice)
}

func testGetBooks(t *testing.T) {
	books, err := GetBooks()
	if err != nil {
		t.Fatal(err)
	}
	for _, book := range *books {
		fmt.Println(book)
	}
}

func testAddBook(t *testing.T) {
	book := model.Book{
		Title: "三体",
		Author: "刘慈欣",
		Price: 100,
		Sales: 108,
		Stock: 999,
	}
	err := AddBook(&book)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteBook(t *testing.T) {
	err := DeleteBook(34)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetBookByID(t *testing.T) {
	book, err := GetBookByID(3)
	if book == nil || err != nil {
		t.Fatal(err)
	}
}

func TestUpdateBook(t *testing.T) {
	book := model.Book{
		ID: 14,
		Title: "三体",
		Author: "刘慈欣",
		Price: 100,
		Sales: 108,
		Stock: 999,
	}
	err := UpdateBook(&book)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetPageBooks(t *testing.T) {
	page, err := GetPageBooks(2)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(page)
	fmt.Println(page.Books)
}

func TestGetPageBooksByPrice(t *testing.T) {
	page, err := GetPageBooksByPrice(1, 99, 101)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", page)
	fmt.Println(page.Books)
}