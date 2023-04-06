package services

import (
	"awesomeProject/adapter/output/protos/integrator"
	"awesomeProject/application/port/output"
	"context"
	"fmt"
)

type validateRequest struct {
	validateRequestPort output.ValidateRequest
}

func NewValidateRequestServices(
	validateRequestPort output.ValidateRequest) *validateRequest {
	return &validateRequest{validateRequestPort: validateRequestPort}
}

func (vl *validateRequest) ValidateTransactionService() {
	sum, err := vl.validateRequestPort.ValidateRequestPort(context.Background(), &integrator.IntegratorRequest{
		FirstNumber:  0,
		SecondNumber: 0,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(sum)
}
