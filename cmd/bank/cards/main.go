package main

import (
	"bank/pkg/bank/cards"
	"bank/pkg/bank/types"
	"fmt"
)

func main() {
	payment := []types.Payment{
		{
			Balance: 10,
			ID:      123,
		},
		{
			Balance: 20,
			ID:      231,
		},
		{
			Balance: 30,
			ID:      324,
		},
		{
			Balance: 40,
			ID:      4312,
		},
	}
	//res :=
	fmt.Println(cards.Avg(payment))

}
