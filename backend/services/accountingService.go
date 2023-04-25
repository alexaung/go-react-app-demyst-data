//backend/services/accountingService.go

package services

import (
	"github.com/alexaung/go-react-app/backend/api/accounting"
	"github.com/alexaung/go-react-app/backend/models"
)

type AccountingService struct {
	Provider accounting.AccountingProvider
}

func NewAccountingService(provider accounting.AccountingProvider) *AccountingService {
	return &AccountingService{Provider: provider}
}

func (s *AccountingService) GetBalanceSheet() ([]models.BalanceSheet, error) {
	balanceSheet, err := s.Provider.GetBalanceSheet()
	if err != nil {
		return nil, err
	}

	return balanceSheet, nil
}
