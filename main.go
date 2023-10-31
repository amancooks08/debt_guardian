package main

import(
	system "financial_system/debt_guardian"
)

func main() {
	for {
		exit := system.System()
		if exit {
			break
		}
	}
}