package main

import "fmt"

type book struct {
	Author  string
	Title   string
	Year    uint16
	Price   float32
	Publish string
}

const semyon int = 5

var (
	i      int
	b1, b2 uint16
	books  [semyon]book
)

func input() {
	for i := 0; i < semyon; i++ {
		fmt.Println("book[", i+1, "]\n")
		fmt.Print("Author: ")
		fmt.Scan(&books[i].Author)
		fmt.Print("Title: ")
		fmt.Scan(&books[i].Title)
		fmt.Print("Year: ")
		fmt.Scan(&books[i].Year)
		fmt.Print("Price: ")
		fmt.Scan(&books[i].Price)
		fmt.Print("Publish: ")
		fmt.Scan(&books[i].Publish)
	}
}

func search() {
	fmt.Print("Enter first year bound: ")
	fmt.Scan(&b1)
	fmt.Print("Enter second year bound: ")
	fmt.Scan(&b2)
	for i := 0; i < semyon; i++ {
		if books[i].Year >= b1 && books[i].Year <= b2 {
			fmt.Println(books[i].Title)
		}
	}

}

func main() {
	input()
	search()
}
