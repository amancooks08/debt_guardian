# DebtGuardian
Given a list of transactions between individuals in the form A,B,X, where A owes B X amount of money. The **DebtGuardian** system can process these transactions and provide insights into the financial relationships between individuals.


# Functions that are to be performed
1. Process Transactions: The system must take a list of transactions in the
form A,B,X and process them to update the respective debts between the individuals
involved.
2. Query Debt Owed by Person: A user must be able to query how much debt a specific
person owes to all other users.
3. Query Money Owed to Person: A user must be able to query how much money is owed
to a specific person by all other users.
4. Query Person with Most Money Owed: The system should be able to identify the
person to whom the most money is owed.
5. Query Person with Most Debt: The system should be able to identify the person who
owes the most money to others.
6. Command Line Interface: The system should provide a command-line interface for
users to input transactions and query the insights mentioned above.

# Steps to access the system
1. Clone the repository to your local machine.
2. Ensure that you have Golang installed on your machine.
3. Open a terminal and navigate to the directory where the repository is cloned.
4. Run the `go mod tidy` command to ensure you have all the required modules present.
5. Run the `go build` command to ensure the binary is made, and the build is done.
6. Run the main game file using the go run command: `go run main.go` .
7. Follow the on-screen instructions to access **DebtGuardian**. Use the keyboard to input your choices during the game.
8. Have hassle free debts using **DebtGuardian**


# Contributing
If you would like to contribute to the game, feel free to submit a pull request. Please make sure to follow the coding standards and include appropriate comments in your code. Also, update the README.md file with any necessary changes.


# License
This project is licensed under the MIT License, which allows for free use, modification, and distribution of the code. However, please acknowledge the original authors if you use or modify this project for your own purposes.