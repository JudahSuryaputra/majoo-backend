package access_token

import (
	"majoo-backend/models/db"

	"github.com/gocraft/dbr"
)

func GetUserAccessToken(sess *dbr.Session, accessToken *string) (*db.AccessToken, error) {
	var userAccessToken db.AccessToken

	query := sess.Select("*").
		From(db.AccessToken{}.TableName()).
		Where("token = ?", accessToken)

	err := query.LoadOne(&userAccessToken)
	if err != nil {
		if err != dbr.ErrNotFound {
			return &userAccessToken, err
		}
		return nil, err
	}

	return &userAccessToken, nil
}
