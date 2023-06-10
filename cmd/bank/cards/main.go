package main

import (
	"bank/pkg/bank/cards"
	"bank/pkg/bank/types"
	"fmt"
)

func main() {
	//payment := []types.Payment{
	//	{
	//		Balance: 10,
	//		ID:      123,
	//	},
	//	{
	//		Balance: 20,
	//		ID:      231,
	//	},
	//	{
	//		Balance: 30,
	//		ID:      324,
	//	},
	//	{
	//		Balance: 40,
	//		ID:      4312,
	//	},
	//}
	////res :=
	//fmt.Println(cards.Avg(payment))

	payment := []types.Payment{
		{
			Balance:  456,
			ID:       999,
			Category: "alcohol",
		},
		{
			Balance:  5,
			ID:       111,
			Category: "alcoholfree",
		}, {
			Balance:  2,
			ID:       7666,
			Category: "alcoholfree",
		},
	}
	fmt.Println(cards.TotalInCategory(payment, "alcoholfree"))

}
