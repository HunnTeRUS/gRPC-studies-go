package grpc_client_service

import (
	"awesomeProject/adapter/output/protos/integrator"
	"awesomeProject/application/domain"
	"awesomeProject/application/port/output"
	"awesomeProject/configuration/configmap"
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/any"
)

type ValidateRequest interface {
	DoAuthorize(ctx context.Context,
		req domain.IntegrateAuthorizationDomain) (*integrator.IntegratorResponse, error)
}

type validateRequestClient struct {
	configurations configmap.ExternalModules
}

func NewValidateRequestClientService(
	configurations configmap.ExternalModules) output.ValidateRequest {
	return &validateRequestClient{configurations}
}

func (v *validateRequestClient) ValidateRequestPort(
	ctx context.Context,
	reqDomain domain.IntegrateAuthorizationDomain) (moduleDomainResponses map[string]*domain.ModuleValidation, err error) {

	for module, configuration := range v.configurations.MapAuthorization {
		fmt.Printf("Getting configuration for %s module \n", module)

		if configuration.IsEnable && isProgramAllowed(reqDomain.AuthorizationRequest.ProgramId, configuration) {
			fmt.Printf("Request is allowed for %s module with product %d\n", module, reqDomain.AuthorizationRequest.ProgramId)
			moduleRes, err := configuration.AuthInterface.DoAuthorize(ctx, reqDomain)
			if err != nil {
				moduleDomainResponses[module] = createModuleDomainResponse(
					module,
					integrator.IntegratorResponse_SKIPPED,
					"Module was skipped because occurs an error when calling it",
					"",
					nil)
				continue
			}

			moduleDomainResponses[module] = createModuleDomainResponse(
				module,
				moduleRes.Status,
				moduleRes.Reason,
				moduleRes.CustomCode,
				moduleRes.Metadata)
			continue
		}

		moduleDomainResponses[module] = createModuleDomainResponse(
			module,
			integrator.IntegratorResponse_SKIPPED,
			fmt.Sprintf(
				"Module was skipped because is disabled or the product %d does not allow it",
				reqDomain.AuthorizationRequest.ProgramId),
			"",
			nil)
	}
	return
}

func createModuleDomainResponse(
	name string,
	status integrator.IntegratorResponse_Status,
	reason string,
	customCode string,
	metadata map[string]*any.Any) *domain.ModuleValidation {
	return &domain.ModuleValidation{
		Name:       name,
		Status:     domain.IntegratorDomain_Status(status),
		Reason:     reason,
		CustomCode: customCode,
		Metadata:   metadata,
	}
}

func isProgramAllowed(productId int64, configuration configmap.ModuleConfiguration) bool {
	for _, product := range configuration.AllowedProgramIds {
		if product == productId {
			return true
		}
	}

	return false
}
