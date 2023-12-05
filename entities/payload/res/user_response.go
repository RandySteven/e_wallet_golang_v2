package res

type UserResponse struct {
	ID    uint
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserLoginResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}
