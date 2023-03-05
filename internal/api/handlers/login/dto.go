package login

type loginIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginOut struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
