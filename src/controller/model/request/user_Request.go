package request

type UserRequest struct {
	Name     string `json:"name" biding:"required, min = 3,max = 100"`
	Email    string `json:"email"biding:"required, email"`
	Password string `json:"password"biding:"required, min =6,containany= !@#$%&*"`
	age      int8   `json:"age"biding:"required,min=4, max = 90`
}
