package gift

import "fmt"

type Book struct {
	Title  string
	Author string
}

func (b *Book) Present() {
	fmt.Printf("Книга: %s, Автор: %s\n", b.Title, b.Author)
}

type Chocolate struct {
	Name   string
	Origin string
}

func (c *Chocolate) Present() {
	fmt.Printf("Шоколад: %s, Производитель: %s\n", c.Name, c.Origin)
}

type GiftCard struct {
	Store  string
	Amount int
}

func (gc *GiftCard) Present() {
	fmt.Printf("Подарочная карта для %s на сумму %d тенге\n", gc.Store, gc.Amount)
}
