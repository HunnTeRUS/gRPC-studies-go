package grpc_client_service

import (
	"awesomeProject/adapter/output/protos/integrator"
	"awesomeProject/application/port/output"
	"awesomeProject/configuration/configmap"
	"context"
	"fmt"
)

type PRODUCT_NAME string

type ValidateRequest interface {
	DoAuthorize(ctx context.Context,
		req *integrator.IntegratorRequest) (*integrator.IntegratorResponse, error)
}

type validateRequestClient struct {
	mapAuthorization map[string]configmap.ModuleConfiguration
}

func NewValidateRequestClientService(
	mapAuthorization map[string]configmap.ModuleConfiguration) output.ValidateRequest {
	return &validateRequestClient{mapAuthorization}
}

func (v *validateRequestClient) ValidateRequestPort(
	ctx context.Context,
	req *integrator.IntegratorRequest) (*integrator.IntegratorResponse, error) {

	for module, configuration := range v.mapAuthorization {
		fmt.Printf("Getting configuration for %s module \n", module)

		if configuration.IsEnable && isProgramAllowed(req.GetProductId(), configuration) {
			fmt.Printf("Request is allowed for %s module with product %d\n", module, req.GetProductId())
			_, err := configuration.AuthInterface.DoAuthorize(ctx, req)
			if err != nil {
				return nil, err
			}
		}

	}
	return nil, nil
}

func isProgramAllowed(productId int64, configuration configmap.ModuleConfiguration) bool {
	for _, product := range configuration.AllowedProgramIds {
		if product == productId {
			return true
		}
	}

	return false
}
