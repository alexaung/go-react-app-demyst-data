package tests

import (
	"testing"

	"github.com/alexaung/go-react-app/backend/models"
	"github.com/alexaung/go-react-app/backend/rules"
)

func TestEvaluatePreAssessment(t *testing.T) {
	tests := []struct {
		name         string
		loanAmount   int
		balanceSheet []models.BalanceSheet
		expected     int
	}{
		{
			name:       "Full_approval",
			loanAmount: 100000,
			balanceSheet: []models.BalanceSheet{
				{Year: 2020, Month: 12, ProfitOrLoss: 250000, AssetsValue: 1234},
				{Year: 2020, Month: 11, ProfitOrLoss: 1150, AssetsValue: 5789},
				{Year: 2020, Month: 10, ProfitOrLoss: 2500, AssetsValue: 22345},
				{Year: 2020, Month: 9, ProfitOrLoss: 187000, AssetsValue: 223452},
				{Year: 2020, Month: 8, ProfitOrLoss: 187000, AssetsValue: 223452},
				{Year: 2020, Month: 7, ProfitOrLoss: 187000, AssetsValue: 223452},
				{Year: 2020, Month: 6, ProfitOrLoss: 187000, AssetsValue: 223452},
				{Year: 2020, Month: 5, ProfitOrLoss: 187000, AssetsValue: 223452},
				{Year: 2020, Month: 4, ProfitOrLoss: 187000, AssetsValue: 223452},
				{Year: 2020, Month: 3, ProfitOrLoss: 187000, AssetsValue: 223452},
				{Year: 2020, Month: 2, ProfitOrLoss: 187000, AssetsValue: 223452},
				{Year: 2020, Month: 1, ProfitOrLoss: 187000, AssetsValue: 223452},
			},
			expected: 100,
		},
		{
			name:       "Partial_approval",
			loanAmount: 10000,
			balanceSheet: []models.BalanceSheet{
				{Year: 2020, Month: 12, ProfitOrLoss: 250000, AssetsValue: 1234},
				{Year: 2020, Month: 11, ProfitOrLoss: 1150, AssetsValue: 1000},
				{Year: 2020, Month: 10, ProfitOrLoss: 2500, AssetsValue: 1000},
				{Year: 2020, Month: 9, ProfitOrLoss: 187000, AssetsValue: 1000},
				{Year: 2020, Month: 8, ProfitOrLoss: 187000, AssetsValue: 1000},
				{Year: 2020, Month: 7, ProfitOrLoss: 187000, AssetsValue: 1000},
				{Year: 2020, Month: 6, ProfitOrLoss: 187000, AssetsValue: 1000},
				{Year: 2020, Month: 5, ProfitOrLoss: 187000, AssetsValue: 1000},
				{Year: 2020, Month: 4, ProfitOrLoss: 187000, AssetsValue: 1000},
				{Year: 2020, Month: 3, ProfitOrLoss: 187000, AssetsValue: 1000},
				{Year: 2020, Month: 2, ProfitOrLoss: 187000, AssetsValue: 1000},
				{Year: 2020, Month: 1, ProfitOrLoss: 187000, AssetsValue: 1000},
			},
			expected: 60,
		},
		{
			name:       "Default_approval",
			loanAmount: 10000,
			balanceSheet: []models.BalanceSheet{
				{Year: 2020, Month: 12, ProfitOrLoss: 250000, AssetsValue: 1234},
				{Year: 2020, Month: 11, ProfitOrLoss: -1150, AssetsValue: 5789},
				{Year: 2020, Month: 10, ProfitOrLoss: 2500, AssetsValue: 22345},
				{Year: 2020, Month: 9, ProfitOrLoss: -187000, AssetsValue: 223452},
			},
			expected: 20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			preAssessment, _ := rules.EvaluatePreAssessment(tt.loanAmount, tt.balanceSheet)
			if preAssessment != tt.expected {
				t.Errorf("Expected pre-assessment %d, but got %d", tt.expected, preAssessment)
			}
		})
	}
}
