package config

import (
	"context"
	"ecommerce/constants"
	"fmt"
	"log"


	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Customer_Collection *mongo.Collection
var Cart_Collection *mongo.Collection
var Seller_Collection *mongo.Collection
var Inventory_Collection *mongo.Collection
var Feedback_Collection *mongo.Collection
var Buynow_Collection *mongo.Collection
var Admin_Collection *mongo.Collection



func init() {
	clientoption := options.Client().ApplyURI(constants.Connectstring)

	client, err := mongo.Connect(context.TODO(), clientoption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb sucessfully connected")
	Customer_Collection = client.Database(constants.DB_Name).Collection(constants.Customer_collection)
	Cart_Collection = client.Database(constants.DB_Name).Collection(constants.Cart_collection)
	Seller_Collection = client.Database(constants.DB_Name).Collection(constants.Seller_Collection)
	Inventory_Collection = client.Database(constants.DB_Name).Collection(constants.Inventory_Collection)
	Feedback_Collection = client.Database(constants.DB_Name).Collection(constants.Feedback_Collection)
	Buynow_Collection = client.Database(constants.DB_Name).Collection(constants.BuyItems_Collection)
	Admin_Collection = client.Database(constants.DB_Name).Collection(constants.Admin_Collection)

	fmt.Println("All Collection Connected")
}
