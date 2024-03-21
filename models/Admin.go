package models

// To Delete Data
type Delete struct {
	Collection string `json:"collection" bson:"collection"`
	IdValue    string `json:"idValue" bson:"idValue"`
}

// To Upadte Feild
type Update struct {
	Collection string `json:"collection" bson:"collection"`
	IdName     string `json:"email" bson:"email"`
	Field      string `json:"field" bson:"field"`
	New_Value  string `json:"newvalue" bson:"newvalue"`
}

// Admin Signup Data
type AdminData struct {
	AdminID    string `json:"adminid" bson:"adminid"`
	Email      string `json:"email" bson:"email"`
	Password   string `json:"password" bson:"password"`
	TOTP       string `json:"totp" bson:"totp"`
	IP_Address string `json:"ip" bson:"ip"`
	SecretKey  string `json:"secretkey" bson:"secretkey"`
	WrongInput int    `json:"wronginput" bson:"wronginput"`
	Token      string `json:"token" bson:"token"`
}

// Admin Sign in data
type AdminSignup struct {
	AdminName       string `json:"name" bson:"name"`
	Email           string `bson:"email" json:"email"`
	Password        string `bson:"password" json:"password"`
	ConfirmPassword string `bson:"confirmpassword" json:"confirmpassword"`
	IP              string `bson:"ip" json:"ip"`
}

// Data Needed for Admin Page
type AdminPageData struct {
	UserCount        int64 `json:"usercount" bson:"usercount"`
	SellerCount      int64 `json:"sellercount" bson:"sellercount"`
	ProductCount     int64 `json:"productcount" bson:"productount"`
	SalesCount       int64 `json:"salescount" bson:"salescount"`
	TotalSalesAmount int32 `json:"totalsalesamount" bson:"totalsalesamount"`
}

// Data Needed for Admin Page -- > Need To Combine Both
type Sales struct {
	TotalSalesAmount int `bson:"totalsalesamount"`
	TotalNoOfSales   int `bson:"totalnoofsales"`
}

// Create Worker
type Workers struct {
	UserName string `bson:"username" json:"username"`
	Email    string `bson:"email" json:"email"`
	Role     string `bson:"role" json:"role"`
	No       string `bson:"no" json:"no"`
	Salary   int64  `bson:"salary" json:"salary"`
	Status   string `bson:"status" json:"status"`
	Image    string `bson:"image" json:"image"`
}

// Get Every Single Data
type Getdata struct {
	Id         string `json:"id" bson:"id"`
	Collection string `json:"collection" bson:"collection"`
}

// Single Data Returing Structure
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
	CustomerId         string `json:"customerid" bson:"customerid"`
	Name               string `json:"name" bson:"name"`
	IsEmailVerified    bool   `json:"isemailverified" bson:"isemailverified"`
	WrongInput         int    `json:"wronginput" bson:"wronginput"`
	VerificationString string `json:"verification" bson:"verification"`
	BlockedUser        bool   `json:"blockeduser" bson:"blockeduser"`
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

// Upload Event to Calender
type UploadCalender struct {
	AdminEmail string   `json:"email" bson:"email"`
	Title      string   `json:"title" bson:"title"`
	Start      string   `json:"start" bson:"start"`
	End        string   `json:"end" bson:"end"`
	Todos      []string `json:"todos" bson:"todos"`
}

// Input to Get Email
type GetCalender struct {
	AdminEmail string `json:"email" bson:"email"`
}

// Block User
type Block struct {
	Email      string `json:"email" bson:"email"`
	Collection string `json:"collection" bson:"collection"`
}

// ShutDown
type ShutDown struct {
	Token    string `json:"token" bson:"token"`
	Password string `json:"password" bson:"password"`
}
