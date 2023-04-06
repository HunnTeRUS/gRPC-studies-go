package configmap

import "awesomeProject/adapter/output/grpc_client_service"

type ModuleConfiguration struct {
	IsEnable          bool                                `yaml:"is_enable" json:"is_enable"`
	AllowedProgramIds []int64                             `yaml:"allowed_program_ids" json:"allowed_program_ids"`
	AuthInterface     grpc_client_service.ValidateRequest `yaml:"-" json:"-"`
}
