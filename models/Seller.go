package models

// Seller Creation
type Seller struct {
	SellerId           string `json:"sellerid" bson:"sellerid"`
	Seller_Name        string `json:"sellername" bson:"sellername"`
	Seller_Email       string `json:"selleremail" bson:"selleremail"`
	Password           string `json:"password" bson:"password"`
	ConfirmPassword    string `json:"confirmpassword" bson:"confirmpassword"`
	Phone_No           int    `json:"phoneno" bson:"phoneno"`
	Address            string `json:"address" bson:"address"`
	Image              string `json:"image" bson:"image"`
	WrongInput         int    `json:"wronginput" bson:"wronginput"`
	BlockedUser        bool   `json:"blockeduser" bson:"blockeduser"`
	IsEmailVerified    bool   `json:"isemailverified" bson:"isemailverified"`
	IsApproved         bool   `json:"isapproved" bson:"isapproved"`
	VerificationString string `json:"verification" bson:"verification"`
}

// Add Items To Inventory By Seller
type Inventory struct {
	ProductId        string   `json:"productid" bson:"productid"`
	SellerName       string   `json:"sellername" bson:"sellername"`
	SellerId         string   `json:"sellerid" bson:"sellerid"`
	ItemCategory     string   `json:"itemcategory" bson:"itemcategory"`
	ItemName         string   `json:"itemname" bson:"itemname"`
	Price            float64  `json:"price" bson:"price"`
	Quantity         string   `json:"quantity" bson:"quantity"`
	Stock_Available  int32    `json:"sellerquantity" bson:"sellerquantity"`
	Image            string   `json:"image" bson:"image"`
	ShortDiscription string   `json:"shortdis" bosn:"shortdis"`
	LongDiscription  string   `json:"longdis" bosn:"longdis"`
	Features         []string `json:"features" bson:"features"`
}

// To Delete Product
type DeleteBySeller struct {
	Token     string `json:"token" bson:"token"`
	ProductID string `json:"productid" bson:"productid"`
}

// Update Products
type UpdateProduct struct {
	Token     string `json:"token" bson:"token"`
	ProductID string `json:"productid" bson:"productid"`
	Attribute   string `json:"attribute" bson:"attribute"`
	New_Value   string  `json:"newvalue" bson:"newvalue"`
}

// Seller DrashBoardNeeded Details
type DrashBoard struct {
	Orders          int64 `json:"orders" bson:"orders"`
	OrdersCompleted int64 `json:"orderscompleted" bson:"orderscompleted"`
	OrdersPending   int64 `json:"ordespending" bson:"orderspending"`
	TotalAmount     int64 `json:"totalamount" bson:"totalamount"`
}

// Update Tracking Details
type OrderTracking struct {
	Feild   string `json:"feild" bson:"feild"`
	Token   string `json:"token" bson:"token"`
	OrderID string `json:"orderid" bson:"orderid"`
}
