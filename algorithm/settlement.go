package algorithm

type Balance struct {
	UserID uint
	Amount float64
}

type Transaction struct {
	From   uint
	To     uint
	Amount float64
}

func Settle(balances []Balance) []Transaction {

	var debtors []Balance
	var creditors []Balance

	for _, b := range balances {

		if b.Amount < 0 {
			debtors = append(debtors, b)
		}

		if b.Amount > 0 {
			creditors = append(creditors, b)
		}
	}

	var result []Transaction

	i, j := 0, 0

	for i < len(debtors) && j < len(creditors) {

		d := &debtors[i]
		c := &creditors[j]

		min := -d.Amount
		if c.Amount < min {
			min = c.Amount
		}

		result = append(result, Transaction{
			From:   d.UserID,
			To:     c.UserID,
			Amount: min,
		})

		d.Amount += min
		c.Amount -= min

		if d.Amount == 0 {
			i++
		}

		if c.Amount == 0 {
			j++
		}
	}

	return result
}
