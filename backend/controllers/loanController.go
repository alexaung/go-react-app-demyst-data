//backend/controllers/loanController.go

package controllers

import (
	"net/http"

	"github.com/alexaung/go-react-app/backend/api/decisionEngine"
	"github.com/alexaung/go-react-app/backend/rules"
	"github.com/alexaung/go-react-app/backend/services"

	"github.com/gin-gonic/gin"
)

type LoanController struct {
	AccountingServices    map[string]*services.AccountingService
	DecisionEngineService *services.DecisionEngineService
}

type LoanApplication struct {
	BusinessDetails    map[string]interface{} `json:"businessDetails"`
	LoanAmount         int                    `json:"loanAmount"`
	AccountingProvider string                 `json:"accountingProvider"`
}

func NewLoanController(accountingServices map[string]*services.AccountingService, decisionEngineService *services.DecisionEngineService) *LoanController {
	return &LoanController{
		AccountingServices:    accountingServices,
		DecisionEngineService: decisionEngineService,
	}
}

func (lc *LoanController) HandleLoanApplication(c *gin.Context) {
	var application LoanApplication
	err := c.ShouldBindJSON(&application)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accountingService, ok := lc.AccountingServices[application.AccountingProvider]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Accounting provider not supported"})
		return
	}

	balanceSheet, err := accountingService.GetBalanceSheet()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	preAssessment, profitSummary := rules.EvaluatePreAssessment(application.LoanAmount, balanceSheet)

	application.BusinessDetails["profitSummary"] = profitSummary

	decisionRequest := decisionEngine.DecisionRequest{
		BusinessDetails: application.BusinessDetails,
		PreAssessment:   preAssessment,
	}

	decisionResponse, err := lc.DecisionEngineService.RequestLoanDecision(decisionRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	approvalStatus := "Denied"
	if preAssessment == rules.PartialApprovalPreAssessment {
		approvalStatus = "Partially Approved"
	} else if preAssessment == rules.FullApprovalPreAssessment {
		approvalStatus = "Fully Approved"
	}

	c.JSON(http.StatusOK, gin.H{
		"decision":       decisionResponse,
		"preAssessment":  preAssessment,
		"approvalStatus": approvalStatus,
	})
}

func (lc *LoanController) FetchBalanceSheet(c *gin.Context) {
	accountingProvider := c.Param("accountingProvider")
	accountingService, ok := lc.AccountingServices[accountingProvider]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Accounting provider not supported"})
		return
	}

	balanceSheet, err := accountingService.GetBalanceSheet()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, balanceSheet)
}
