package request

type UserRequest struct {
	FirstName string `json:"first_name" binding:"required,min=2,max=80"`
	LastName  string `json:"last_name" binding:"required,min=2,max=80"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8,containsany=!@#$%^&*()_+"`
	Role      string `json:"role,omitempty" binding:"omitempty,oneof=ADMIN USER"`
	IsActive  bool   `json:"is_active,omitempty" binding:"omitempty,boolean"`
}
