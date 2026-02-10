package main

import (
	"fmt"
	"os"

	"github.com/Hdeee1/go-store-inventory-CLI/services"
)

func main() {
	services.InitData()
	for {
		displayMenu()
		choice := getUserChoice()

		switch choice {
		case 1 :
			services.ShowAllItems()
		case 2:
			services.SearchItem()
		case 3:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice, try again!")
		}
		fmt.Println()
	}
}

func displayMenu() {
	fmt.Println("1. Show All Items")
	fmt.Println("2. Search Item")
	fmt.Println("3. Quit")
	fmt.Println("Chose options:")
}

func getUserChoice() int {
	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Scanln()
		return -1
	}
	return choice
}