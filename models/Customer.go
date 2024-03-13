package models

// Sign Up Customer
type Customer struct {
	CustomerId         string `json:"customerid" bson:"customerid"`
	Email              string `json:"email" bson:"email"`
	Name               string `json:"name" bson:"name"`
	Phone_No           int    `json:"phonenumber" bson:"phonenumber"`
	Password           string `json:"password" bson:"password"`
	ConfirmPassword    string `json:"confirmpassword" bson:"confirmpassword"`
	Address            string `json:"address" bson:"address"`
	IsEmailVerified    bool   `json:"isemailverified" bson:"isemailverified"`
	WrongInput         int    `json:"wronginput" bson:"wronginput"`
	VerificationString string `json:"verification" bson:"verification"`
	BlockedUser        bool   `json:"blockeduser" bson:"blockeduser"`
}

// Address for customer when ordres
type Address struct {
	FirstName   string `json:"firstname" bson:"firstname"`
	LastName    string `json:"lastname" bson:"lastname"`
	Phone_No    int    `json:"phonenumber" bson:"phonenumber"`
	House_No    string `json:"houseno" bson:"houseno"`
	Street_Name string `json:"streetname" bson:"streetname"`
	City        string `json:"city" bson:"city"`
	Pincode     int64  `json:"pincode" bson:"pincode"`
}

//Customer Login
type Login struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

// Search For Products
type Search struct {
	ProductName string `json:"productName" bson:"productName"`
}

// To add Items To Cart Input
type Addtocart struct { 
	Token       string `json:"token" bson:"token"`
	ProductName string `json:"productName" bson:"productName"`
}

// To Add Items in Cart To DB
type Addcart struct {
	CustomerId   string  `json:"customerid" bson:"customerid"`
	ProductName  string  `json:"productname" bson:"productname"`
	Price        float64 `json:"price" bson:"price"`
	Quantity     int32   `json:"quantity" bson:"quantity"`
	Image        string  `json:"image" bson:"image"`
	SellerID     string  `json:"sellerid" bson:"sellerid"`
	SellerName   string  `json:"sellername" bson:"sellername"`
	ItemCategory string  `json:"itemcategory" bson:"itemcategory"`
}

//Delete Items From Cart
type DeleteProduct struct {
	Token    string `json:"token" bson:"token"`
	Name     string `json:"name" bson:"name"`
	Quantity int    `json:"quantity" bson:"quantity"`
}

//Send to Buyed Items 
type BuyNow struct {
	Token       string  `json:"token"`
	TotalAmount float64 `json:"totalAmount"`
	ItemsToBuy  []Item  `json:"itemsToBuy"`
}

// Name of Quantity of previous
type Item struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

// To Send Total Amount
type TotalAmount struct {
	TotalAmount float64 `json:"totalamount" bson:"totalamount"`
} 

// Delete Order
type DeleteOrder struct {
	Id string `json:"_id" bson:"_id"`
}

//Email Verification
type VerifyEmail struct {
	Email              string `json:"email" bson:"email"`
	VerificationString string `json:"verification" bson:"verification"`
}