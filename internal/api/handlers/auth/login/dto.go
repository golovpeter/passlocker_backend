package login

type loginIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginOut struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type User struct {
	UserID       int    `db:"user_id"`
	Email        string `db:"email"`
	HashPassword string `db:"password_hash"`
}
