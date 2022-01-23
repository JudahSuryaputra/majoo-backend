package User

import (
	"errors"
	"majoo-backend/http/enum"
	"majoo-backend/http/formatter"
	"majoo-backend/http/utils"
	"majoo-backend/models/requests"
	"majoo-backend/repositories/merchant"
	"majoo-backend/repositories/transaction"
	"majoo-backend/repositories/user"
	"net/http"
	"strconv"

	"github.com/gocraft/dbr"
)

type GetMerchantTransactions struct {
	DBConn *dbr.Connection
}

func (c GetMerchantTransactions) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userName := r.Context().Value("UserName").(string)

	sess := c.DBConn.NewSession(nil)

	err := utils.CheckAuthorization(r, sess)
	if err != nil {
		formatter.ERROR(w, http.StatusForbidden, errors.New(enum.UnauthorizedUser))
		return
	}

	currentUser, err := user.GetUserByUserName(sess, userName)
	if err != nil {
		formatter.ERROR(w, http.StatusBadRequest, errors.New(enum.PasswordIncorrect))
		return
	}

	currentMerchant, err := merchant.GetMerchantByUserID(sess, currentUser.ID)
	if err != nil {
		formatter.ERROR(w, http.StatusBadRequest, errors.New("merchant not found"))
		return
	}

	var request requests.GetTransactionRequest
	page := 1
	if r.FormValue("from") != "" {
		from := r.FormValue("from")
		request.From = &from
	}

	if r.FormValue("to") != "" {
		to := r.FormValue("to")
		request.To = &to
	}

	if r.FormValue("page") != "" {
		page, _ = strconv.Atoi(r.FormValue("page"))
		request.Page = page
	}

	currentTransactions, err := transaction.GetMerchantTransactions(sess, request, currentMerchant.ID)
	if err != nil {
		formatter.ERROR(w, http.StatusBadRequest, errors.New("transaction not found"))
		return
	}

	formatter.JSON(w, http.StatusOK, currentTransactions)
	return
}
