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

func (fs *FinancialSystem) queryMostMoneyOwed() string {
	maxOwed := 0
	person := ""
	for key, value := range fs.moneyOwed {
		if value > maxOwed {
			maxOwed = value
			person = key
		}
	}
	return person
}

func (fs *FinancialSystem) queryMostDebt() string {
	maxDebt := 0
	person := ""
	for key, value := range fs.debts {
		if value > maxDebt {
			maxDebt = value
			person = key
		}
	}
	return person
}