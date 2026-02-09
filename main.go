package main

import (
	"fmt"
	"os"

	"github.com/Hdeee1/go-store-inventory-CLI/services"
)

func main() {
	services.InitData()
	for {
		var choice int
		fmt.Println("1. Show All Items")
		fmt.Println("2. Search Item")
		fmt.Println("3. Quit")

		fmt.Println("Chose options:")
		fmt.Scanln(&choice)

		switch choice {
		case 1 :
			services.ShowAllItems()
			return
		case 2:
			services.SearchItem()
			return
		case 3:
			os.Exit(0)
		default:
			fmt.Println("")
		}
	}
}