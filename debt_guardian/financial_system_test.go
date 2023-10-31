package debt_guardian

import "testing"

func TestFinancialSystem_addTransaction(t *testing.T) {
	fs := NewFinancialSystem()

    // Test 1: Add a transaction and check if debts and moneyOwed are updated correctly
    fs.addTransaction("Alice", "Bob", 100)
    if fs.queryDebt("Alice") != 100 || fs.queryMoneyOwed("Bob") != 100 {
        t.Errorf("AddTransaction test 1 failed")
    }

    // Test 2: Add another transaction for the same person and check if the amounts are cumulative
    fs.addTransaction("Alice", "Bob", 50)
    if fs.queryDebt("Alice") != 150 || fs.queryMoneyOwed("Bob") != 150 {
        t.Errorf("AddTransaction test 2 failed")
    }
}
