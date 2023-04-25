//backend/api/accounting/accounting.go

package accounting

import "github.com/alexaung/go-react-app/backend/models"

// AccountingProvider is an interface for different accounting software providers.
type AccountingProvider interface {
    GetBalanceSheet() ([]models.BalanceSheet, error)
}
