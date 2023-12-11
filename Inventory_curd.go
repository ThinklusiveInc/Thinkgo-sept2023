package main

import (
	"context"
	"fmt"

	//"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateItem(item InventoryItem) error {
	_, err := collection.InsertOne(context.TODO(), item)
	if err != nil {
		return err
	}
	fmt.Println("Item created successfully")
	return nil
}

func GetItems() ([]InventoryItem, error) {
	var items []InventoryItem
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	if err := cursor.All(context.TODO(), &items); err != nil {
		return nil, err
	}
	return items, nil
}

func UpdateItem(id primitive.ObjectID, updates bson.M) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updates}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("item not found")
	}

	fmt.Println("Item updated successfully")
	return nil
}

func DeleteItem(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("item not found")
	}

	fmt.Println("Item deleted successfully")
	return nil
}
