package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type InventoryItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Quantity    int                `bson:"quantity"`
}
