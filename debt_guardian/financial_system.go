package debt_guardian

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

func (fs *FinancialSystem) addTransaction(personA, personB string, amount int) {
	fs.debts[personA] += amount
	fs.moneyOwed[personB] += amount
}

func (fs *FinancialSystem) queryDebt(person string) int {
	return fs.debts[person]
}

func (fs *FinancialSystem) queryMoneyOwed(person string) int {
	return fs.moneyOwed[person]
}
