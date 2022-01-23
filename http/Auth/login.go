package Auth

import (
	"encoding/json"
	"errors"
	"majoo-backend/http/enum"
	"majoo-backend/http/formatter"
	"majoo-backend/http/utils"
	"majoo-backend/models/requests"
	"majoo-backend/repositories/user"
	"net/http"

	"github.com/gocraft/dbr"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	DBConn *dbr.Connection
}

func (c Login) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request requests.LoginRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&request)
	if err != nil {
		formatter.ERROR(w, http.StatusBadRequest, err)
		return
	}

	sess := c.DBConn.NewSession(nil)

	currentUser, err := user.GetUserByUserName(sess, request.UserName)
	if err != nil {
		formatter.ERROR(w, http.StatusBadRequest, errors.New(enum.PasswordIncorrect))
		return
	}

	if currentUser == nil {
		formatter.ERROR(w, http.StatusBadRequest, errors.New(enum.PasswordIncorrect))
		return
	}

	err = CheckPasswordHash(currentUser.Password, request.Password)
	if err != nil {
		formatter.ERROR(w, http.StatusBadRequest, errors.New(enum.PasswordIncorrect))
		return
	}

	accessToken, err := utils.EncodeAuthToken(currentUser.ID, currentUser.UserName)
	if err != nil {
		formatter.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = utils.InsertAccessToken(sess, currentUser.UserName, accessToken)
	if err != nil {
		formatter.ERROR(w, http.StatusBadRequest, err)
		return
	}

	formatter.JSON(w, http.StatusOK, struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: accessToken,
	})

	return
}

func CheckPasswordHash(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errors.New(enum.PasswordIncorrect)
	}
	return nil
}
