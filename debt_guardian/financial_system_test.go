package debt_guardian

import (
	"testing"
)

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

func TestFinancialSystem_queryMostMoneyOwned(t *testing.T) {
	fs := NewFinancialSystem()

	// Test 1: Query most money owed when there are no transactions
	mostOwed := fs.queryMostMoneyOwed()
	if mostOwed != "" {
		t.Errorf("Test 1 failed, expected an empty string, got %s", mostOwed)
	}

	// Test 2: Query most money owed when there is a single transaction
	fs.addTransaction("Alice", "Bob", 100)
	mostOwed = fs.queryMostMoneyOwed()
	if mostOwed != "Bob" {
		t.Errorf("Test 2 failed, expected 'Bob', got %s", mostOwed)
	}

	// Test 3: Query most money owed when there are multiple transactions
	fs.addTransaction("Charlie", "Bob", 150)
	mostOwed = fs.queryMostMoneyOwed()
	if mostOwed != "Bob" {
		t.Errorf("Test 3 failed, expected 'Bob', got %s", mostOwed)
	}

	// Test 4: Query most money owed when there is a tie
	fs.addTransaction("Dave", "Bob", 100)
	mostOwed = fs.queryMostMoneyOwed()
	if mostOwed != "Bob" {
		t.Errorf("Test 4 failed, expected 'Bob', got %s", mostOwed)
	}
}

func TestFinancialSystem_queryMostDebt(t *testing.T) {
	fs := NewFinancialSystem()

	// Test 1: Query most debt when there are no transactions
	mostDebt := fs.queryMostDebt()
	if mostDebt != "" {
		t.Errorf("Test 1 failed, expected an empty string, got %s", mostDebt)
	}

	// Test 2: Query most debt when there is a single transaction
	fs.addTransaction("Alice", "Bob", 100)
	mostDebt = fs.queryMostDebt()
	if mostDebt != "Alice" {
		t.Errorf("Test 2 failed, expected 'Alice', got %s", mostDebt)
	}

	// Test 3: Query most debt when there are multiple transactions
	fs.addTransaction("Charlie", "Alice", 150)
	mostDebt = fs.queryMostDebt()
	if mostDebt != "Charlie" {
		t.Errorf("Test 3 failed, expected 'Charlie', got %s", mostDebt)
	}

	// Test 4: Query most debt when there is a tie
	fs.addTransaction("Alice", "Dave", 100)
	mostDebt = fs.queryMostDebt()
	if mostDebt != "Alice" {
		t.Errorf("Test 4 failed, expected 'Alice', got %s", mostDebt)
	}
}
