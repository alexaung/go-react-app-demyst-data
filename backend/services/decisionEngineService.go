// backend/services/decisionEngineService.go

package services

import (
	"errors"

	"github.com/alexaung/go-react-app/backend/api/decisionEngine"
)

type DecisionEngineService struct {
	decisionEngine *decisionEngine.DecisionEngine
}

func NewDecisionEngineService(decisionEngine *decisionEngine.DecisionEngine) *DecisionEngineService {
	return &DecisionEngineService{
		decisionEngine: decisionEngine,
	}
}

func (des *DecisionEngineService) RequestLoanDecision(request decisionEngine.DecisionRequest) (decisionEngine.DecisionResponse, error) {
	// Validate the input data
	err := validateDecisionRequest(request)
	if err != nil {
		return decisionEngine.DecisionResponse{}, err
	}

	// Request the decision from the decision engine
	decision, err := des.decisionEngine.RequestDecision(request)
	if err != nil {
		return decisionEngine.DecisionResponse{}, err
	}

	return decision, nil
}

func validateDecisionRequest(request decisionEngine.DecisionRequest) error {
	if len(request.BusinessDetails) == 0 {
		return errors.New("business details are required")
	}

	if request.PreAssessment < 0 {
		return errors.New("preAssessment value must be greater than or equal to 0")
	}

	return nil
}
