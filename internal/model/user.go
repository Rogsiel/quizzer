package model

import (
	"time"
)

type User struct {
ID                int64     `json:"id"`
UserName          string    `json:"user_name"`
Email             string    `json:"email"`
HashedPassword    string    `json:"hashed_password"`
PasswordChangedAt time.Time `json:"password_changed_at"`
CreatedAt         time.Time `json:"created_at"`
}
