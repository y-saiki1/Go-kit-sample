package auth

type RegisterRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
}