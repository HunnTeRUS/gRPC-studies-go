package grpc_client

import (
	grpc_client_service "awesomeProject/adapter/output/grpc_client_service"
	"awesomeProject/adapter/output/protos/integrator"
	"context"
)

type antifraudClient struct {
}

func _() grpc_client_service.ValidateRequest {
	return &antifraudClient{}
}

func (ac *antifraudClient) DoAuthorize(
	ctx context.Context,
	req *integrator.IntegratorRequest) (*integrator.IntegratorResponse, error) {
	return nil, nil
}
