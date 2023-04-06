package services

import (
	"awesomeProject/application/domain"
	"awesomeProject/application/port/input"
	"awesomeProject/application/port/output"
	"context"
	"fmt"
)

type validateRequest struct {
	validateRequestPort output.ValidateRequest
}

func NewValidateRequestServices(
	validateRequestPort output.ValidateRequest) input.ValidateRequestUseCase {
	return &validateRequest{validateRequestPort: validateRequestPort}
}

func (vl *validateRequest) ValidateRequestServices(
	ctx context.Context,
	request domain.AuthorizationRequest) (*domain.AuthorizationResponse, error) {

	fmt.Println("Init validation services for transaction")
	authorizationDomain := domain.IntegrateAuthorizationDomain{
		AuthorizationRequest:  request,
		EnrichedData:          domain.EnrichedData{},
		AuthorizationResponse: domain.AuthorizationResponse{},
	}

	/* TODO: Call enrichment services to get more data */
	authorizationDomain.EnrichedData = domain.EnrichedData{
		PersonId: "",
		CbCardId: "",
	}

	err := authorizationDomain.ValidateRequest(ctx, vl.validateRequestPort)
	if err != nil {
		return nil, err
	}

	heavierDecisionModuleName, heavierDecisionWeight := authorizationDomain.TakeHeavierDecision()
	authorizationDomain.TakeDecisionAndAnswer(heavierDecisionModuleName, heavierDecisionWeight, request)

	return nil, nil
}
