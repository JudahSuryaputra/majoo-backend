package db

import "time"

type User struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	UserName  string    `db:"user_name" json:"user_name"`
	Password  string    `db:"password" json:"-"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	CreatedBy int       `db:"created_by" json:"created_by"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy int       `db:"updated_by" json:"updated_by"`
}

func (u User) TableName() string {
	return "users"
}
