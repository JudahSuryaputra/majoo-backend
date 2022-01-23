package access_token

import (
	"majoo-backend/models/db"
	"majoo-backend/models/requests"

	"github.com/gocraft/dbr"
)

func InsertAccessToken(sess *dbr.Session, r requests.InsertAccessTokenRequest) error {
	accessToken := db.AccessToken{
		UserName: r.UserName,
		Token:    r.Token,
	}

	columns := []string{
		"user_name",
		"token",
	}

	_, err := sess.InsertInto(db.AccessToken{}.TableName()).
		Columns(columns...).
		Record(accessToken).
		Exec()
	if err != nil {
		return err
	}

	return nil
}
