package services

import (
	"fmt"

	"github.com/Hdeee1/go-store-inventory-CLI/models"
)

var Products []models.Product

func InitData() {
	products := []models.Product{
		{Name: "Laptop", Price: 50000000, Stock: 10},
		{Name: "Monitor", Price: 10000000, Stock: 3},
		{Name: "Mouse", Price: 100000, Stock: 0},
	}
	Products = append(Products, products...)
}

func ShowAllItems() {
	
	for _, item := range Products {
		if item.Stock == 0 {
			fmt.Printf("Product Name: %v, Price: %v, Stock: Sold Out \n", item.Name, item.Price)
		} else if item.Stock > 5 {
			fmt.Printf("Product Name: %v, Price: %v, Stock: Low Stock\n", item.Name, item.Price)
		} else {
			fmt.Printf("Product Name: %v, Price: %v, Stock: Available\n", item.Name, item.Price)
		}
	}
}

func SearchItem() {
	itemName := make(map[string]models.Product)
	for _, item := range Products {
		itemName[item.Name] = item
	}

	var userInput string
	fmt.Printf("Search item name (ex: Laptop): ")
	fmt.Scanln(&userInput)

	searchName := userInput
	data, find := itemName[searchName]
	if find {
		fmt.Printf("Find item, name: %s, price: %v, stock: %v \n", data.Name, data.Price, data.Stock)
	} else {
		fmt.Println("Sorry, item not found")
	}
}