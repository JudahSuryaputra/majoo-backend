package user

import (
	"majoo-backend/models/db"

	"github.com/gocraft/dbr"
)

func GetUserByUserName(sess *dbr.Session, userName string) (*db.User, error) {
	var user db.User

	query := sess.Select("*").
		From(db.User{}.TableName()).
		Where("user_name = ?", userName)

	err := query.LoadOne(&user)
	if err != nil {
		return nil, err
	}
	
	return &user, nil
}
