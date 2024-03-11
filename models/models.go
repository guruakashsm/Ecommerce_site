package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	BlockedUser        bool `json:"blockeduser" bson:"blockeduser"`
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
	Image           string `json:"image" bson:"image"`
	WrongInput         int    `json:"wronginput" bson:"wronginput"`
	BlockedUser        bool `json:"blockeduser" bson:"blockeduser"`
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
	Role     string `json:"role" bson:"role"`
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
type Userdata struct {
	Token    string `json:"token" bson:"token"`
	UserName string `json:"username" bson:"username"`
}

type AdminData struct {
	Id         primitive.ObjectID `json:"_id" bson:"_id"`
	Email      string             `json:"email" bson:"email"`
	Password   string             `json:"password" bson:"password"`
	TOTP       string             `json:"totp" bson:"totp"`
	IP_Address string             `json:"ip" bson:"ip"`
	SecretKey  string             `json:"secretkey" bson:"secretkey"`
	WrongInput int                `json:"wronginput" bson:"wronginput"`
	Token      string             `json:"token" bson:"token"`
}
type Admin struct {
	Email      string `json:"email" bson:"email"`
	Password   string `json:"password" bson:"password"`
	IP_Address string `json:"ip" bson:"ip"`
	SecretKey  string `json:"secretkey" bson:"secretkey"`
	WrongInput int    `json:"wronginput" bson:"wronginput"`
	Token      string `json:"token" bson:"token"`
}

type AdminPageData struct {
	UserCount        int64 `json:"usercount" bson:"usercount"`
	SellerCount      int64 `json:"sellercount" bson:"sellercount"`
	ProductCount     int64 `json:"productcount" bson:"productount"`
	SalesCount       int64 `json:"salescount" bson:"salescount"`
	TotalSalesAmount int32 `json:"totalsalesamount" bson:"totalsalesamount"`
}

type Sales struct {
	TotalSalesAmount int `bson:"totalsalesamount"`
	TotalNoOfSales   int `bson:"totalnoofsales"`
}

type Workers struct {
	UserName string `bson:"username" json:"username"`
	Email    string `bson:"email" json:"email"`
	Role     string `bson:"role" json:"role"`
	No       string `bson:"no" json:"no"`
	Salary   int64  `bson:"salary" json:"salary"`
	Status   string `bson:"status" json:"status"`
	Image    string `bson:"image" json:"image"`
}

type AdminSignup struct {
	AdminName       string `json:"name" bson:"name"`
	Email           string `bson:"email" json:"email"`
	Password        string `bson:"password" json:"password"`
	ConfirmPassword string `bson:"confirmpassword" json:"confirmpassword"`
	IP              string `bson:"ip" json:"ip"`
}

type Getdata struct {
	Id         string `json:"id" bson:"id"`
	Collection string `json:"collection" bson:"collection"`
}

type ReturnData struct {
	// worker
	UserName string `bson:"username" json:"username"`
	Role     string `bson:"role" json:"role"`
	No       string `bson:"no" json:"no"`
	Salary   int64  `bson:"salary" json:"salary"`
	Status   string `bson:"status" json:"status"`
	// inventory

	ItemCategory    string  `json:"itemcategory" bson:"itemcategory"`
	ItemName        string  `json:"itemname" bson:"itemname"`
	Price           float64 `json:"price" bson:"price"`
	Quantity        string  `json:"quantity" bson:"quantity"`
	Stock_Available int32   `json:"sellerquantity" bson:"sellerquantity"`
	// seller

	Seller_Email string `json:"selleremail" bson:"selleremail"`
	//customer
	CustomerId string `json:"customerid" bson:"customerid"`
	Name       string `json:"name" bson:"name"`
	IsEmailVerified    bool   `json:"isemailverified" bson:"isemailverified"`
	WrongInput         int    `json:"wronginput" bson:"wronginput"`
	VerificationString string `json:"verification" bson:"verification"`
	BlockedUser        bool `json:"blockeduser" bson:"blockeduser"`
	//common feilds

	Seller_Name     string `json:"sellername" bson:"sellername"`
	Image           string `json:"image" bson:"image"`
	Email           string `json:"email" bson:"email"`
	SellerId        string `json:"sellerid" bson:"sellerid"`
	Address         string `json:"address" bson:"address"`
	Phone_No        int    `json:"phonenumber" bson:"phonenumber"`
	Password        string `json:"password" bson:"password"`
	ConfirmPassword string `json:"confirmpassword" bson:"confirmpassword"`
}

type UploadCalender struct {
	AdminEmail string   `json:"email" bson:"email"`
	Title      string   `json:"title" bson:"title"`
	Start      string   `json:"start" bson:"start"`
	End        string   `json:"end" bson:"end"`
	Todos      []string `json:"todos" bson:"todos"`
}

type GetCalender struct {
	AdminEmail string `json:"email" bson:"email"`
}

type VerifyEmail struct {
	Email              string `json:"email" bson:"email"`
	VerificationString string `json:"verification" bson:"verification"`
}

type Block struct {
	Email              string `json:"email" bson:"email"`
	Collection string `json:"collection" bson:"collection"`
}
