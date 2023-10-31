package debt_guardian

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func System() (exit bool) {
	financialSystem := NewFinancialSystem()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter a command (add_transaction, query_debt, query_money_owed, query_most_money_owed, query_most_debt, exit):")
		scanner.Scan()
		command := scanner.Text()

		if command == "add_transaction" {
			fmt.Println("Enter transaction (A,B,X):")
			scanner.Scan()
			transaction := scanner.Text()
			parts := strings.Split(transaction, ",")
			if len(parts) == 3 {
				personA := parts[0]
				personB := parts[1]
				if(personA == "" || personB == "") {
					fmt.Println("invalid name entered, name can't be empty")
				}
				amount := 0
				_, err := fmt.Sscanf(parts[2], "%d", &amount)
				if err == nil {
					financialSystem.addTransaction(personA, personB, amount)
				}
			}
		} else if command == "query_debt" {
			fmt.Println("Enter person:")
			scanner.Scan()
			person := scanner.Text()
			if person == "" {	
				fmt.Println("invalid name entered, name can't be empty")
			}
			debt := financialSystem.queryDebt(person)
			fmt.Printf("%s owes %d\n", person, debt)
		} else if command == "query_money_owed" {
			fmt.Println("Enter person:")
			scanner.Scan()
			person := scanner.Text()
			if person == "" {
				fmt.Println("invalid name entered, name can't be empty")
			}
			moneyOwed := financialSystem.queryMoneyOwed(person)
			fmt.Printf("%s is owed %d\n", person, moneyOwed)
		} else if command == "query_most_money_owed" {
			person := financialSystem.queryMostMoneyOwed()
			fmt.Printf("%s is owed the most money.\n", person)
		} else if command == "query_most_debt" {
			person := financialSystem.queryMostDebt()
			fmt.Printf("%s owes the most money.\n", person)
		} else if command == "exit" {
			return true
		}
	}
}
