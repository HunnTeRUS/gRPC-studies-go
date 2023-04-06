package domain

import (
	"awesomeProject/application/port/output"
	"awesomeProject/configuration/configmap"
	"context"
	"github.com/golang/protobuf/ptypes/any"
	"time"
)

type IntegratorDomain_Status int32

const LOWER_DECISION_WEIGHT = -1

const (
	IntegratorDomain_APPROVED IntegratorDomain_Status = 0
	IntegratorDomain_SKIPPED  IntegratorDomain_Status = 1
	IntegratorDomain_DECLINED IntegratorDomain_Status = 2
	IntegratorDomain_REFERRED IntegratorDomain_Status = 3
)

type IntegrateAuthorizationDomain struct {
	AuthorizationRequest  AuthorizationRequest
	EnrichedData          EnrichedData
	AuthorizationResponse AuthorizationResponse
}

type AuthorizationRequest struct {
	AuthorizationId string
	CardId          int64
	ProgramId       int64
	CustomCode      string
	ResponseCode    string
}

type AuthorizationResponse struct {
	Timestamp         string
	Approve           bool
	Referral          bool
	CustomCode        string
	ResponseCode      string
	ForceApprove      bool
	ModuleValidations map[string]*ModuleValidation
	Metadata          map[string]*any.Any
}

type ModuleValidation struct {
	Name         string
	Status       IntegratorDomain_Status
	Reason       string
	ResponseCode string
	CustomCode   string
	Metadata     map[string]*any.Any
}

type EnrichedData struct {
	PersonId string
	CbCardId string
}

func (ia IntegrateAuthorizationDomain) ValidateRequest(
	ctx context.Context,
	validateRequestPort output.ValidateRequest) error {
	mapResponse, err := validateRequestPort.ValidateRequestPort(ctx, ia)
	if err != nil {
		return err
	}

	ia.AuthorizationResponse.ModuleValidations = mapResponse
	return nil
}

func (ia *IntegrateAuthorizationDomain) TakeHeavierDecision() (
	heavierDecisionModuleName string, heavierDecisionWeight int8) {
	heavierDecisionWeight = LOWER_DECISION_WEIGHT

	for moduleName, module := range ia.AuthorizationResponse.ModuleValidations {
		if module.Status != IntegratorDomain_SKIPPED &&
			configmap.Configuration.MapAuthorization[moduleName].DecisionWeight >
				heavierDecisionWeight {
			heavierDecisionModuleName = moduleName
			heavierDecisionWeight = configmap.Configuration.MapAuthorization[moduleName].DecisionWeight
		}
	}

	return
}

func (ia *IntegrateAuthorizationDomain) TakeDecisionAndAnswer(
	heavierDecisionModuleName string,
	heavierDecisionWeight int8,
	request AuthorizationRequest) {
	if heavierDecisionWeight == LOWER_DECISION_WEIGHT {
		ia.takeProcessorDecision(request)
		return
	}

	isApproved := ia.AuthorizationResponse.ModuleValidations[heavierDecisionModuleName].Status == IntegratorDomain_APPROVED
	isReferred := ia.AuthorizationResponse.ModuleValidations[heavierDecisionModuleName].Status == IntegratorDomain_REFERRED

	ia.AuthorizationResponse = AuthorizationResponse{
		Timestamp:         time.Now().String(),
		Approve:           isApproved,
		Referral:          isReferred,
		ForceApprove:      isApproved,
		CustomCode:        ia.getCustomCode(heavierDecisionModuleName),
		ResponseCode:      ia.getResponseCode(heavierDecisionModuleName),
		ModuleValidations: ia.AuthorizationResponse.ModuleValidations,
		Metadata:          nil,
	}
}

func (ia *IntegrateAuthorizationDomain) takeProcessorDecision(request AuthorizationRequest) {
	/*
		TODO: Verify with the team if we can take approve and referral fields without mapping all custom_code and
		response_codes in configmap. Just keep decision without rules.
	*/
	ia.AuthorizationResponse = AuthorizationResponse{
		Timestamp:         time.Now().String(),
		Approve:           false,
		Referral:          false,
		ForceApprove:      false,
		CustomCode:        request.CustomCode,
		ResponseCode:      request.ResponseCode,
		ModuleValidations: ia.AuthorizationResponse.ModuleValidations,
		Metadata:          nil,
	}
}

func (ia *IntegrateAuthorizationDomain) getResponseCode(heavierDecisionModuleName string) string {
	return ia.AuthorizationResponse.ModuleValidations[heavierDecisionModuleName].ResponseCode
}

func (ia *IntegrateAuthorizationDomain) getCustomCode(heavierDecisionModuleName string) string {
	return ia.AuthorizationResponse.ModuleValidations[heavierDecisionModuleName].CustomCode
}
