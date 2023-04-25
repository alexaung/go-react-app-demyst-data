//backend/main.go

package main

import (
	"log"

	"github.com/alexaung/go-react-app/backend/api/accounting"
	"github.com/alexaung/go-react-app/backend/api/decisionEngine"
	"github.com/alexaung/go-react-app/backend/controllers"
	"github.com/alexaung/go-react-app/backend/middleware"
	"github.com/alexaung/go-react-app/backend/routes"
	"github.com/alexaung/go-react-app/backend/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create instances of services
	accountingProviders := map[string]accounting.AccountingProvider{
		"MYOB": accounting.NewMYOB(),
		"Xero": accounting.NewXero(),
	}

	accountingServices := make(map[string]*services.AccountingService)
	for k, v := range accountingProviders {
		accountingServices[k] = services.NewAccountingService(v)
	}

	decisionEngine := decisionEngine.NewDecisionEngine()

	// Inject services into the controller
	loanController := controllers.NewLoanController(accountingServices, decisionEngine)

	// Set up the Gin router
	router := gin.Default()

	// Use CORS middleware
	router.Use(middleware.CorsMiddleware())

	// Define the routes
	routes.DefineRoutes(router, loanController)

	// Start the server
	port := ":8080"
	log.Printf("Server listening on port %s", port)

	if err := router.Run(port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
