package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunCLI(store *MemoryAccountStore) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== CLI BANKING SYSTEM ===")
		fmt.Println("1. Create Account")
		fmt.Println("2. Check Balance")
		fmt.Println("3. Deposit")
		fmt.Println("4. Withdraw")
		fmt.Println("5. Transfer")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			reader.ReadString('\n')
			continue
		}
	
		reader.ReadString('\n')

		switch choice {
		case 1:
			var acc Account
			fmt.Print("Enter Account ID: ")
			idInput, _ := reader.ReadString('\n')
			acc.ID = strings.TrimSpace(idInput)

			fmt.Print("Enter Owner Name: ")
			nameInput, _ := reader.ReadString('\n')
			acc.Owner = strings.TrimSpace(nameInput)

			fmt.Print("Initial Balance: ")
			fmt.Scan(&acc.Balance)
			reader.ReadString('\n')

			err := store.AddAccount(acc)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Account created successfully!")
			}

		case 2:
			fmt.Print("Enter Account ID: ")
			idInput, _ := reader.ReadString('\n')
			id := strings.TrimSpace(idInput)

			acc, exists := store.GetAccount(id)
			if !exists {
				fmt.Println("Account not found.")
			} else {
				fmt.Printf("Account: %s | Owner: %s | Balance: $%.2f\n", acc.ID, acc.Owner, acc.Balance)
			}

		case 3:
			fmt.Print("Enter Account ID: ")
			idInput, _ := reader.ReadString('\n')
			id := strings.TrimSpace(idInput)

			var amount float64
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&amount)
			reader.ReadString('\n')

			err := store.Deposit(id, amount)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful!")
			}

		case 4:
			fmt.Print("Enter Account ID: ")
			idInput, _ := reader.ReadString('\n')
			id := strings.TrimSpace(idInput)

			var amount float64
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&amount)
			reader.ReadString('\n')

			err := store.Withdraw(id, amount)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdrawal successful!")
			}

		case 5:
			fmt.Print("Enter Source Account ID: ")
			fromInput, _ := reader.ReadString('\n')
			fromID := strings.TrimSpace(fromInput)

			fmt.Print("Enter Destination Account ID: ")
			toInput, _ := reader.ReadString('\n')
			toID := strings.TrimSpace(toInput)

			var amount float64
			fmt.Print("Enter amount to transfer: ")
			fmt.Scan(&amount)
			reader.ReadString('\n')

			err := store.Transfer(fromID, toID, amount)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Transfer successful!")
			}

		case 6:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Choose between 1 and 6.")
		}
	}
}