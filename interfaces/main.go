package main

import (
	"example/go-gin/interfaces/db"
	"fmt"
)

type Article struct {
	Title  string
	Author string
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

type Stringer interface {
	String() string
}

func (b Book) String() string {
	return fmt.Sprintf("The %q book was written by %s", b.Title, b.Author)
}

func (a Article) String() string {
	return fmt.Sprintf("The book is: %s", a.Title)
}

func main() {
	a := Article{
		Title:  "Understanding Go",
		Author: "Ruben V",
	}

	b := Book{
		Title:  "The organized Mind",
		Author: "Daniel J. Levitin",
		Pages:  100,
	}

	Print(a)
	Print(b)

	c := Circle{Radius: 10}
	s := Square{Height: 10, Width: 5}

	l := Less(c, s)
	fmt.Printf("%+v is the smallest\n", l)

	PrintArea(s)

	PriceCheck(1)
}

func Print(s Stringer) {
	fmt.Println(s)
}

func PriceCheck(itemId int) (float64, bool) {
	item := db.LoadItem(itemId)
	if item == nil {
		return 0, false
	}
	return item.Price, true
}
