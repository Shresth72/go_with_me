package main

import "fmt"

type Book struct {
  name string
  author string
  pages int
}

// Encapsulation
// Hiding sensitive data from users
func (b Book) printDetails() {
  fmt.Printf("Book %s was wirtten by %s.\n", b.name, b.author)
  fmt.Printf("It contains %d pages\n", b.pages)
}

func (b Book) PrintName() {
  fmt.Printf("Name of the book is %s\n", b.name)
}

// Polymorphism: Ability to use a single type in different forms
// Done using interfaces
type Printable interface {
  printDetails()
}

func printInfo(p Printable) {
  p.printDetails()
}

// Inheritance
// When a class acquires the properties of its superclass then we can say it is inheritance. Subclass class the terms used for the class which acquire properties.
type EBook struct {
  Book
  fileSize int
}

func (e EBook) printDetails() {
  fmt.Printf("E-Book %s was written by %s.\n", e.name, e.author)
	fmt.Printf("It contains %d pages and the file size is %dMB.\n", e.pages, e.fileSize)
}

// Abstraction: Hiding implementation details and showing only essential info
// The Printable interface is an example of Abstraction

func main() {
  book1 := Book{"Monster", "Shine", 130}
	ebook1 := EBook{Book{"Go Programming", "Alice", 200}, 5}

	// Encapsulation and Polymorphism
	printInfo(book1)
	printInfo(ebook1)

	// Direct access to specific method
	ebook1.printDetails()
}
