package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Name     string
	Category string
	Price    float64
	Quantity int
}

type Inventory struct {
	Products []Product
}

func (i *Inventory) AddProduct(name, category string, price float64, quantity int) {
	i.Products = append(i.Products, Product{
		Name:     name,
		Category: category,
		Price:    price,
		Quantity: quantity,
	})
}

func (i *Inventory) TotalValue() float64 {
	total := 0.0

	for _, product := range i.Products {
		total += product.Price * float64(product.Quantity)
	}

	return total
}

func (i *Inventory) SortByValue() {
	sort.Slice(i.Products, func(a, b int) bool {
		valueA := i.Products[a].Price * float64(i.Products[a].Quantity)
		valueB := i.Products[b].Price * float64(i.Products[b].Quantity)
		return valueA > valueB
	})
}

func (i *Inventory) PrintReport() {
	fmt.Println("Inventory Report")
	fmt.Println("================")

	for _, product := range i.Products {
		value := product.Price * float64(product.Quantity)

		fmt.Printf(
			"%s | %s | $%.2f | %d units | $%.2f\n",
			product.Name,
			product.Category,
			product.Price,
			product.Quantity,
			value,
		)
	}

	fmt.Println("================")
	fmt.Printf("Products: %d\n", len(i.Products))
	fmt.Printf("Total Value: $%.2f\n", i.TotalValue())
}

func main() {
	inventory := Inventory{}

	inventory.AddProduct("Laptop", "Electronics", 899.99, 5)
	inventory.AddProduct("Keyboard", "Accessories", 79.50, 12)
	inventory.AddProduct("Mouse", "Accessories", 39.99, 20)
	inventory.AddProduct("Monitor", "Electronics", 249.99, 8)

	inventory.SortByValue()
	inventory.PrintReport()
}