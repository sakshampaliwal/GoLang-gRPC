package main

import (
    "fmt"
)

// Define a struct type named Account to represent a bank account
type Account struct {
    AccountNumber int
    HolderName    string
    Balance       float64
}

// Function to deposit money into an account
func (a *Account) Deposit(amount float64) {
    a.Balance += amount
}

// Function to withdraw money from an account
func (a *Account) Withdraw(amount float64) error {
    if amount > a.Balance {
        return fmt.Errorf("insufficient funds")
    }
    a.Balance -= amount
    return nil
}

// Function to display account details
func (a *Account) Display() {
    fmt.Printf("Account Number: %d\n", a.AccountNumber)
    fmt.Printf("Holder Name: %s\n", a.HolderName)
    fmt.Printf("Balance: $%.2f\n", a.Balance)
    fmt.Println("---------------------------")
}

func main() {
    // Create a new instance of Account struct
    account1 := Account{
        AccountNumber: 123456,
        HolderName:    "Alice",
        Balance:       1000.00,
    }

    // Deposit some money
    account1.Deposit(500.00)

    // Withdraw some money
    err := account1.Withdraw(200.00)
    if err != nil {
        fmt.Println("Error:", err)
    }

    // Display account details
    account1.Display()
}
