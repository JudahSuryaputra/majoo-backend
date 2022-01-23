package db

import "time"

type Outlet struct {
	ID         int       `db:"id" json:"id"`
	MerchantID string    `db:"merchant_id" json:"merchant_id"`
	OutletName string    `db:"outlet_name" json:"outlet_name"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	CreatedBy  int       `db:"created_by" json:"created_by"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy  int       `db:"updated_by" json:"updated_by"`
}

func (c Outlet) TableName() string {
	return "outlets"
}
