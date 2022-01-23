package db

import "time"

type Transaction struct {
	ID         int       `db:"id" json:"id"`
	MerchantID int       `db:"merchant_id" json:"merchant_id"`
	OutletID   int       `db:"outlet_id" json:"outlet_id"`
	BillTotal  float64   `db:"bill_total" json:"bill_total"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	CreatedBy  int       `db:"created_by" json:"created_by"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy  int       `db:"updated_by" json:"updated_by"`
}

func (c Transaction) TableName() string {
	return "transactions"
}
