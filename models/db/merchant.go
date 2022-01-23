package db

import "time"

type Merchant struct {
	ID           int       `db:"id" json:"id"`
	UserID       int       `db:"user_id" json:"user_id"`
	MerchantName string    `db:"merchant_name" json:"merchant_name"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	CreatedBy    int       `db:"created_by" json:"created_by"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy    int       `db:"updated_by" json:"updated_by"`
}

func (c Merchant) TableName() string {
	return "merchants"
}
