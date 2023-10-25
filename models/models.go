package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	CustomerId      string `json:"customerid" bson:"customerid"`
	Email           string `json:"email" bson:"email"`
	Name            string `json:"name" bson:"name"`
	Phone_No        int    `json:"phonenumber" bson:"phonenumber"`
	Age             int    `json:"age" bson:"age"`
	Password        string `json:"password" bson:"password"`
	ConfirmPassword string `json:"confirmpassword" bson:"confirmpassword"`
	FirstName       string `json:"firstname" bson:"firstname"`
	LastName        string `json:"lastname" bson:"lastname"`
	House_No        string `json:"houseno" bson:"houseno"`
	Street_Name     string `json:"streetname" bson:"streetname"`
	City            string `json:"city" bson:"city"`
	Pincode         int64  `json:"pincode" bson:"pincode"`
}
type Address struct {
	FirstName   string `json:"firstname" bson:"firstname"`
	LastName    string `json:"lastname" bson:"lastname"`
	Phone_No    int    `json:"phonenumber" bson:"phonenumber"`
	House_No    string `json:"houseno" bson:"houseno"`
	Street_Name string `json:"streetname" bson:"streetname"`
	City        string `json:"city" bson:"city"`
	Pincode     int64  `json:"pincode" bson:"pincode"`
}
type Inventory struct {
	SellerId        string  `json:"sellerid" bson:"sellerid"`
	ItemCategory    string  `json:"itemcategory" bson:"itemcategory"`
	ItemName        string  `json:"itemname" bson:"itemname"`
	Price           float64 `json:"price" bson:"price"`
	Quantity        string  `json:"quantity" bson:"quantity"`
	Stock_Available int32   `json:"sellerquantity" bson:"sellerquantity"`
	Image           string  `json:"image" bson:"image"`
}
type Inventory1 struct {
	SellerName      string  `json:"sellername" bson:"sellername"`
	ItemCategory    string  `json:"itemcategory" bson:"itemcategory"`
	ItemName        string  `json:"itemname" bson:"itemname"`
	Price           float64 `json:"price" bson:"price"`
	Quantity        string  `json:"quantity" bson:"quantity"`
	Image           string  `json:"image" bson:"image"`
	Stock_Available int32   `json:"sellerquantity" bson:"sellerquantity"`
}

type Addtocart1 struct {
	CustomerId     string  `json:"customerid" bson:"customerid"`
	Name           string  `json:"name" bson:"name"`
	Price          float64 `json:"price" bson:"price"`
	SellerQuantity int32   `json:"sellerquantity" bson:"sellerquantity"`
}
type Seller struct {
	SellerId        string `json:"sellerid" bson:"sellerid"`
	Seller_Name     string `json:"sellername" bson:"sellername"`
	Seller_Email    string `json:"selleremail" bson:"selleremail"`
	Password        string `json:"password" bson:"password"`
	ConfirmPassword string `json:"confirmpassword" bson:"confirmpassword"`
	Phone_No        int    `json:"phoneno" bson:"phoneno"`
	Address         string `json:"address" bson:"address"`
}
type Delete struct {
	Collection string `json:"collection" bson:"collection"`
	IdValue    string `json:"idValue" bson:"idValue"`
}
type Addtocart struct {
	Token          string  `json:"token" bson:"token"`
	Name           string  `json:"name" bson:"name"`
	Price          float64 `json:"price" bson:"price"`
	Sellerquantity int32   `json:"quantity" bson:"quantity"`
}
type Login struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
type Cart struct {
	CustomerId string  `json:"token"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Quantity   int32   `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}
type Orders struct {
	Item_id    string  `json:"itemid" bson:"itemid"`
	Item_Name  string  `json:"itemname" bson:"itemname"`
	Quantity   int64   `json:"quantity" bson:"quantity"`
	Total_Cost float64 `json:"totalcost" bson:"totalcost"`
}
type Customer_Response struct {
	Id              primitive.ObjectID `json:"_id" bson:"_id"`
	Name            string             `json:"name" bson:"name"`
	Phone_No        int                `json:"phonenumber" bson:"phonenumber"`
	Age             int                `json:"age" bson:"age"`
	Password        string             `json:"password" bson:"password"`
	ConfirmPassword string             `json:"confirmpassword" bson:"confirmpassword"`
	FirstName       string             `json:"firstname" bson:"firstname"`
	LastName        string             `json:"lastname" bson:"lastname"`
	House_No        string             `json:"houseno" bson:"houseno"`
	Street_Name     string             `json:"streetname" bson:"streetname"`
	City            string             `json:"city" bson:"city"`
	Pincode         int64              `json:"pincode" bson:"pincode"`
}

type Update struct {
	Collection string `json:"collection" bson:"collection"`
	IdName     string `json:"email" bson:"email"`
	Field      string `json:"field" bson:"field"`
	New_Value  string `json:"newvalue" bson:"newvalue"`
}
type DeleteBySeller struct {
	ProductName string `json:"productname" bson:"productname"`
}
type DeleteProduct struct {
	Token    string `json:"token" bson:"token"`
	Name     string `json:"name" bson:"name"`
	Quantity int    `json:"quantity" bson:"quantity"`
}
type UpdateProduct struct {
	ProductName string `json:"productname" bson:"productname"`
	Attribute   string `json:"attribute" bson:"attribute"`
	New_Value   int32  `json:"newvalue" bson:"newvalue"`
}

type Feedback struct {
	Role     string `json:"role" bson:"role"`
	Email    string `json:"email" bson:"email"`
	Feedback string `json:"feedback" bson:"feedback"`
}

type FeedbacktoAdmin struct {
	Email    string `json:"email" bson:"email"`
	Feedback string `json:"feedback" bson:"feedback"`
}

type BuyNow struct {
	Token       string  `json:"token"`
	TotalAmount float64 `json:"totalAmount"`
	ItemsToBuy  []Item  `json:"itemsToBuy"`
}

type Item struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type TotalAmount struct {
	TotalAmount float64 `json:"totalamount" bson:"totalamount"`
}

type CustomerOrder struct {
	Itemstobuy []Item  `json:"itemstobuy" bson:"itemstobuy"`
	Address    Address `json:"address" bson:"address"`
}

type Customerorder struct {
	Id         string  `json:"_id" bson:"_id"`
	Itemstobuy []Item  `json:"itemstobuy" bson:"itemstobuy"`
	Address    Address `json:"address" bson:"address"`
}
type DeleteOrder struct {
	Id string `json:"_id" bson:"_id"`
}
type Token struct {
	Token string `json:"token" bson:"token"`
}
