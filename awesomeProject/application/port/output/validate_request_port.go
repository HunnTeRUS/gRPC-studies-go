package output

import (
	"awesomeProject/application/domain"
	"context"
)

type ValidateRequest interface {
	ValidateRequestPort(
		context.Context,
		domain.IntegrateAuthorizationDomain) (map[string]*domain.ModuleValidation, error)
}
