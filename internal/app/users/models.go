package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name" gorm:"text;not null;default:null"`
}

/* type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// FirstName string `json:"first_name"`
	// LastName  string `json:"last_name"`
	Password string `json:"password"`
	// Email        string    `json:"email"`
	// Phone        string    `json:"phone"`
	// Token        string    `json:"-"`
	// RefreshToken string    `json:"-"`
	// CreatedAt    time.Time `json:"-"`
	// UpdatedAt    time.Time `json:"-"`
	// UserID       string    `json:"-"`
	// Address      Address   `json:"address"`
}
*/
/* type User struct {
	ID
	First_Name
	Last_Name
	Password
	Email
	Phone
	Token
	Refresh_Token
	Created_At
	Updated_At
	User_ID
	// UserCart        []products.ProductUser
	Address_Details []Address
	Order_Status
} */

type Address struct {
	ID        string `json:"-"`
	HouseName string `json:"housename"`
	City      string `json:"city"`
	Country   string `json:"country"`
}

// type Address struct {
// Address_id primitive.ObjectID `bson:"_id"`
// House      *string            `json:"house_name" bson:"house_name"`
// Street     *string            `json:"street_name" bson:"street_name"`
// City       *string            `json:"city_name" bson:"city_name"`
// Pincode    *string            `json:"pincode" bson:"pincode"`
// }
