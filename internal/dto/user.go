package dto

type UserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthOutput struct {
	AccessToken string `json:"access_token"`
}
