package service

import (
	"context"
	"ecommerce/config"
	"ecommerce/constants"
	"ecommerce/models"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateCart(cart models.Cart) *mongo.UpdateResult {
	filter := bson.D{
		{Key: "$and", Value: []interface{}{
			bson.D{{Key: "customerid", Value: cart.CustomerId}},
			bson.D{{Key: "name", Value: cart.Name}},
		}},
	}
	_, err := config.Cart_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	update := bson.M{"$set": bson.M{"name": cart.Name, "quantity": cart.Quantity, "totalprice": cart.Price, "price": cart.Price / float64(cart.Quantity)}}
	result, err := config.Cart_Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return result

}

func Cart(customerid string) []models.Cart {

	filter := bson.M{"customerid": customerid}
	cursor, err := config.Cart_Collection.Find(context.Background(), filter)
	if err != nil {

		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var Cart []models.Cart
	for cursor.Next(context.Background()) {
		var cart models.Cart
		err := cursor.Decode(&cart)
		if err != nil {

			log.Fatal(err)
		}
		Cart = append(Cart, cart)
	}
	return Cart
}
func Search(productName string) []models.Inventory1 {

	filter := bson.M{"itemcategory": productName}
	cursor, err := config.Inventory_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var Inventory []models.Inventory1
	for cursor.Next(context.Background()) {
		var inventory models.Inventory1
		err := cursor.Decode(&inventory)
		if err != nil {
			log.Fatal(err)
		}
		Inventory = append(Inventory, inventory)
	}
	return Inventory
}
func Getalldata() []models.Customer {
	filter := bson.D{}
	cursor, err := config.Customer_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var Profiles []models.Customer
	for cursor.Next(context.Background()) {
		var profile models.Customer
		err := cursor.Decode(&profile)
		if err != nil {
			log.Fatal(err)
		}
		Profiles = append(Profiles, profile)
	}
	return Profiles
}
func Getinventorydata() []models.Inventory {
	filter := bson.D{}
	cursor, err := config.Inventory_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var Inventorydata []models.Inventory
	for cursor.Next(context.Background()) {
		var inventory models.Inventory
		err := cursor.Decode(&inventory)
		if err != nil {
			log.Fatal(err)
		}
		Inventorydata = append(Inventorydata, inventory)
	}
	return Inventorydata
}

func Insert(profile models.Customer) int {

	name := profile.Name
	regexPattern := "^[a-zA-Z ]+$"
	regexpObject := regexp.MustCompile(regexPattern)
	match := regexpObject.FindString(name)

	if match == "" {
		return 4
	}

	phonecount := countdigits(profile.Phone_No)

	if phonecount != 10 {
		return 5
	}

	if profile.Password != profile.ConfirmPassword {
		return 3
	}

	pincount := countdigits(int(profile.Pincode))

	if pincount != 6 {
		return 6
	}

	filter := bson.M{"email": profile.Email}
	cursor, err := config.Customer_Collection.Find(context.Background(), filter)
	defer cursor.Close(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	profile.CustomerId = GenerateUniqueCustomerID()

	if cursor.RemainingBatchLength() == 0 {

		inserted, err := config.Customer_Collection.InsertOne(context.Background(), profile)
		if err != nil {
			log.Fatal(err)
			return 0
		}

		fmt.Println("Inserted", inserted.InsertedID)
		return 1
	}
	return 2
}
func Addtocart(addtocart models.Addtocart1) bool {
	filter := bson.D{
		{Key: "$and", Value: []interface{}{
			bson.D{{Key: "customerid", Value: addtocart.CustomerId}},
			bson.D{{Key: "name", Value: addtocart.Name}},
		}},
	}

	cursor, err := config.Cart_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())
	type addcart struct {
		CustomerId string  `json:"customerid" bson:"customerid"`
		Name       string  `json:"name" bson:"name"`
		Price      float64 `json:"price" bson:"price"`
		Quantity   int32   `json:"quantity" bson:"quantity"`
	}
	// Check if any items were found
	if cursor.RemainingBatchLength() == 0 {
		// Item not found, so insert a new item with quantity 1
		cart := addcart{CustomerId: addtocart.CustomerId, Name: addtocart.Name, Price: addtocart.Price, Quantity: 1}
		_, err := config.Cart_Collection.InsertOne(context.Background(), cart)
		if err != nil {
			log.Fatal(err)
			return false
		}

		return true
	}

	// Item already exists, update its quantity
	var cart addcart
	for cursor.Next(context.Background()) {
		err = cursor.Decode(&cart)

		if err != nil {
			log.Fatal(err)
		}
	}
	// Item already exists, update its quantity
	cart.Quantity++
	cart.Price = cart.Price + addtocart.Price
	// Use the UpdateOne method to increment the quantity
	update := bson.M{"$set": bson.M{"quantity": cart.Quantity, "price": cart.Price}}
	_, err = config.Cart_Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func Getallsellerdata() []models.Seller {
	filter := bson.D{}
	cursor, err := config.Seller_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var Seller []models.Seller
	for cursor.Next(context.Background()) {
		var seller models.Seller
		err := cursor.Decode(&seller)
		if err != nil {
			log.Fatal(err)
		}
		Seller = append(Seller, seller)
	}
	return Seller
}
func CreateSeller(seller models.Seller) bool {
	if seller.Password != seller.ConfirmPassword {
		return false
	}
	filter := bson.M{"selleremail": seller.Seller_Email}
	cursor, err := config.Seller_Collection.Find(context.Background(), filter)
	defer cursor.Close(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	if cursor.RemainingBatchLength() == 0 {
		seller.SellerId = GenerateUniqueCustomerID()
		_, err := config.Seller_Collection.InsertOne(context.Background(), seller)
		if err != nil {

			log.Fatal(err)

		}
		return true
	}
	return false
}
func Login(details models.Login) (string, bool, error) {
	var customer models.Customer

	filter := bson.M{"email": details.Email}
	err := config.Customer_Collection.FindOne(context.Background(), filter).Decode(&customer)
	if err != nil {
		// Handle the case where the user is not found
		return "", false, err
	}

	// Verify the password (You should use a secure password hashing library here)
	if customer.Password != details.Password {
		// Passwords don't match
		return "", false, nil
	}

	token, err := CreateToken(customer.Email, customer.CustomerId)
	if err != nil {
		return "", false, err

	}

	return token, true, nil
}
func Inventory(inventory models.Inventory) (bool, error) {
	filter := bson.M{"itemname": inventory.ItemName}
	cursor, err := config.Inventory_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	if cursor.RemainingBatchLength() == 0 {

		inventory1 := models.Inventory1{
			ItemCategory: inventory.ItemCategory,
			ItemName:     inventory.ItemName,
			Price:        inventory.Price,
			Quantity:     inventory.Quantity,
			Image:        inventory.Image,
		}
		var seller models.Seller
		err := config.Seller_Collection.FindOne(context.TODO(), bson.M{"sellerid": inventory.SellerId}).Decode(&seller)
		if err != nil {

			return false, err
		}
		inventory1.SellerName = seller.Seller_Name
		_, err1 := config.Inventory_Collection.InsertOne(context.Background(), inventory1)
		if err1 != nil {

			return false, err1
		}

		return true, nil
	} else {

		return false, fmt.Errorf("item name already exists")
	}

}

func Update(update models.Update) bool {
	if update.Collection == "seller" {
		filter := bson.M{"selleremail": update.IdName}
		update1 := bson.M{"$set": bson.M{update.Field: update.New_Value}}
		options := options.Update()
		_, err := config.Seller_Collection.UpdateOne(context.TODO(), filter, update1, options)
		if err != nil {
			return false
		}
		return true
	} else if update.Collection == "customer" {
		if update.Field == "phonenumber" || update.Field == "age" || update.Field == "pincode" {

			intValue, err := strconv.Atoi(update.New_Value)
			if err != nil {
				// Handle the error, e.g., return an error response or log it
			} else {
				update.New_Value = strconv.Itoa(intValue)
			}
			if !isValidNumber(update.New_Value) {
				return false
			}
			filter := bson.M{"email": update.IdName}
			update1 := bson.M{"$set": bson.M{update.Field: intValue}}
			options := options.Update()
			_, err1 := config.Customer_Collection.UpdateOne(context.TODO(), filter, update1, options)

			if err1 != nil {

				return false
			}

		}

		filter := bson.M{"email": update.IdName}
		update1 := bson.M{"$set": bson.M{update.Field: update.New_Value}}
		options := options.Update()
		_, err := config.Customer_Collection.UpdateOne(context.TODO(), filter, update1, options)
		fmt.Println("updated")
		if err != nil {
			return false
		}

		return true

	} else if update.Collection == "inventory" {
		if update.Field == "price" {
			// Check if New_Value is a valid integer
			intValue, err := strconv.Atoi(update.New_Value)
			if err != nil {
				// Handle the error, e.g., return an error response or log it
				return false
			}

			// Check if the input value is a valid number (numeric characters only)
			if !isValidNumber(update.New_Value) {
				return false
			}

			filter := bson.M{"itemname": update.IdName}
			update1 := bson.M{"$set": bson.M{update.Field: intValue}}
			options := options.Update()
			_, err1 := config.Inventory_Collection.UpdateOne(context.TODO(), filter, update1, options)
			if err1 != nil {

				return false
			}
			return true
		}

		filter := bson.M{"itemname": update.IdName}
		update1 := bson.M{"$set": bson.M{update.Field: update.New_Value}}
		options := options.Update()
		_, err := config.Inventory_Collection.UpdateOne(context.TODO(), filter, update1, options)
		if err != nil {

			return false
		}
		return true
	}

	return false
}

func Delete(delete models.Delete) bool {
	if delete.Collection == "customer" {
		filter := bson.M{"email": delete.IdValue}
		_, err := config.Customer_Collection.DeleteOne(context.Background(), filter)
		if err != nil {
			log.Fatal(err)
			return false
		}
		return true
	}
	if delete.Collection == "seller" {
		filter := bson.M{"selleremail": delete.IdValue}
		_, err := config.Seller_Collection.DeleteOne(context.Background(), filter)
		if err != nil {
			log.Fatal(err)
			return false
		}
		return true
	}
	if delete.Collection == "inventory" {
		filter := bson.M{"itemname": delete.IdValue}
		_, err := config.Inventory_Collection.DeleteOne(context.Background(), filter)
		if err != nil {
			log.Fatal(err)
			return false
		}
		return true
	}
	return true
}

func isValidNumber(s string) bool {
	numericRegex := regexp.MustCompile("^[0-9]+$")
	return numericRegex.MatchString(s)
}

func CheckSeller(check models.Login) (string, bool, error) {
	var seller models.Seller
	filter := bson.M{"selleremail": check.Email}
	config.Seller_Collection.FindOne(context.Background(), filter).Decode(&seller)
	if check.Password != seller.Password {
		return "", false, fmt.Errorf("InvalidPassword")
	}
	result, err := CreateToken(seller.Seller_Email, seller.SellerId)
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func DeleteProduct(delete models.DeleteProduct) bool {
	customerid, err := ExtractCustomerID(delete.Token, constants.SecretKey)
	if err != nil {
		log.Fatal(err)
		return false
	}
	filter1 := bson.M{"customerid": customerid}
	filter2 := bson.M{"name": delete.Name}
	combinedFilter := bson.M{
		"$and": []bson.M{filter1, filter2},
	}
	_, err = config.Cart_Collection.DeleteOne(context.Background(), combinedFilter)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

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

func DeleteProductBySeller(delete models.DeleteBySeller) int {
	filter := bson.M{"itemname": delete.ProductName}
	_, err := config.Inventory_Collection.DeleteOne(context.Background(), filter)
	if err != nil {

		return 0
	}
	return 1
}

func countdigits(number int) int {
	count := 0
	for number > 0 {
		count++
		number = number / 10
	}
	return int(count)
}

func Feedback(feedback models.Feedback) int {
	insertedid, err := config.Feedback_Collection.InsertOne(context.Background(), feedback)
	if err != nil {
		log.Fatal(err)
		return 3
	}
	fmt.Println(insertedid.InsertedID)
	return 1
}

func CustomerFeedback() []models.FeedbacktoAdmin {
	filter := bson.M{"role": "customer"}
	cursor, err := config.Feedback_Collection.Find(context.Background(), filter)
	var Feedback []models.FeedbacktoAdmin
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.Background()) {
		var feedback models.FeedbacktoAdmin
		err := cursor.Decode(&feedback)
		if err != nil {
			log.Fatal(err)
		}
		Feedback = append(Feedback, feedback)
	}
	return Feedback
}

func SellerFeedback() []models.FeedbacktoAdmin {
	filter := bson.M{"role": "seller"}
	cursor, err := config.Feedback_Collection.Find(context.Background(), filter)
	var Feedback []models.FeedbacktoAdmin
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.Background()) {
		var feedback models.FeedbacktoAdmin
		err := cursor.Decode(&feedback)
		if err != nil {
			log.Fatal(err)
		}
		Feedback = append(Feedback, feedback)
	}
	return Feedback
}

func Deletefeedback(delete models.FeedbacktoAdmin) int32 {
	filter1 := bson.M{"email": delete.Email}
	filter2 := bson.M{"feedback": delete.Feedback}
	combinedFilter := bson.M{
		"$and": []bson.M{filter1, filter2},
	}
	_, err := config.Feedback_Collection.DeleteMany(context.Background(), combinedFilter)
	if err != nil {
		return 0
	}
	return 1

}

func GetUser(id string) models.Address {
	var address models.Address
	filter1 := bson.M{"customerid": id}
	config.Customer_Collection.FindOne(context.Background(), filter1).Decode(&address)
	return address
}

func CustomerOrders(ItemsToBuy []models.Item, Data models.Address) {

	var order models.CustomerOrder
	order.Itemstobuy = ItemsToBuy
	order.Address = Data
	id, err := config.Buynow_Collection.InsertOne(context.Background(), order)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)

}

func Orders() []models.Customerorder {
	var Order []models.Customerorder
	filter := bson.M{}
	cursor, err := config.Buynow_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.Background()) {
		var order models.Customerorder
		err := cursor.Decode(&order)
		if err != nil {
			log.Fatal(err)
		}
		Order = append(Order, order)
	}
	fmt.Println(Order)
	return Order
}
func DeleteOrder(delete models.DeleteOrder) {
	fmt.Println(delete.Id)

	// Parse the ID string to an ObjectId
	objectID, err := primitive.ObjectIDFromHex(delete.Id)
	if err != nil {
		log.Fatal(err)
	}

	// Create a filter to match the ObjectId
	filter := bson.M{"_id": objectID}

	id, err := config.Buynow_Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id)
}


func CustomerOrder(token string) []models.Customerorder {
	id,err := ExtractCustomerID(token,constants.SecretKey)
	if err != nil{
		log.Fatal(err)
	}
	var customer models.Customer
	filter_customer := bson.M{"customerid":id}
	config.Customer_Collection.FindOne(context.Background(),filter_customer).Decode(&customer)
	var Order []models.Customerorder
	filter := bson.M{"address.phonenumber":customer.Phone_No}
	cursor, err := config.Buynow_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.Background()) {
		var order models.Customerorder
		err := cursor.Decode(&order)
		if err != nil {
			log.Fatal(err)
		}
		Order = append(Order, order)
	}
	fmt.Println(Order)
	return Order
}
