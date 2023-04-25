package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexaung/go-react-app/backend/api/accounting"
	"github.com/alexaung/go-react-app/backend/api/decisionEngine"
	"github.com/alexaung/go-react-app/backend/controllers"
	"github.com/alexaung/go-react-app/backend/services"
	"github.com/gin-gonic/gin"
)

func TestHandleLoanApplication(t *testing.T) {
	accountingServices := map[string]*services.AccountingService{
		"Xero": services.NewAccountingService(accounting.NewXero()),
	}

	decisionEngine := decisionEngine.NewDecisionEngine()
	decisionEngineService := services.NewDecisionEngineService(decisionEngine)
	loanController := controllers.NewLoanController(accountingServices, decisionEngineService)

	router := gin.Default()
	router.POST("/loan", loanController.HandleLoanApplication)

	loanApplication := map[string]interface{}{
		"loanAmount":          1000,
		"accountingProvider":  "Xero",
		"requestBalanceSheet": true,
		"businessDetails": map[string]interface{}{
			"name":    "Test Business",
			"address": "123 Test St",
		},
	}

	jsonData, _ := json.Marshal(loanApplication)

	req, _ := http.NewRequest("POST", "/loan", bytes.NewBuffer(jsonData))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected HTTP status code 200, got %d", resp.Code)
	}

	// You can add more tests here to verify the content of the response.
}
