package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type Database interface {
	Users
	Tokens
	Passwords
}

type Users interface {
	ExistUser(result *bool, email string) error
	InsertUser(email string, passwordHash string) (sql.Result, error)
	SelectUserData(userData any, email string) error
	DeleteUser(accessToken string) (sql.Result, error)
}

type Tokens interface {
	InsertTokens(userID int, newDeviceID uuid.UUID, accessToken string, refreshToken string) (sql.Result, error)
	ExistRefreshToken(result *bool, deviceID string, refreshToken string) error
	ExistAccessToken(result *bool, deviceID string, accessToken string) error
	UpdateTokens(newAccessToken string, newRefreshToken string, deviceID string) (sql.Result, error)
}

type Passwords interface {
	InsertPassword(userID int, serviceName string, link string, email string, login string, password string) (sql.Result, error)
	SelectMaxPasswordID(passwordID *int) error
	SelectPasswordUserID(result *int, passwordID int) error
	DeletePassword(passwordId int) (sql.Result, error)
	SelectAllPasswords(result any, userID int) error
}
