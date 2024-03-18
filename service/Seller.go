package service

import (
	"context"
	"ecommerce/config"
	"ecommerce/constants"
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
		_, err = config.Inventory_Collection.InsertOne(context.Background(), inventory)
		if err != nil {

			return false, err
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
	err := config.Seller_Collection.FindOne(context.Background(), filter).Decode(&seller)
	if err != nil {
		return "No Data Found", false, err
	}

	if !seller.IsApproved {
		return "Not Approved yet", false, nil
	}
	if !seller.IsEmailVerified {
		return "Email Not Verified", false, nil
	}
	if check.Password != seller.Password {
		seller.WrongInput++
		update := bson.M{"$set": bson.M{"wronginput": seller.WrongInput}}
		_, err = config.Seller_Collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return "Internal Error", false, err
		}
		return "Wrong Password", false, nil

	}
	if seller.BlockedUser {
		return "Your ID Has been blocked by admin", false, nil
	}
	if seller.WrongInput == 10 {
		return "To many no of attempts", false, nil
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

func RegisterSeller(seller models.Seller) (string, error) {
	if seller.Password != seller.ConfirmPassword {
		return "Password Mismatch", nil
	}
	if seller.Password == "" {
		return "Please Enter the password", nil
	}
	if seller.Seller_Email == "" {
		return "Please Enter the Email", nil
	}
	filter := bson.M{"selleremail": seller.Seller_Email}
	cursor, err := config.Seller_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
		return "Error in Searching", err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var existingSeller models.Seller
		err := cursor.Decode(&existingSeller)
		if err != nil {
			log.Println(err)
			return "Server Error", err
		}

		if existingSeller.IsEmailVerified {
			return "Email already Exists and Verified", nil
		}

		seller.SellerId = existingSeller.SellerId
		seller.BlockedUser = existingSeller.BlockedUser
		seller.WrongInput = existingSeller.WrongInput
		seller.IsEmailVerified = false // Reset to false as it's not verified yet
		seller.IsApproved = false
		seller.VerificationString = GenerateOTP(6)

		_, err = config.Seller_Collection.ReplaceOne(context.Background(), filter, seller)
		if err != nil {
			log.Println(err)
			return "Error In Updating", err
		}

		go SendEmailforVerification(seller.Seller_Email, seller.VerificationString, seller.Seller_Name)
		return "Verify Your Email", nil
	}

	// If no matching email found, insert the new seller
	seller.SellerId = GenerateUniqueCustomerID()
	seller.BlockedUser = false
	seller.WrongInput = 0
	seller.IsEmailVerified = false
	seller.IsApproved = false
	seller.VerificationString = GenerateOTP(6)

	_, err = config.Seller_Collection.InsertOne(context.Background(), seller)
	if err != nil {
		log.Println(err)
		return "Error In Creating", err
	}

	go SendEmailforVerification(seller.Seller_Email, seller.VerificationString, seller.Seller_Name)
	return "Verify Your Email", nil
}

// Email Verification for Seller
func EmailVerificationforSeller(data models.VerifyEmail) (string, error) {
	filter := bson.M{"selleremail": data.Email}
	var Seller models.Seller
	err := config.Seller_Collection.FindOne(context.Background(), filter).Decode(&Seller)
	if err != nil {
		return "Email Not Exists", err
	}
	if Seller.VerificationString == data.VerificationString {
		update := bson.M{"$set": bson.M{"isemailverified": true}}
		_, err := config.Seller_Collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return "Error in Updating", err
		}
		return "Email Verification Successful", nil
	} else {
		return "Wrong OTP", nil
	}

}

func SellerDrashbordDetails(token models.Token) (models.DrashBoard, string, error) {
	var drashboard models.DrashBoard
	id, err := ExtractCustomerID(token.Token, constants.SecretKey)
	if err != nil {
		log.Println(err)
		return drashboard, "Error in Extracting", err
	}
	log.Println(id)

	completedCount, err := config.Buynow_Collection.CountDocuments(context.Background(), map[string]interface{}{
		"status.confirmed": "completed",
		"status.delivered": "completed",
		"sellerid":         id,
	})
	if err != nil {
		fmt.Println("Error counting completed documents:", err)
		return drashboard, "Error in Finding", err
	}

	// Count the number of documents with delivery not completed
	notCompletedCount, err := config.Buynow_Collection.CountDocuments(context.Background(), map[string]interface{}{
		"status.confirmed": "completed",
		"sellerid":         id,
		"status.delivered": map[string]interface{}{"$ne": "completed"},
	})
	if err != nil {
		fmt.Println("Error counting not completed documents:", err)
		return drashboard, "Error in Finding", err
	}
	filter := bson.M{"sellerid": id}
	cursor, err := config.Buynow_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
		return drashboard, "Error in Searching", err
	}
	for cursor.Next(context.Background()) {
		var existingSeller models.Seller
		err := cursor.Decode(&existingSeller)
		if err != nil {
			log.Println(err)
			return drashboard, "Server Error", err
		}

		return drashboard, "Verify Your Email", nil
	}

	drashboard.OrdersCompleted = completedCount
	drashboard.OrdersPending = notCompletedCount

	return drashboard, "Success", nil
}
