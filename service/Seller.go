package service

import (
	"context"
	"ecommerce/config"
	"ecommerce/constants"
	"ecommerce/models"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Add Items To Inventory
func Inventory(inventory models.Inventory) (string, error) {

	sellerid, err := ExtractCustomerID(inventory.SellerName, constants.SecretKey)
	if err != nil {
		return "Login Expired", err
	}
	inventory.SellerId = sellerid

	filter := bson.M{"itemname": inventory.ItemName}
	cursor, err := config.Inventory_Collection.Find(context.Background(), filter)
	if err != nil {
		return "Error while Finding", err
	}
	if cursor.RemainingBatchLength() == 0 {

		var seller models.Seller
		err := config.Seller_Collection.FindOne(context.TODO(), bson.M{"sellerid": inventory.SellerId}).Decode(&seller)
		if err != nil {
			return "Seller not found", err
		}
		inventory.SellerName = seller.Seller_Name
		inventory.ProductId = GenerateUniqueProductID()
		_, err = config.Inventory_Collection.InsertOne(context.Background(), inventory)
		if err != nil {
			return "Error while Adding Product", err
		}

		return "Success", nil
	} else {
		return "Item Name Already exists", nil
	}

}

// Generate Product ID
func GenerateUniqueProductID() string {
	// Implement your logic to generate a unique customer ID (e.g., UUID, timestamp, etc.)
	// For example, you can use a combination of timestamp and random characters
	return fmt.Sprintf("%d%s", time.Now().UnixNano(), GetRandomString(8))
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
func UpdateProductBySeller(update models.UpdateProduct) (string,error) {
	id,err := ExtractCustomerID(update.Token,constants.SecretKey)
	if err != nil{
		return "Login Expired",err
	}

	var inventory models.Inventory

	filter := bson.M{"productid": update.ProductID}
	err = config.Inventory_Collection.FindOne(context.Background(),filter).Decode(&inventory)
	if err != nil{
		return "Product Not Found",err
	}

	if inventory.SellerId != id{
		return "Product Not Found",nil
	}

	updatefilter := bson.M{"$set": bson.M{update.Attribute: update.New_Value}}

	if(update.Attribute == "price" || update.Attribute == "sellerquantity"){
		value,err := stringToInt(update.New_Value)
		if err != nil{
			return "Error in converting",err
		}
		updatefilter = bson.M{"$set": bson.M{update.Attribute: value}}
	}else{
		updatefilter = bson.M{"$set": bson.M{update.Attribute: update.New_Value}}
	}
	
	options := options.Update()
	_, err = config.Inventory_Collection.UpdateOne(context.TODO(), filter, updatefilter, options)
	if err != nil {
		return "Error in Updating",nil
	}
	return "Updated Successfully",nil
}

// Delete Product
func DeleteProductBySeller(delete models.DeleteBySeller) (string, error) {
	id, err := ExtractCustomerID(delete.Token, constants.SecretKey)
	if err != nil {
		return "Login Expired", err
	}
	var inventory models.Inventory
	filter := bson.M{"productid": delete.ProductID}
	err = config.Inventory_Collection.FindOne(context.Background(),filter).Decode(&inventory)
	if err != nil{
		return "Product Not Found",err
	}
	if inventory.SellerId != id{
		return "Product Not Found",nil
	}
	_, err = config.Inventory_Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return "Error in Deleting",err
	}
	return "Deleted Successfully",nil
}

// Display All Orders
func Orders(token models.Token) ([]models.AddOrder, string, error) {
	id, err := ExtractCustomerID(token.Token, constants.SecretKey)
	if err != nil {
		return nil, "Login Expired", err
	}
	var Order []models.AddOrder
	filter := bson.M{"sellerid": id}
	cursor, err := config.Buynow_Collection.Find(context.Background(), filter)
	if err != nil {
		return nil, "Error in Finding", err
	}
	for cursor.Next(context.Background()) {
		var order models.AddOrder
		err := cursor.Decode(&order)
		if err != nil {
			return nil, "Error in Extracting", err
		}
		Order = append(Order, order)
	}
	return Order, "Success", nil
}

// Display All Completed Orders
func CompletedOrders(token models.Token) ([]models.AddOrder, string, error) {
	id, err := ExtractCustomerID(token.Token, constants.SecretKey)
	if err != nil {
		return nil, "Login Expired", err
	}
	var Order []models.AddOrder
	filter := bson.M{
		"sellerid": id,
		"$and": []bson.M{
			{"status.processing": "completed"},
			{"status.quality": "completed"},
			{"status.dispatched": "completed"},
			{"status.delivered": "completed"},
		},
	}
	cursor, err := config.Buynow_Collection.Find(context.Background(), filter)
	if err != nil {
		return nil, "Error in Finding", err
	}
	for cursor.Next(context.Background()) {
		var order models.AddOrder
		err := cursor.Decode(&order)
		if err != nil {
			return nil, "Error in Extracting", err
		}
		Order = append(Order, order)
	}
	return Order, "Success", nil
}

// Display All Pending Orders
func PendingOrders(token models.Token) ([]models.AddOrder, string, error) {
	id, err := ExtractCustomerID(token.Token, constants.SecretKey)
	if err != nil {
		return nil, "Login Expired", err
	}
	var Order []models.AddOrder
	filter := bson.M{
		"sellerid": id,
		"$or": []bson.M{
			{"status.processing": "pending"},
			{"status.quality": "pending"},
			{"status.dispatched": "pending"},
		},
	}
	cursor, err := config.Buynow_Collection.Find(context.Background(), filter)
	if err != nil {
		return nil, "Error in Finding", err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var order models.AddOrder
		if err := cursor.Decode(&order); err != nil {
			return nil, "Error in Extracting", err
		}
		Order = append(Order, order)
	}

	if err := cursor.Err(); err != nil {
		return nil, "Error during cursor iteration", err
	}

	return Order, "Success", nil
}

// Yet To Deliver Order
func YettoDeliverOrders(token models.Token) ([]models.AddOrder, string, error) {
	id, err := ExtractCustomerID(token.Token, constants.SecretKey)
	if err != nil {
		return nil, "Login Expired", err
	}
	var Order []models.AddOrder
	filter := bson.M{
		"sellerid": id,
		"$and": []bson.M{
			{"status.processing": "completed"},
			{"status.quality": "completed"},
			{"status.dispatched": "completed"},
			{"status.delivered": "pending"},
		},
	}
	cursor, err := config.Buynow_Collection.Find(context.Background(), filter)
	if err != nil {
		return nil, "Error in Finding", err
	}
	for cursor.Next(context.Background()) {
		var order models.AddOrder
		err := cursor.Decode(&order)
		if err != nil {
			return nil, "Error in Extracting", err
		}
		Order = append(Order, order)
	}
	return Order, "Success", nil
}

// Create Seller
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
		return "Error in Searching", err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var existingSeller models.Seller
		err := cursor.Decode(&existingSeller)
		if err != nil {
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
		return drashboard, "Error in Extracting", err
	}
	completedCount, err := config.Buynow_Collection.CountDocuments(context.Background(), map[string]interface{}{
		"status.confirmed": "completed",
		"status.delivered": "completed",
		"sellerid":         id,
	})
	if err != nil {
		return drashboard, "Error in Finding", err
	}

	// Count the number of documents with delivery not completed
	notCompletedCount, err := config.Buynow_Collection.CountDocuments(context.Background(), map[string]interface{}{
		"status.confirmed": "completed",
		"sellerid":         id,
		"status.delivered": map[string]interface{}{"$ne": "completed"},
	})

	if err != nil {
		return drashboard, "Error in Finding", err
	}

	filter := bson.M{"sellerid": id}
	cursor, err := config.Buynow_Collection.Find(context.Background(), filter)
	if err != nil {
		return drashboard, "Error in Searching", err
	}
	for cursor.Next(context.Background()) {
		var order models.AddOrder
		err := cursor.Decode(&order)
		if err != nil {

			return drashboard, "Server Error", err
		}
		drashboard.Orders++
		drashboard.TotalAmount += int64(order.ItemsToBuy.TotalPrice)
	}

	drashboard.OrdersCompleted = completedCount
	drashboard.OrdersPending = notCompletedCount

	return drashboard, "Success", nil
}

// Get All Products
func GetAllProducts(token models.Token) ([]models.Inventory, string, error) {
	id, err := ExtractCustomerID(token.Token, constants.SecretKey)
	if err != nil {
		return nil, "Login Expired", err
	}
	filter := bson.M{"sellerid": id}
	cursor, err := config.Inventory_Collection.Find(context.Background(), filter)
	if err != nil {
		return nil, "No Products found", err
	}
	var Inventory []models.Inventory
	for cursor.Next(context.Background()) {
		var inventory models.Inventory
		err := cursor.Decode(&inventory)
		if err != nil {
			return nil, "Server Error", err
		}
		Inventory = append(Inventory, inventory)
	}
	return Inventory, "Success", nil
}

// Get all buyed Customer
func GetAllBuyedCustomer(token models.Token) ([]models.Customer, string, error) {
	id, err := ExtractCustomerID(token.Token, constants.SecretKey)
	if err != nil {
		return nil, "Login Expired", err
	}
	var CustomerId []string
	filter := bson.M{"sellerid": id}
	cursor, err := config.Buynow_Collection.Find(context.Background(), filter)
	if err != nil {
		return nil, "Error in Searching", err
	}
	for cursor.Next(context.Background()) {
		var order models.AddOrder
		err := cursor.Decode(&order)
		if err != nil {
			return nil, "Server Error", err
		}
		CustomerId = append(CustomerId, order.CustomerId)
	}
	CustomerId = uniqueStrings(CustomerId)
	var Customer []models.Customer
	for _, value := range CustomerId {
		var customer models.Customer
		filter = bson.M{"customerid": value}
		err := config.Customer_Collection.FindOne(context.Background(), filter).Decode(&customer)
		if err != nil {
			return nil, "Error in Finding", err
		}
		customer.BlockedUser = false
		customer.ConfirmPassword = ""
		customer.Password = ""
		customer.WrongInput = 0
		Customer = append(Customer, customer)
	}
	return Customer, "Success", nil
}

// Find Unique Customer
func uniqueStrings(input []string) []string {
	uniqueMap := make(map[string]bool)
	var uniqueSlice []string

	for _, str := range input {
		if !uniqueMap[str] {
			uniqueMap[str] = true
			uniqueSlice = append(uniqueSlice, str)
		}
	}

	return uniqueSlice
}

// Get Order Details
func GetCustromerOrderforSeller(details models.GetOrder) (*models.AddOrder, string, error) {
	var orderDetails models.AddOrder

	id, err := ExtractCustomerID(details.Token, constants.SecretKey)
	if err != nil {
		return &orderDetails, "Login Expired", err
	}

	filter := bson.M{"orderid": details.OrderID}
	err = config.Buynow_Collection.FindOne(context.Background(), filter).Decode(&orderDetails)
	if err != nil {
		return &orderDetails, "No Result found", err
	}
	if orderDetails.SellerId != id {
		return nil, "Order not Found", nil
	}
	return &orderDetails, "Success", nil
}

// Update Order Tracking
func UpdateOrderTracking(details models.OrderTracking) (string, error) {
	id, err := ExtractCustomerID(details.Token, constants.SecretKey)
	if err != nil {
		return "Login Expired", err
	}
	log.Println(details)
	var order models.AddOrder
	filter := bson.M{"orderid": details.OrderID}
	err = config.Buynow_Collection.FindOne(context.Background(), filter).Decode(&order)
	if err != nil {
		return "Error in Finding", err
	}
	if order.SellerId != id {
		return "Order not Found", nil
	}
	if order.Status.Product_Delivered == "completed" {
		return "Order Already Delivered", nil
	}
	if order.Status.Product_Dispatched == "completed" {
		return "Order Already Dispatched", nil
	}
	if details.Feild == "processing" && order.Status.Processing_Order == "pending" {
		update := bson.M{"$set": bson.M{"status.processing": "completed"}}
		_, err = config.Buynow_Collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return "Error in Updating", err
		}
		return "Updated Successfully", nil
	} else if order.Status.Processing_Order == "completed" && details.Feild == "processing" {
		return "Processing Already Completed", nil
	}

	if details.Feild == "quality" && order.Status.Quality_Check == "pending" && order.Status.Processing_Order == "completed" {
		update := bson.M{"$set": bson.M{"status.quality": "completed"}}
		_, err = config.Buynow_Collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return "Error in Updating", err
		}
		return "Updated Successfully", nil
	} else if details.Feild == "quality" && order.Status.Processing_Order == "pending" {
		return "Please Complete Processing before this", nil
	} else if order.Status.Quality_Check == "completed" && details.Feild == "quality" {
		return "Quality Check Already Completed", nil
	}

	if details.Feild == "dispatched" && order.Status.Product_Dispatched == "pending" && order.Status.Quality_Check == "completed" {
		update := bson.M{"$set": bson.M{"status.dispatched": "completed"}}
		_, err = config.Buynow_Collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return "Error in Updating", err
		}
		return "Updated Successfully", nil
	} else if details.Feild == "dispatched" && order.Status.Quality_Check == "pending" {
		return "Please Complete Quality Check before this", nil
	} else if order.Status.Product_Dispatched == "completed" && details.Feild == "dispatched" {
		return "Product Dispatched Already Completed", nil
	}

	return "Unable to update data", nil

}

//Get Product Data
func GetProductData(details models.DeleteBySeller)(*models.Inventory,string,error){
	id,err := ExtractCustomerID(details.Token,constants.SecretKey)
	if err != nil{
		return nil,"Login Expired",err
	}
	var inventory models.Inventory
	filter  := bson.M{"productid":details.ProductID}
	err = config.Inventory_Collection.FindOne(context.Background(),filter).Decode(&inventory)
	if err != nil{
		return nil,"Error in Finding",err
	}

	if inventory.SellerId != id{
		return nil,"Product not Found",nil
	}

	return &inventory,"Success",nil
}
