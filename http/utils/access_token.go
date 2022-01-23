package utils

import (
	"errors"
	"majoo-backend/http/enum"
	"majoo-backend/models/requests"
	"majoo-backend/repositories/access_token"
	"net/http"
	"strings"

	"github.com/gocraft/dbr"
)

func CheckAuthorization(r *http.Request, sess *dbr.Session) error {
	authorization := r.Header.Get("Authorization")
	accessToken := strings.Split(authorization, "Bearer ")
	if len(accessToken) != 2 {
		return errors.New(enum.UnauthorizedUser)
	}
	currentToken, err := access_token.GetUserAccessToken(sess, &accessToken[1])
	if currentToken == nil || err != nil {
		return errors.New(enum.UnauthorizedUser)
	}

	return nil
}

func InsertAccessToken(sess *dbr.Session, userName string, accessToken string) error {
	insertAccessTokenRequest := requests.InsertAccessTokenRequest{
		UserName: userName,
		Token:    &accessToken,
	}

	err := access_token.InsertAccessToken(sess, insertAccessTokenRequest)
	if err != nil {
		return err
	}

	return nil
}
