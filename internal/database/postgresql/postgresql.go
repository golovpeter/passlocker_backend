package postgresql

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Conn *sqlx.DB
}

func NewDatabase(dbUrl string) *Config {
	return &Config{
		Conn: OpenConnection(dbUrl),
	}
}

func (c *Config) ExistUser(result *bool, email string) error {
	return c.Conn.Get(result, "select exists(select email from users where email = $1)", email)
}

func (c *Config) InsertUser(email string, passwordHash string) (sql.Result, error) {
	return c.Conn.Exec("insert into users (email, password_hash) values ($1, $2)", email, passwordHash)
}

func (c *Config) SelectUserData(userData any, email string) error {
	return c.Conn.Get(userData, "select user_id, email, password_hash from users where email = $1", email)
}

func (c *Config) DeleteUser(accessToken string) (sql.Result, error) {
	return c.Conn.Exec("delete from tokens where access_token = $1", accessToken)
}

func (c *Config) InsertTokens(userID int, newDeviceID uuid.UUID, accessToken string, refreshToken string) (sql.Result, error) {
	return c.Conn.Exec("insert into tokens values ($1, $2, $3, $4)", userID, newDeviceID, accessToken, refreshToken)
}

func (c *Config) ExistRefreshToken(result *bool, deviceID string, refreshToken string) error {
	return c.Conn.Get(result,
		"select exists(select refresh_token, device_id from tokens where device_id = $1 and refresh_token=$2)",
		deviceID,
		refreshToken,
	)
}

func (c *Config) ExistAccessToken(result *bool, deviceID string, accessToken string) error {
	return c.Conn.Get(
		result,
		"select exists(select device_id, access_token from tokens where device_id = $1 and access_token = $2)",
		deviceID,
		accessToken,
	)
}

func (c *Config) UpdateTokens(newAccessToken string, newRefreshToken string, deviceID string) (sql.Result, error) {
	return c.Conn.Exec(
		"update tokens set access_token=$1, refresh_token=$2 where device_id = $3",
		newAccessToken,
		newRefreshToken,
		deviceID,
	)
}

func (c *Config) InsertPassword(userID int, serviceName string, link string, email string, login string, password string) (sql.Result, error) {
	return c.Conn.Exec(
		"insert into passwords (user_id, service_name, link, email, login, password) values ($1, $2, $3, $4, $5, $6)",
		userID,
		serviceName,
		link,
		email,
		login,
		password,
	)
}

func (c *Config) SelectMaxPasswordID(passwordID *int) error {
	return c.Conn.Get(passwordID, "select max(id) from passwords")
}

func (c *Config) SelectPasswordUserID(result *int, passwordID int) error {
	return c.Conn.Get(result, "select user_id from passwords where id = $1", passwordID)
}

func (c *Config) DeletePassword(passwordId int) (sql.Result, error) {
	return c.Conn.Exec("delete from passwords where id = $1", passwordId)
}

func (c *Config) SelectAllPasswords(result any, userID int) error {
	return c.Conn.Select(
		result,
		"select id, service_name, link, email, login, password from passwords where user_id = $1",
		userID,
	)
}
