package output

import (
	"awesomeProject/adapter/output/protos/integrator"
	"context"
)

type ValidateRequest interface {
	ValidateRequestPort(
		context.Context,
		*integrator.IntegratorRequest) (*integrator.IntegratorResponse, error)
}
