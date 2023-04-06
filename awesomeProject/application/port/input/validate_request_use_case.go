package input

import (
	"awesomeProject/application/domain"
	"context"
)

type ValidateRequestUseCase interface {
	ValidateRequestServices(
		context.Context,
		domain.AuthorizationRequest) (*domain.AuthorizationResponse, error)
}
