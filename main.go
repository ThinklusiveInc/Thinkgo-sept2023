package main

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	item := InventoryItem{
		Name:        "Laptop",
		Description: "High-performance laptop",
		Quantity:    10,
	}

	// Create Item
	err := CreateItem(item)
	if err != nil {
		log.Fatal(err)
	}

	// Get Items
	items, err := GetItems()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Items:", items)

	// Update Item
	updateID := items[0].ID
	err = UpdateItem(updateID, bson.M{"quantity": 8})
	if err != nil {
		log.Fatal(err)
	}

	// Get Updated Items
	updatedItems, err := GetItems()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated Items:", updatedItems)

	// Delete Item
	deleteID := updatedItems[0].ID
	err = DeleteItem(deleteID)
	if err != nil {
		log.Fatal(err)
	}

	// Get Remaining Items
	remainingItems, err := GetItems()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Remaining Items:", remainingItems)
}
