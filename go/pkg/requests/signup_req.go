package requests

type SignUp struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
