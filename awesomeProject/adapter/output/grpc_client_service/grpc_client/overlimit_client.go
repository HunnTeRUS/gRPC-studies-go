package grpc_client

import (
	grpc_client_service "awesomeProject/adapter/output/grpc_client_service"
	"awesomeProject/adapter/output/protos/integrator"
	"context"
)

type overlimit_client struct {
}

func NewOverlimitClient() grpc_client_service.ValidateRequest {
	return &overlimit_client{}
}

func (ac *overlimit_client) DoAuthorize(
	ctx context.Context,
	req *integrator.IntegratorRequest) (*integrator.IntegratorResponse, error) {
	return nil, nil
}
