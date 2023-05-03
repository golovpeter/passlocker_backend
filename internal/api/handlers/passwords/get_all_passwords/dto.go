package get_all_passwords

type PasswordsOut struct {
	ID          int    `db:"id" json:"id"`
	ServiceName string `db:"service_name" json:"serviceName"`
	Link        string `db:"link" json:"link"`
	Email       string `db:"email" json:"email"`
	Login       string `db:"login" json:"login"`
	Password    string `db:"password" json:"password"`
}
