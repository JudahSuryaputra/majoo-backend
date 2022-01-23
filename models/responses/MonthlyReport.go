package responses

type MonthlyReport struct {
	Date  string `db:"date" json:"date"`
	Omzet int    `db:"omzet" json:"omzet"`
}
