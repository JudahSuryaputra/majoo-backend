package merchant

import (
	"majoo-backend/models/db"

	"github.com/gocraft/dbr"
)

func GetMerchantByUserID(sess *dbr.Session, userID int) (*db.Merchant, error) {
	var merchant db.Merchant

	query := sess.Select("*").
		From(db.Merchant{}.TableName()).
		Where("user_id = ?", userID)

	err := query.LoadOne(&merchant)
	if err != nil {
		return nil, err
	}

	return &merchant, nil
}
