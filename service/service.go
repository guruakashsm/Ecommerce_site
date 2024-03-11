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
	var data models.Cart
	config.Cart_Collection.FindOne(context.Background(), filter).Decode(&data)
	if data.Quantity > cart.Quantity {
		var inventory models.Inventory
		filter := bson.M{"itemname": cart.Name}
		config.Inventory_Collection.FindOne(context.Background(), filter).Decode(&inventory)
		quantity := data.Quantity - cart.Quantity
		inventory.Stock_Available = inventory.Stock_Available + quantity
		update := bson.M{"$set": bson.M{"sellerquantity": inventory.Stock_Available}}
		config.Inventory_Collection.UpdateOne(context.Background(), filter, update)
	}
	if data.Quantity < cart.Quantity {
		var inventory models.Inventory
		filter := bson.M{"itemname": cart.Name}
		config.Inventory_Collection.FindOne(context.Background(), filter).Decode(&inventory)
		quantity := cart.Quantity - data.Quantity
		inventory.Stock_Available = inventory.Stock_Available - quantity
		update := bson.M{"$set": bson.M{"sellerquantity": inventory.Stock_Available}}
		config.Inventory_Collection.UpdateOne(context.Background(), filter, update)
	}

	update := bson.M{"$set": bson.M{"name": cart.Name, "quantity": cart.Quantity, "totalprice": cart.Price, "price": cart.Price / float64(cart.Quantity)}}
	result, err := config.Cart_Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println(err)
	}
	return result

}

func Cart(customerid string) []models.Cart {

	filter := bson.M{"customerid": customerid}
	cursor, err := config.Cart_Collection.Find(context.Background(), filter)
	if err != nil {

		log.Println(err)
	}
	defer cursor.Close(context.Background())
	var Cart []models.Cart
	for cursor.Next(context.Background()) {
		var cart models.Cart
		err := cursor.Decode(&cart)
		if err != nil {

			log.Println(err)
		}
		Cart = append(Cart, cart)
	}
	return Cart
}
func Search(productName string) []models.Inventory1 {

	filter := bson.M{"itemcategory": productName}
	cursor, err := config.Inventory_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	defer cursor.Close(context.Background())
	var Inventory []models.Inventory1
	for cursor.Next(context.Background()) {
		var inventory models.Inventory1
		err := cursor.Decode(&inventory)
		if inventory.Stock_Available <= 0 {
			// filter := bson.M{"itemname":inventory.ItemName}
			// _,err:=config.Inventory_Collection.DeleteOne(context.Background(),filter)
			// if err != nil {
			// 	log.Println(err)
			// }
			continue
		}
		if err != nil {
			log.Println(err)
		}

		Inventory = append(Inventory, inventory)
	}
	return Inventory
}
func Getalldata() []models.Customer {
	filter := bson.D{}
	cursor, err := config.Customer_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	defer cursor.Close(context.Background())
	var Profiles []models.Customer
	for cursor.Next(context.Background()) {
		var profile models.Customer
		err := cursor.Decode(&profile)
		if err != nil {
			log.Println(err)
		}
		Profiles = append(Profiles, profile)
	}
	return Profiles
}
func Getinventorydata() []models.Inventory {
	filter := bson.D{}
	cursor, err := config.Inventory_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	defer cursor.Close(context.Background())
	var Inventorydata []models.Inventory
	for cursor.Next(context.Background()) {
		var inventory models.Inventory
		err := cursor.Decode(&inventory)
		if err != nil {
			log.Println(err)
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

	filter := bson.M{"email": profile.Email}
	existingCustomer := config.Customer_Collection.FindOne(context.Background(), filter)

	// Check if the document already exists
	if existingCustomer.Err() == nil {
		var existingProfile models.Customer // Replace with the actual type of your document
		if err := existingCustomer.Decode(&existingProfile); err != nil {
			log.Println(err)
			return 0
		}

		// Check if the email is verified
		if existingProfile.IsEmailVerified {
			return 2
		} else {
			profile.CustomerId = GenerateUniqueCustomerID()
			profile.IsEmailVerified = false
			profile.WrongInput = 0
			profile.VerificationString = GenerateOTP(6)
			_, updateErr := config.Customer_Collection.ReplaceOne(context.Background(), filter, profile)
			if updateErr != nil {
				return 0
			}
			go SendEmailforCustomerVerification(profile.Email, profile.VerificationString, "GURU")

			return 1
		}
	} else if existingCustomer.Err() == mongo.ErrNoDocuments {

		profile.CustomerId = GenerateUniqueCustomerID()
		profile.IsEmailVerified = false
		profile.BlockedUser = false
		profile.WrongInput = 0
		profile.VerificationString = GenerateOTP(6)

		inserted, insertErr := config.Customer_Collection.InsertOne(context.Background(), profile)
		if insertErr != nil {
			log.Println(insertErr)
			return 0
		}
		go SendEmailforCustomerVerification(profile.Email, profile.VerificationString, profile.Name)
		fmt.Println("Inserted", inserted.InsertedID)
		return 1
	} else {
		log.Println(existingCustomer.Err())
		return 0
	}

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
		log.Println(err)
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
		var inventory models.Inventory
		filter := bson.M{"itemname": addtocart.Name}
		config.Inventory_Collection.FindOne(context.Background(), filter).Decode(&inventory)
		inventory.Stock_Available--
		update := bson.M{"$set": bson.M{"sellerquantity": inventory.Stock_Available}}
		config.Inventory_Collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			log.Println(err)
			return false
		}

		return true
	}

	// Item already exists, update its quantity
	var cart addcart
	for cursor.Next(context.Background()) {
		err = cursor.Decode(&cart)

		if err != nil {
			log.Println(err)
		}
	}
	// Item already exists, update its quantity
	cart.Quantity++
	cart.Price = cart.Price + addtocart.Price
	// Use the UpdateOne method to increment the quantity
	update := bson.M{"$set": bson.M{"quantity": cart.Quantity, "price": cart.Price}}
	_, err = config.Cart_Collection.UpdateOne(context.Background(), filter, update)
	var inventory models.Inventory
	filter1 := bson.M{"itemname": addtocart.Name}
	config.Inventory_Collection.FindOne(context.Background(), filter1).Decode(&inventory)
	inventory.Stock_Available--
	fmt.Println(inventory.Stock_Available)
	update1 := bson.M{"$set": bson.M{"sellerquantity": inventory.Stock_Available}}
	config.Inventory_Collection.UpdateOne(context.Background(), filter1, update1)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func Getallsellerdata() []models.Seller {
	filter := bson.D{}
	cursor, err := config.Seller_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	defer cursor.Close(context.Background())
	var Seller []models.Seller
	for cursor.Next(context.Background()) {
		var seller models.Seller
		err := cursor.Decode(&seller)
		if err != nil {
			log.Println(err)
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
		log.Println(err)
	}
	if cursor.RemainingBatchLength() == 0 {
		seller.SellerId = GenerateUniqueCustomerID()
		seller.BlockedUser = false
		seller.WrongInput = 0
		_, err := config.Seller_Collection.InsertOne(context.Background(), seller)
		if err != nil {

			log.Println(err)

		}
		go SendSellerInvitation(seller.Seller_Email, seller.Seller_Name, seller.Password, "https://anon.up.railway.app/seller/")
		return true
	}
	return false
}
func Login(details models.Login) (string, error, int) {
	var customer models.Customer

	filter := bson.M{"email": details.Email}
	err := config.Customer_Collection.FindOne(context.Background(), filter).Decode(&customer)
	if err != nil {
		return "User not found", err, 0
	}
	if customer.WrongInput == 10 {
		return "Too many no of try", nil, 0
	}
	if !customer.IsEmailVerified {
		return "Please verify your email", nil, 0
	}
	if customer.BlockedUser {
		return "Your ID has been Blocked", nil, 0
	}
	if customer.Password != details.Password {
		customer.WrongInput++
		update := bson.M{"$set": bson.M{"wronginput": customer.WrongInput}}
		config.Customer_Collection.UpdateOne(context.Background(), filter, update)
		return "Wrong Password", nil, 0
	}

	token, err := CreateToken(customer.Email, customer.CustomerId)
	if err != nil {
		return "Internal server error", err, 0

	}

	return token, nil, 1
}
func Inventory(inventory models.Inventory) (bool, error) {
	filter := bson.M{"itemname": inventory.ItemName}
	cursor, err := config.Inventory_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
		return false, err
	}
	if cursor.RemainingBatchLength() == 0 {

		inventory1 := models.Inventory1{
			ItemCategory:    inventory.ItemCategory,
			ItemName:        inventory.ItemName,
			Price:           inventory.Price,
			Quantity:        inventory.Quantity,
			Image:           inventory.Image,
			Stock_Available: inventory.Stock_Available,
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
		go SendEditDataNotification(update.IdName,update.Field,update.New_Value)
		return true
	} else if update.Collection == "customer" {
		if update.Field == "phonenumber" || update.Field == "age" || update.Field == "pincode" {

			intValue, err := strconv.Atoi(update.New_Value)
			if err != nil {
				return false
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
		go SendEditDataNotification(update.IdName,update.Field,update.New_Value)
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
			log.Println(err)
			return false
		}
		return true
	}
	if delete.Collection == "seller" {
		filter := bson.M{"selleremail": delete.IdValue}
		_, err := config.Seller_Collection.DeleteOne(context.Background(), filter)
		if err != nil {
			log.Println(err)
			return false
		}
		return true
	}
	if delete.Collection == "inventory" {
		filter := bson.M{"itemname": delete.IdValue}
		_, err := config.Inventory_Collection.DeleteOne(context.Background(), filter)
		if err != nil {
			log.Println(err)
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
	if seller.BlockedUser == true {
		return "", false, fmt.Errorf("Your ID Has been blocked by admin")
	}
	if seller.WrongInput == 10 {
		return "", false, fmt.Errorf("To many no of attempts")
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
		log.Println(err)
		return false
	}
	filter1 := bson.M{"customerid": customerid}
	filter2 := bson.M{"name": delete.Name}
	combinedFilter := bson.M{
		"$and": []bson.M{filter1, filter2},
	}
	filter3 := bson.M{"itemname": delete.Name}
	var data models.Inventory1
	config.Inventory_Collection.FindOne(context.Background(), filter3).Decode(&data)
	fmt.Println(data)
	delete.Quantity = delete.Quantity + int(data.Stock_Available)
	fmt.Println(delete.Quantity)
	update1 := bson.M{"$set": bson.M{"sellerquantity": delete.Quantity}}
	_, err = config.Inventory_Collection.UpdateOne(context.Background(), filter3, update1)
	if err != nil {
		log.Println(err)
	}
	_, err = config.Cart_Collection.DeleteOne(context.Background(), combinedFilter)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
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
		log.Println(err)
	}
	for cursor.Next(context.Background()) {
		var feedback models.FeedbacktoAdmin
		err := cursor.Decode(&feedback)
		if err != nil {
			log.Println(err)
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
		log.Println(err)
	}
	for cursor.Next(context.Background()) {
		var feedback models.FeedbacktoAdmin
		err := cursor.Decode(&feedback)
		if err != nil {
			log.Println(err)
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
		log.Println(err)
	}
	fmt.Println(id)

}

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
func DeleteOrder(delete models.DeleteOrder) {
	fmt.Println(delete.Id)

	// Parse the ID string to an ObjectId
	objectID, err := primitive.ObjectIDFromHex(delete.Id)
	if err != nil {
		log.Println(err)
	}

	// Create a filter to match the ObjectId
	filter := bson.M{"_id": objectID}

	id, err := config.Buynow_Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(id)
}

func CustomerOrder(token string) []models.Customerorder {
	id, err := ExtractCustomerID(token, constants.SecretKey)
	if err != nil {
		log.Println(err)
	}
	var customer models.Customer
	filter_customer := bson.M{"customerid": id}
	config.Customer_Collection.FindOne(context.Background(), filter_customer).Decode(&customer)
	var Order []models.Customerorder
	filter := bson.M{"address.phonenumber": customer.Phone_No}
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

func Buynow(id string) {
	filter := bson.M{"customerid": id}
	_, err := config.Cart_Collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
}

func Itemstobuy(id string) []models.Item {
	filter := bson.M{"customerid": id}
	cursor, err := config.Cart_Collection.Find(context.Background(), filter)
	if err != nil {

		log.Println(err)
	}
	var Item []models.Item
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var item models.Item
		err := cursor.Decode(&item)
		if err != nil {

			log.Println(err)
		}
		Item = append(Item, item)
	}
	//fmt.Println(Item)
	return Item
}

func TotalAmount(id string) float64 {
	filter := bson.M{"customerid": id}
	cursor, err := config.Cart_Collection.Find(context.Background(), filter)
	if err != nil {

		log.Println(err)
	}
	var Cart float64
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var cart models.Cart
		err := cursor.Decode(&cart)
		if err != nil {

			log.Println(err)
		}
		if cart.TotalPrice == 0 {
			Cart = Cart + cart.Price
		} else {
			Cart = Cart + cart.TotalPrice
		}
	}
	var currentValues models.Sales
	err = config.Sales_Collection.FindOne(context.Background(), bson.M{}).Decode(&currentValues)
	if err != nil {
		log.Println(err)
	}
	// Update the values
	updatedTotalSalesAmount := currentValues.TotalSalesAmount + int(Cart)
	updatedTotalNoOfSales := currentValues.TotalNoOfSales + 1

	// Update the document in the collection
	_, err = config.Sales_Collection.UpdateOne(
		context.Background(),
		bson.M{},
		bson.M{
			"$set": bson.M{
				"totalsalesamount": updatedTotalSalesAmount,
				"totalnoofsales":   updatedTotalNoOfSales,
			},
		},
	)
	return Cart
}

func Validatetoken(token string) bool {
	_, err := ExtractCustomerID(token, constants.SecretKey)
	return err == nil
}

func AdminLoginCheck(login *models.AdminData) (string, int) {

	var correctdata models.AdminData
	filter := bson.M{"email": login.Email}
	err := config.Admin_Collection.FindOne(context.Background(), filter).Decode(&correctdata)
	if err != nil {
		return "", 0
	}
	if correctdata.WrongInput == 4 {
		return "", 1
	}
	if correctdata.IP_Address == login.IP_Address {
		return "", 3
	}
	if correctdata.Password != login.Password {
		correctdata.WrongInput++
		update := bson.M{"$set": bson.M{"wronginput": correctdata.WrongInput}}
		config.Admin_Collection.UpdateOne(context.Background(), filter, update)
		return "", 2
	}

	if !ValidateOTP(login.TOTP, correctdata.SecretKey) {
		correctdata.WrongInput++
		update := bson.M{"$set": bson.M{"wronginput": correctdata.WrongInput}}
		config.Admin_Collection.UpdateOne(context.Background(), filter, update)
		return "", 4
	}
	idString := correctdata.Id.Hex()
	objectID, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		log.Println(err)
	}
	token, err := CreateToken(login.Email, string(objectID.String()))
	if err != nil {
		return "", 5
	}
	log.Println(token)
	update := bson.M{"$set": bson.M{"token": token, "wronginput": 0}}
	config.Admin_Collection.UpdateOne(context.Background(), filter, update)
	return token, 5

}

func AdminNeededData() models.AdminPageData {
	var adminpagedata models.AdminPageData
	var sales models.Sales
	adminpagedata.ProductCount, _ = config.Inventory_Collection.CountDocuments(context.Background(), bson.D{})

	adminpagedata.UserCount, _ = config.Customer_Collection.CountDocuments(context.Background(), bson.D{})

	adminpagedata.SellerCount, _ = config.Seller_Collection.CountDocuments(context.Background(), bson.D{})

	config.Sales_Collection.FindOne(context.Background(), bson.M{}).Decode(&sales)

	adminpagedata.SalesCount = int64(sales.TotalNoOfSales)

	adminpagedata.TotalSalesAmount = int32(sales.TotalSalesAmount)

	return adminpagedata
}

func GetWorkerdata() []models.Workers {
	var workers []models.Workers

	filter := bson.M{}
	cursor, err := config.Worker_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	for cursor.Next(context.Background()) {
		var worker models.Workers
		err := cursor.Decode(&worker)
		if err != nil {
			log.Println(err)
		}
		workers = append(workers, worker)
	}
	return workers
}

func GetFeedBacks() []models.FeedbacktoAdmin {
	filter := bson.M{}
	cursor, err := config.Feedback_Collection.Find(context.Background(), filter)
	var Feedback []models.FeedbacktoAdmin
	if err != nil {
		log.Println(err)
	}

	for cursor.Next(context.Background()) {
		var feedback models.FeedbacktoAdmin
		err := cursor.Decode(&feedback)
		if err != nil {
			log.Println(err)
		}
		Feedback = append(Feedback, feedback)
	}
	return Feedback
}

func CreateWorker(worker models.Workers) string {
	filter := bson.M{"email": worker.Email}
	result := config.Worker_Collection.FindOne(context.Background(), filter)
	if result.Err() == nil {
		return "User Already Exists"
	}
	if result.Err() != nil && result.Err() != mongo.ErrNoDocuments {
		return "Error in Query: " + result.Err().Error()
	}
	_, err := config.Worker_Collection.InsertOne(context.Background(), worker)
	if err != nil {
		return "Error in Creating: " + err.Error()
	}
	return "Created Successfully"
}

func CreateAdmin(admin models.AdminSignup) (string, string) {
	filter := bson.M{"email": admin.Email}

	result := config.Admin_Collection.FindOne(context.Background(), filter)
	if result.Err() == nil {
		return "User Already Exists", ""
	}
	if result.Err() != nil && result.Err() != mongo.ErrNoDocuments {
		return "Error in Query: " + result.Err().Error(), ""
	}
	key, err := GenerateSecret()
	if err != nil {
		return "Error In Generating TOTP", ""
	}
	var AdminData models.Admin
	AdminData.Email = admin.Email
	AdminData.Password = admin.Password
	AdminData.IP_Address = admin.IP
	AdminData.SecretKey = key
	AdminData.Token = ""
	AdminData.WrongInput = 0
	_, err = config.Admin_Collection.InsertOne(context.Background(), AdminData)
	if err != nil {
		return "Error in Creating: " + err.Error(), ""
	}
	go SendAdminInvitation(admin.Email, admin.AdminName, admin.Password, "https://anon.up.railway.app/admin/", admin.IP, key)
	return "Created Successfully", key
}

func GetData(data models.Getdata) (*models.ReturnData, error) {
	var returndata models.ReturnData

	if data.Collection == "customer" {
		log.Println("In customer")
		var profile models.Customer
		filter := bson.M{"email": data.Id}
		err := config.Customer_Collection.FindOne(context.Background(), filter).Decode(&profile)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		returndata.Name = profile.Name
		returndata.CustomerId = profile.CustomerId
		returndata.Address = profile.Address
		returndata.Email = profile.Email
		returndata.Phone_No = profile.Phone_No
		returndata.Password = profile.Password
		returndata.IsEmailVerified = profile.IsEmailVerified
		returndata.BlockedUser = profile.BlockedUser
		returndata.WrongInput = profile.WrongInput
		return &returndata, nil

	} else if data.Collection == "seller" {
		log.Println("In seller")
		var profile models.Seller
		filter := bson.M{"selleremail": data.Id}
		log.Println()
		err := config.Seller_Collection.FindOne(context.Background(), filter).Decode(&profile)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		returndata.Seller_Name = profile.Seller_Name
		returndata.Phone_No = profile.Phone_No
		returndata.Address = profile.Address
		returndata.Password = profile.Password
		returndata.SellerId = profile.SellerId
		returndata.Seller_Email = profile.Seller_Email
		returndata.Seller_Name = profile.Seller_Name
		returndata.Image = profile.Image
		returndata.BlockedUser = profile.BlockedUser
		returndata.WrongInput = profile.WrongInput
		return &returndata, nil
	} else if data.Collection == "inventory" {
		log.Println("In inventory")
		var profile models.Inventory1
		filter := bson.M{"itemname": data.Id}
		err := config.Inventory_Collection.FindOne(context.Background(), filter).Decode(&profile)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		returndata.ItemCategory = profile.ItemCategory
		returndata.ItemName = profile.ItemName
		returndata.Quantity = profile.Quantity
		returndata.Seller_Name = profile.SellerName
		returndata.Price = profile.Price
		returndata.Stock_Available = profile.Stock_Available
		returndata.Image = profile.Image
		return &returndata, nil
	} else if data.Collection == "worker" {
		log.Println("In worker")
		var profile models.Workers
		filter := bson.M{"email": data.Id}
		err := config.Worker_Collection.FindOne(context.Background(), filter).Decode(&profile)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		returndata.Email = profile.Email
		returndata.No = profile.No
		returndata.Role = profile.Role
		returndata.Status = profile.Status
		returndata.UserName = profile.UserName
		returndata.Salary = profile.Salary
		returndata.Image = profile.Image
		return &returndata, nil
	}
	return nil, nil

}

func AddEvent(upload models.UploadCalender) error {
	_, err := config.Calender_Collection.InsertOne(context.Background(), upload)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetEvent(GetData models.GetCalender) ([]models.UploadCalender, error) {
	filter := bson.M{"email": GetData.AdminEmail}
	cursor, err := config.Calender_Collection.Find(context.Background(), filter)
	var Data []models.UploadCalender
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var data models.UploadCalender
		err := cursor.Decode(&data)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		Data = append(Data, data)
	}
	return Data, nil
}

func EmailVerification(data models.VerifyEmail) (string, error) {
	filter := bson.M{"email": data.Email}
	var customer models.Customer
	err := config.Customer_Collection.FindOne(context.Background(), filter).Decode(&customer)
	if err != nil {
		return "", err
	}
	if customer.VerificationString == data.VerificationString {
		update := bson.M{"$set": bson.M{"isemailverified": true}}
		filter := bson.M{"email": data.Email}

		_, err := config.Customer_Collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return "", err
		}
		go SendThankYouEmail(customer.Email, customer.Name)
		return "Signup Successful", nil
	} else {
		return "Wrong OTP", nil
	}

}

func Block(data models.Block) (string, error) {
	if data.Collection == "customer" {
		var customer models.Customer
		filter := bson.M{"email": data.Email}
		err := config.Customer_Collection.FindOne(context.Background(), filter).Decode(&customer)
		if err != nil {
			log.Println(err)
			return "No result Found", err
		}
		message := ""
		if customer.BlockedUser {
			customer.BlockedUser = false
			message = "Customer has been Unblocked"
		} else {
			customer.BlockedUser = true
			message = "Customer has been Blocked"
		}
		update := bson.M{"$set": bson.M{"blockeduser": customer.BlockedUser}}
		_, err = config.Customer_Collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			log.Println(err)
			return "Can't Update Data", err
		}
		go SendBlockingNotification(customer.Email,customer.Name,"Due to improper behaviour")
		return message, nil
	} else if data.Collection == "seller" {
		var Seller models.Seller
		filter := bson.M{"selleremail": data.Email}
		err := config.Seller_Collection.FindOne(context.Background(), filter).Decode(&Seller)
		if err != nil {
			log.Println(err)
			return "No result Found", err
		}
		message := ""
		if Seller.BlockedUser{
			Seller.BlockedUser = false
			message = "Seller has been Unblocked"
		} else {
			Seller.BlockedUser = true
			message = "Seller has been Blocked"
		}
		update := bson.M{"$set": bson.M{"blockeduser": Seller.BlockedUser}}
		_, err = config.Seller_Collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			log.Println(err)
			return "Can't Update Data", err
		}
		go SendBlockingNotification(Seller.Seller_Email,Seller.Seller_Name,"Due to improper behaviour")
		return message, nil

	}
	return "Invalid Collection", nil
}
