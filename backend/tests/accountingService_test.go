package tests

import (
	"testing"

	"github.com/alexaung/go-react-app/backend/api/accounting"
	"github.com/alexaung/go-react-app/backend/services"
)

func TestGetBalanceSheet(t *testing.T) {
	accountingProvider := "Xero" // You can change this to the desired provider

	var provider accounting.AccountingProvider
	switch accountingProvider {
	case "Xero":
		provider = accounting.NewXero()
	case "MYOB":
		// Instantiate MYOB provider here, if needed
		t.Skip("MYOB provider not implemented") // Skip the test if MYOB is not implemented
	default:
		t.Fatalf("Unsupported accounting provider: %s", accountingProvider)
	}

	accountingService := services.NewAccountingService(provider)

	balanceSheet, err := accountingService.GetBalanceSheet()

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if len(balanceSheet) != 4 {
		t.Errorf("Expected balance sheet length of 4, got: %d", len(balanceSheet))
	}

	// Additional tests can be added here to verify the content of the balanceSheet.
	for _, sheet := range balanceSheet {
		if sheet.Year == 0 || sheet.Month == 0 {
			t.Errorf("Year and Month should not be 0")
		}
	}
}
