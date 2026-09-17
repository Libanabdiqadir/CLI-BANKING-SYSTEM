# CLI-BANKING-SYSTEM
A command-line banking application built with **Go**, designed to practice concurrency control (`sync.RWMutex`), modular code organization, and robust error handling.

## Features
* **Account Management:** Create and track unique bank accounts.
* **Transactions:** Deposit, withdraw, and transfer funds between accounts.
* **Thread-Safe Storage:** Uses Go's `sync.RWMutex` to safely handle concurrent operations in memory.
* **Interactive CLI:** Simple command-line menu for easy user interaction.

## Prerequisites
Make sure you have [Go](https://golang.org/) installed on your machine.

## How to Run

Clone the repository and run the application from your terminal:

```bash
git clone [https://github.com/Libanabdiqadir/CLI-Bank-Account.git](https://github.com/Libanabdiqadir/CLI-Bank-Account.git)
cd CLI-Bank-Account
go run .
