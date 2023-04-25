//backend/models/models/loan.go

package models

type Loan struct {
	ID               int64                  `json:"id"`
	BusinessDetails  map[string]interface{} `json:"businessDetails"`
	LoanAmount       int                `json:"loanAmount"`
	AccountingProvider   string             `json:"accountingProvider"`
	RequestBalanceSheet  bool               `json:"requestBalanceSheet"`
	Status           string                `json:"status"`
	PreAssessment    int                   `json:"preAssessment"`
}

type BalanceSheet struct {
	Year        int     `json:"year"`
	Month       int     `json:"month"`
	ProfitOrLoss int `json:"profitOrLoss"`
	AssetsValue  int `json:"assetsValue"`
}

// You can add additional methods for the Loan model as needed
