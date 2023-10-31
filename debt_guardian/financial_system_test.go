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

func TestFinancialSystem_queryDebt(t *testing.T) {
	fs := NewFinancialSystem()

	// Test 1: Query for a person with no transactions and check if the amount is 0
	if fs.queryDebt("Alice") != 0 {
		t.Errorf("QueryDebt test 1 failed")
	}

	// Test 2: Query for a person with transactions and check if the amount is correct
	fs.addTransaction("Alice", "Bob", 100)
	if fs.queryDebt("Alice") != 100 {
		t.Errorf("QueryDebt test 2 failed")
	}
}

func TestFinancialSystem_queryMoneyOwed(t *testing.T) {
	fs := NewFinancialSystem()

	// Test 1: Query for a person with no transactions and check if the amount is 0
	if fs.queryMoneyOwed("Alice") != 0 {
		t.Errorf("QueryMoneyOwed test 1 failed")
	}

	// Test 2: Query for a person with transactions and check if the amount is correct
	fs.addTransaction("Alice", "Bob", 100)
	if fs.queryMoneyOwed("Bob") != 100 {
		t.Errorf("QueryMoneyOwed test 2 failed")
	}
}