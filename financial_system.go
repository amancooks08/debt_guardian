package main

type FinancialSystem struct {
	debts    map[string]int
	moneyOwed map[string]int
}

func NewFinancialSystem() *FinancialSystem {
	return &FinancialSystem{
		debts:    make(map[string]int),
		moneyOwed: make(map[string]int),
	}
}