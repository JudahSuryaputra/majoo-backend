package outlet

import (
	"majoo-backend/models/db"

	"github.com/gocraft/dbr"
)

func GetOutletByID(sess *dbr.Session, merchantID, outletID int) (*db.Outlet, error) {
	var outlet db.Outlet

	query := sess.Select("*").
		From(db.Outlet{}.TableName()).
		Where("id = ? AND merchant_id = ?", outletID, merchantID)

	err := query.LoadOne(&outlet)
	if err != nil {
		return nil, err
	}

	return &outlet, nil
}
