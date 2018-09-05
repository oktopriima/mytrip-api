package auth

type gettokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (this gettokenRequest) GetEmail() string {
	return this.Email
}
func (this gettokenRequest) GetPassword() string {
	return this.Password
}
