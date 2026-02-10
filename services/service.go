package services

import (
	"fmt"

	"github.com/Hdeee1/go-store-inventory-CLI/models"
)

const LowStockThreshold = 5

var (
	Products []models.Product
	ProductMap map[string]models.Product
)

func InitData() {
	products := []models.Product{
		{Name: "Laptop", Price: 50000000, Stock: 10},
		{Name: "Monitor", Price: 10000000, Stock: 3},
		{Name: "Mouse", Price: 100000, Stock: 0},
	}
	Products = products
	buildProductMap()
}

func buildProductMap() {
	ProductMap = make(map[string]models.Product)
	for _, item := range Products {
		ProductMap[item.Name] = item
	}
}

func getStockStatus(stock int) string {
	if stock == 0 {
		return "Sold Out"
	} else if stock <= LowStockThreshold {
		return "Low Stock"
	}
	return "Available"
}

func ShowAllItems() {
	
	for _, item := range Products {
		status := getStockStatus(item.Stock)
		fmt.Printf("Product Name: %v, Price: %v, Stock: %s\n", item.Name, item.Price, status)
	}
}

func SearchItem() {
	var userInput string
	fmt.Printf("Search item name: ")
	fmt.Scanln(&userInput)

	data, found := ProductMap[userInput]
	if found {
		status := getStockStatus(data.Stock)
		fmt.Printf("Found: Name: %s | Price: %v | Status: %v (Stock: %v) \n", data.Name, data.Price, status, data.Stock)
	} else {
		fmt.Println("Sorry, item not found")
	}
}