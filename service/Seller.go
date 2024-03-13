package service

import (
	"context"
	"ecommerce/config"
	"ecommerce/models"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Add Items To Inventory
func Inventory(inventory models.Inventory) (bool, error) {
	filter := bson.M{"itemname": inventory.ItemName}
	cursor, err := config.Inventory_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
		return false, err
	}
	if cursor.RemainingBatchLength() == 0 {

		var seller models.Seller
		err := config.Seller_Collection.FindOne(context.TODO(), bson.M{"sellerid": inventory.SellerId}).Decode(&seller)
		if err != nil {

			return false, err
		}
		inventory.SellerName = seller.Seller_Name
		_, err1 := config.Inventory_Collection.InsertOne(context.Background(), inventory)
		if err1 != nil {

			return false, err1
		}

		return true, nil
	} else {

		return false, fmt.Errorf("item name already exists")
	}

}

// Check Seller for Login
func CheckSeller(check models.Login) (string, bool, error) {
	var seller models.Seller
	filter := bson.M{"selleremail": check.Email}
	config.Seller_Collection.FindOne(context.Background(), filter).Decode(&seller)
	if check.Password != seller.Password {
		return "InvalidPassword", false, nil
	}
	if seller.BlockedUser {
		return "Your ID Has been blocked by admin", false, nil
	}
	if seller.WrongInput == 10 {
		return "To many no of attempts", false,nil
	}

	result, err := CreateToken(seller.Seller_Email, seller.SellerId)
	if err != nil {
		return "Error In Creating Token ", false, err
	}

	return result, true, nil
}

// Update Product
func UpdateProductBySeller(update models.UpdateProduct) int {
	filter := bson.M{"itemname": update.ProductName}
	update1 := bson.M{"$set": bson.M{update.Attribute: update.New_Value}}
	options := options.Update()
	_, err := config.Inventory_Collection.UpdateOne(context.TODO(), filter, update1, options)
	if err != nil {

		return 0
	}
	return 1
}

// Delete Product
func DeleteProductBySeller(delete models.DeleteBySeller) int {
	filter := bson.M{"itemname": delete.ProductName}
	_, err := config.Inventory_Collection.DeleteOne(context.Background(), filter)
	if err != nil {

		return 0
	}
	return 1
}

// Display Orders
func Orders() []models.Customerorder {
	var Order []models.Customerorder
	filter := bson.M{}
	cursor, err := config.Buynow_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	for cursor.Next(context.Background()) {
		var order models.Customerorder
		err := cursor.Decode(&order)
		if err != nil {
			log.Println(err)
		}
		Order = append(Order, order)
	}
	fmt.Println(Order)
	return Order
}
