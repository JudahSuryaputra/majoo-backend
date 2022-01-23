package transaction

import (
	"fmt"
	"majoo-backend/http/enum"
	"majoo-backend/models/db"
	"majoo-backend/models/requests"
	"majoo-backend/models/responses"

	"github.com/gocraft/dbr"
)

func GetMerchantTransactions(sess *dbr.Session, r requests.GetTransactionRequest, merchantID int) ([]responses.MonthlyReport, error) {
	offset := (r.Page - 1) * enum.TransactionPagination
	monthlyReport := []responses.MonthlyReport{}

	columns := []string{
		"DATE(t.created_at) as date",
		"SUM(t.bill_total) as omzet",
	}

	query := sess.Select(columns...).
		From(dbr.I(db.Transaction{}.TableName()).As("t")).
		Where("merchant_id = ? AND created_at BETWEEN ? AND ?", merchantID, r.From, r.To).
		GroupBy("date").
		OrderBy("date")

	query = query.Limit(enum.TransactionPagination).Offset(uint64(offset))

	_, err := query.Load(&monthlyReport)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return monthlyReport, nil
}

func GetOutletTransactions(sess *dbr.Session, r requests.GetTransactionRequest, outletID int) ([]responses.MonthlyReport, error) {
	offset := (r.Page - 1) * enum.TransactionPagination
	monthlyReport := []responses.MonthlyReport{}

	columns := []string{
		"DATE(t.created_at) as date",
		"SUM(t.bill_total) as omzet",
	}

	query := sess.Select(columns...).
		From(dbr.I(db.Transaction{}.TableName()).As("t")).
		Where("outlet_id = ? AND created_at BETWEEN ? AND ?", outletID, r.From, r.To).
		GroupBy("date").
		OrderBy("date")

	query = query.Limit(enum.TransactionPagination).Offset(uint64(offset))

	_, err := query.Load(&monthlyReport)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return monthlyReport, nil
}
