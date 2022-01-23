package User

import (
	"errors"
	"majoo-backend/http/enum"
	"majoo-backend/http/formatter"
	"majoo-backend/http/utils"
	"majoo-backend/models/requests"
	"majoo-backend/repositories/merchant"
	"majoo-backend/repositories/outlet"
	"majoo-backend/repositories/transaction"
	"majoo-backend/repositories/user"
	"net/http"
	"strconv"

	"github.com/gocraft/dbr"
	"github.com/gorilla/mux"
)

type GetOutletTransactions struct {
	DBConn *dbr.Connection
}

func (c GetOutletTransactions) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userName := r.Context().Value("UserName").(string)
	pathVariable := mux.Vars(r)
	outletID, _ := strconv.Atoi(pathVariable["id"])

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
		formatter.ERROR(w, http.StatusBadRequest, errors.New(enum.UnauthorizedUser))
		return
	}

	currentOutlet, err := outlet.GetOutletByID(sess, currentMerchant.ID, outletID)
	if err != nil {
		formatter.ERROR(w, http.StatusForbidden, errors.New(enum.UnauthorizedUser))
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

	currentTransactions, err := transaction.GetOutletTransactions(sess, request, currentOutlet.ID)
	if err != nil {
		formatter.ERROR(w, http.StatusBadRequest, errors.New("transaction not found"))
		return
	}

	formatter.JSON(w, http.StatusOK, currentTransactions)
	return
}
