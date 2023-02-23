package login

type In struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Out struct {
	AccessToken string `json:"access_token"`
}
