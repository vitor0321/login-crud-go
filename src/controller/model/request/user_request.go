package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%^&*()_+"`
	Name     string `json:"name" binding:"required,min=4,max=70"`
	Age      int16  `json:"age" binding:"required,min=1,max=140"`
}
