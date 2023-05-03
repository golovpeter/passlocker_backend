package add_password

// TODO: сделать все json поля в snake-кейсе
type addPasswordIn struct {
	ServiceName string `json:"serviceName"`
	Link        string `json:"link"`
	Login       string `json:"login"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}
