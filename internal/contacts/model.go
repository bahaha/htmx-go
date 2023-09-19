package contacts

type Contact struct {
	ID        int    `json:"id"         form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name"  form:"last_name"`
	Phone     string `json:"phone"      form:"phone"`
	Email     string `json:"email"      form:"email"`
}
