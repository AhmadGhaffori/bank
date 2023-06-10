package types

type Payment struct {
	ID       int
	Balance  Money
	Category Category
}

type Money int
type Category string
