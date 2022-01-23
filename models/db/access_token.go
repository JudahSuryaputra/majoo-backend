package db

import "time"

type AccessToken struct {
	ID        int       `db:"id" json:"id"`
	UserName  string    `db:"user_name" json:"user_name"`
	Token     *string   `db:"token" json:"token,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"-"`
}

func (c AccessToken) TableName() string {
	return "access_tokens"
}
