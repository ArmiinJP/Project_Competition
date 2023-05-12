package main

import (
	"fmt"
	"strings"
)


type Library struct {
	capacity int
	numExist int
	books []string
	borrow map[string]string

}

func NewLibrary(cap int) *Library {
	return &Library{
		capacity: cap,
		numExist: 0,
		books: []string{},
		borrow: map[string]string{},
	}
}

func (library *Library) AddBook(name string) string {

	name = strings.ToLower(name)
	for _, v := range(library.books){
		if v == name{
			return "The book is already in the library"
		}
	}

	if library.numExist == library.capacity{
		return "Not enough capacity"
	}

	library.books = append(library.books, name)
	library.numExist += 1
	
	return "OK"
}

func (library *Library) BorrowBook(bookName, personName string) string {
	bookName = strings.ToLower(bookName)
	//personName = strings.ToLower(personName)

	for bName, pName := range library.borrow{
		if bName == bookName{
			return fmt.Sprintf("The book is already borrowed by %s", pName)
		}
	}

	check := false
	for _, bName := range library.books{
		if bName == bookName{
			check = true
		}
	}

	if check {
		library.borrow[bookName] = personName
		return "OK"
	} else {
		return "The book is not defined in the library"
	}
}

func (library *Library) ReturnBook(bookName string) string {
	bookName = strings.ToLower(bookName)

	for bName := range library.borrow{
		if bName == bookName{
			delete(library.borrow, bName)
			return "OK"
		}
	}

	check := false
	for _, bName := range library.books{
		if bName == bookName{
			check = true
		}
	}
	if check {
		return "The book has not been borrowed"
	} else {
		return "The book is not defined in the library"
	}

}


func main(){
	l1 := NewLibrary(3)
	fmt.Println(l1.AddBook("b1"))
	fmt.Println(l1.AddBook("b2"))
	fmt.Println(l1.AddBook("b2"))
	fmt.Println(l1.AddBook("b3"))
	fmt.Println(l1.AddBook("b1"))
	fmt.Println(l1.AddBook("b4"))
	
	fmt.Println(l1.BorrowBook("b1", "p1"))
	fmt.Println(l1.BorrowBook("b2", "p2"))
	fmt.Println(l1.BorrowBook("b1", "p2"))
	fmt.Println(l1.BorrowBook("b4", "p2"))
	fmt.Println(l1.BorrowBook("b3", "p3"))

	fmt.Println(l1.ReturnBook("b1"))
	fmt.Println(l1.BorrowBook("b1", "p2"))
	fmt.Println(l1.BorrowBook("b1", "p3"))



}