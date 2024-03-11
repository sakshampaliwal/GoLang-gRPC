package main

import (
    "fmt"
)

// Define a struct type named Product to represent a product in the inventory
type Product struct {
    ID       int
    Name     string
    Price    float64
    Quantity int
}

// Define a struct type named Customer to represent a customer
type Customer struct {
    ID    int
    Name  string
    Email string
}

// Define a struct type named Order to represent a customer order
type Order struct {
    ID          int
    CustomerID  int
    ProductID   int
    Quantity    int
    TotalAmount float64
}

func main() {
    // Create instances of Product struct representing different products in the inventory
    product1 := Product{ID: 1, Name: "Apple", Price: 1.00, Quantity: 100}
    product2 := Product{ID: 2, Name: "Banana", Price: 0.50, Quantity: 200}

    // Create instances of Customer struct representing different customers
    customer1 := Customer{ID: 101, Name: "Alice", Email: "alice@example.com"}
    customer2 := Customer{ID: 102, Name: "Bob", Email: "bob@example.com"}

    // Create instances of Order struct representing customer orders
    order1 := Order{ID: 1001, CustomerID: 101, ProductID: 1, Quantity: 10, TotalAmount: 10.00}
    order2 := Order{ID: 1002, CustomerID: 102, ProductID: 2, Quantity: 20, TotalAmount: 10.00}

    // Display details of each product, customer, and order
    displayProductDetails(product1)
    displayProductDetails(product2)
    displayCustomerDetails(customer1)
    displayCustomerDetails(customer2)
    displayOrderDetails(order1)
    displayOrderDetails(order2)
}

// Function to display details of a product
func displayProductDetails(product Product) {
    fmt.Println("Product ID:", product.ID)
    fmt.Println("Name:", product.Name)
    fmt.Println("Price:", product.Price)
    fmt.Println("Quantity:", product.Quantity)
    fmt.Println("---------------------------")
}

// Function to display details of a customer
func displayCustomerDetails(customer Customer) {
    fmt.Println("Customer ID:", customer.ID)
    fmt.Println("Name:", customer.Name)
    fmt.Println("Email:", customer.Email)
    fmt.Println("---------------------------")
}

// Function to display details of an order
func displayOrderDetails(order Order) {
    fmt.Println("Order ID:", order.ID)
    fmt.Println("Customer ID:", order.CustomerID)
    fmt.Println("Product ID:", order.ProductID)
    fmt.Println("Quantity:", order.Quantity)
    fmt.Println("Total Amount:", order.TotalAmount)
    fmt.Println("---------------------------")
}
