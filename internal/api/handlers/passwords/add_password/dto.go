package add_password

type addPasswordIn struct {
	ServiceName string `json:"service_name"`
	Link        string `json:"link"`
	Login       string `json:"login"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}
