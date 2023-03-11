package get_all_passwords

type passwordsOut struct {
	ServiceName string `db:"service_name"`
	Link        string `db:"link"`
	Email       string `db:"email"`
	Login       string `db:"login"`
	Password    string `db:"password"`
}
