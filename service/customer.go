package service

import (
	"context"
	"ecommerce/config"
	"ecommerce/constants"
	"ecommerce/models"
	"fmt"
	"log"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Create Customer
func CreateCustomer(profile models.Customer) int {
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

// Login Customer
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

// Add To Cart
func Addtocart(addtocart models.Addtocart) (string, error) {
	id, err := ExtractCustomerID(addtocart.Token, constants.SecretKey)
	if err != nil {
		log.Println(err)
		return "Token Expired (or) Invalid", nil
	}
	filter := bson.D{
		{Key: "$and", Value: []interface{}{
			bson.D{{Key: "customerid", Value: id}},
			bson.D{{Key: "productname", Value: addtocart.ProductName}},
		}},
	}

	cursor, err := config.Cart_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
		return "Internal Server Error", err
	}

	var inventoryData models.Inventory
	inventoryfilter := bson.M{"itemname": addtocart.ProductName}
	err = config.Inventory_Collection.FindOne(context.Background(), inventoryfilter).Decode(&inventoryData)
	if err != nil {
		log.Println(err)
		return "Error in Fetching", err
	}
	defer cursor.Close(context.Background())

	if cursor.RemainingBatchLength() == 0 {
		// Item not found, so insert a new item with quantity 1
		cart := models.Addcart{
			CustomerId:   id,
			ProductName:  addtocart.ProductName,
			Price:        inventoryData.Price,
			Quantity:     1,
			Image:        inventoryData.Image,
			SellerID:     inventoryData.SellerId,
			SellerName:   inventoryData.SellerName,
			ItemCategory: inventoryData.ItemCategory,
		}
		_, err := config.Cart_Collection.InsertOne(context.Background(), cart)
		if err != nil {
			log.Println(err)
			return "Error in Inserting", err
		}

	} else {
		// Item already exists, update its quantity
		var cart models.Addcart
		for cursor.Next(context.Background()) {
			err = cursor.Decode(&cart)

			if err != nil {

				return "Error in Decode", err
			}
		}
		// Item already exists, update its quantity
		cart.Quantity++
		cart.Price = cart.Price + inventoryData.Price
		// Use the UpdateOne method to increment the quantity

		update := bson.M{"$set": bson.M{"quantity": cart.Quantity, "price": cart.Price}}
		_, err = config.Cart_Collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return "Error In Updating", err
		}

	}

	inventoryData.Stock_Available--
	update := bson.M{"$set": bson.M{"sellerquantity": inventoryData.Stock_Available}}
	_, err = config.Inventory_Collection.UpdateOne(context.Background(), inventoryfilter, update)
	if err != nil {
		return "Error In Updating", err
	}

	return "Item Added to Cart", nil
}

// Search and get Products
func Search(productName string) []models.Inventory {

	filter := bson.M{"itemcategory": productName}
	cursor, err := config.Inventory_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	defer cursor.Close(context.Background())
	var Inventory []models.Inventory
	for cursor.Next(context.Background()) {
		var inventory models.Inventory
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

// Get All items in Cart
func GetAllItemsinCart(token models.Token) ([]models.Addcart, string, error) {
	id, err := ExtractCustomerID(token.Token, constants.SecretKey)
	if err != nil {
		return nil, "Login Expired", err
	}
	filter := bson.M{"customerid": id}
	cursor, err := config.Cart_Collection.Find(context.Background(), filter)
	if err != nil {
		return nil, "No result Found", err
	}
	defer cursor.Close(context.Background())
	var Cart []models.Addcart
	for cursor.Next(context.Background()) {
		var cart models.Addcart
		err := cursor.Decode(&cart)
		if err != nil {
			log.Println(err)
			return nil, "Error in Decode", err
		}
		Cart = append(Cart, cart)
	}
	return Cart, "Success", nil
}

// Update Cart
func UpdateCart(cart models.Addcart) *mongo.UpdateResult {
	filter := bson.D{
		{Key: "$and", Value: []interface{}{
			bson.D{{Key: "customerid", Value: cart.CustomerId}},
			bson.D{{Key: "name", Value: cart.ProductName}},
		}},
	}
	var data models.Addcart
	config.Cart_Collection.FindOne(context.Background(), filter).Decode(&data)
	if data.Quantity > cart.Quantity {
		var inventory models.Inventory
		filter := bson.M{"itemname": cart.ProductName}
		config.Inventory_Collection.FindOne(context.Background(), filter).Decode(&inventory)
		quantity := data.Quantity - cart.Quantity
		inventory.Stock_Available = inventory.Stock_Available + quantity
		update := bson.M{"$set": bson.M{"sellerquantity": inventory.Stock_Available}}
		config.Inventory_Collection.UpdateOne(context.Background(), filter, update)
	}
	if data.Quantity < cart.Quantity {
		var inventory models.Inventory
		filter := bson.M{"itemname": cart.ProductName}
		config.Inventory_Collection.FindOne(context.Background(), filter).Decode(&inventory)
		quantity := cart.Quantity - data.Quantity
		inventory.Stock_Available = inventory.Stock_Available - quantity
		update := bson.M{"$set": bson.M{"sellerquantity": inventory.Stock_Available}}
		config.Inventory_Collection.UpdateOne(context.Background(), filter, update)
	}

	update := bson.M{"$set": bson.M{"name": cart.ProductName, "quantity": cart.Quantity, "totalprice": cart.Price, "price": cart.Price / float64(cart.Quantity)}}
	result, err := config.Cart_Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println(err)
	}
	return result

}

// Delete Products in Cart
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
	var data models.Inventory
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

// Get Customer Address
func GetUser(id string) models.Address {
	var address models.Address
	filter1 := bson.M{"customerid": id}
	config.Customer_Collection.FindOne(context.Background(), filter1).Decode(&address)
	return address
}

// Add Ordered Items to Db
func CustomerOrders(ItemsToBuy []models.Item, Data models.Address) {

	var order models.Customerorder
	order.Itemstobuy = ItemsToBuy
	order.Address = Data
	id, err := config.Buynow_Collection.InsertOne(context.Background(), order)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(id)

}

// Delete Orders
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

//Display Customer Order
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

//Delete Items In Cart
func DeleteItemsInCart(id string) {
	filter := bson.M{"customerid": id}
	_, err := config.Cart_Collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
}

//Fetch itmes from Cart
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


//Display Total Amount in Cart
func TotalAmount(id string) float64 {
	filter := bson.M{"customerid": id}
	cursor, err := config.Cart_Collection.Find(context.Background(), filter)
	if err != nil {

		log.Println(err)
	}
	var Cart float64
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var cart models.Addcart
		err := cursor.Decode(&cart)
		if err != nil {

			log.Println(err)
		}
		// if cart.TotalPrice == 0 {
		// 	Cart = Cart + cart.Price
		// } else {
		// 	Cart = Cart + cart.TotalPrice
		// }
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