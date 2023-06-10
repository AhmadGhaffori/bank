package cards

import "bank/pkg/bank/types"

//func Avg(payment []types.Payment) types.Money {
//	var sum types.Money
//	for _, card := range payment {
//		sum += card.Balance
//	}
//	l := len(payment)
//	return sum / types.Money(l)
//}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, categories := range payments {
		if categories.Category == category {
			sum += categories.Balance
		}

	}
	return sum
}
