//backend/routes/routes.go

package routes

import (
    "github.com/alexaung/go-react-app/backend/controllers"
    "github.com/gin-gonic/gin"
)

// Define the routes and group them
func DefineRoutes(router *gin.Engine, loanController *controllers.LoanController) {
    api := router.Group("/api")
    {
        v1 := api.Group("/v1")
        {
            v1.POST("/loan", loanController.HandleLoanApplication)
            v1.GET("/balance-sheet/:accountingProvider", loanController.FetchBalanceSheet)
        }
    }
}
