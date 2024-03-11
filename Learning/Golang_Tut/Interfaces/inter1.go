package main

import (
    "fmt"
)

// Define an interface named AccountOperations
type AccountOperations interface {
    Deposit(amount float64)
    Withdraw(amount float64) error
    Display()
}

// Define a struct type named Account to represent a bank account
type Account struct {
    AccountNumber int
    HolderName    string
    Balance       float64
}

// Implement the methods of AccountOperations interface for Account struct
func (a *Account) Deposit(amount float64) {
    a.Balance += amount
}

func (a *Account) Withdraw(amount float64) error {
    if amount > a.Balance {
        return fmt.Errorf("insufficient funds")
    }
    a.Balance -= amount
    return nil
}

func (a *Account) Display() {
    fmt.Printf("Account Number: %d\n", a.AccountNumber)
    fmt.Printf("Holder Name: %s\n", a.HolderName)
    fmt.Printf("Balance: $%.2f\n", a.Balance)
    fmt.Println("---------------------------")
}

// Function to process account operations
func ProcessOperations(account AccountOperations, depositAmount, withdrawAmount float64) {
    account.Display() // Display initial account details
    account.Deposit(depositAmount)
    fmt.Printf("Deposited: $%.2f\n", depositAmount)
    account.Display() // Display account details after deposit
    err := account.Withdraw(withdrawAmount)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("Withdrawn: $%.2f\n", withdrawAmount)
        account.Display() // Display account details after withdrawal
    }
}

func main() {
    // Create a new instance of Account struct
    account1 := Account{
        AccountNumber: 123456,
        HolderName:    "Alice",
        Balance:       1000.00,
    }

    // Process operations on account1
    ProcessOperations(&account1, 500.00, 200.00)
}
