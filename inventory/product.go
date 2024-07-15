package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Product struct {
	Name     string
	Quantity int
}

type products map[string]Product

func (p products) saveAll(newProducts []Product) {
	for _, product := range newProducts {
		p[uuid.NewString()] = product
	}
}

func (p products) display() {
	fmt.Printf("\n=========== ALL PRODUCTS (%d) ===========\n\n", len(p))

	for id, product := range p {
		fmt.Printf("ID: %s\n", id)
		fmt.Printf("Name: %s\n", product.Name)
		fmt.Printf("Quantity: %d\n\n", product.Quantity)
	}
}
