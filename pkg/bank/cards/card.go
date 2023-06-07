package cards

import "bank/pkg/bank/types"

func Avg(payment []types.Payment) types.Money {
	var sum types.Money
	for _, card := range payment {
		sum += card.Balance
	}
	l := len(payment)
	return sum / types.Money(l)
}
