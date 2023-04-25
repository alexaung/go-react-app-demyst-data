//backend/rules/loanRules.go

package rules

import "github.com/alexaung/go-react-app/backend/models"

const (
	DefaultPreAssessment         = 20
	PartialApprovalPreAssessment = 60
	FullApprovalPreAssessment    = 100
)

func EvaluatePreAssessment(loanAmount int, balanceSheet []models.BalanceSheet) (int, map[int]int) {
	var profitMonths, totalAssetValue int
	profitSummary := make(map[int]int)

	for _, sheet := range balanceSheet {
		if sheet.ProfitOrLoss > 0 {
			profitMonths++
		}
		totalAssetValue += sheet.AssetsValue
		profitSummary[sheet.Year] += sheet.ProfitOrLoss
	}

	averageAssetValue := totalAssetValue / int(len(balanceSheet))

	preAssessment := DefaultPreAssessment

	if len(balanceSheet) >= 12 && profitMonths == len(balanceSheet) { // Check if the balance sheet has at least 12 months of data and all of them have profit
		if averageAssetValue > loanAmount { // Check if the average asset value across 12 months is greater than the loan amount
			preAssessment = FullApprovalPreAssessment // preAssessment: 100
		} else {
			preAssessment = PartialApprovalPreAssessment // preAssessment: 60
		}
	}

	return preAssessment, profitSummary // preAssessment: 20
}
