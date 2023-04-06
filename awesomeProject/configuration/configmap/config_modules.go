package configmap

import "awesomeProject/adapter/output/grpc_client_service"

var Configuration ExternalModules

type ExternalModules struct {
	MapAuthorization map[string]ModuleConfiguration
	ResponseCodeMap  map[string]ResponseCodeRules
}

type ResponseCodeRules struct {
	IsApprovalCode bool
	IsReferralCode bool
	IsDenialCode   bool
}

type BrandResponseCode struct {
	Mastercard string
	Visa       string
}

type ModuleConfiguration struct {
	IsEnable          bool                                `yaml:"is_enable" json:"is_enable"`
	AllowedProgramIds []int64                             `yaml:"allowed_program_ids" json:"allowed_program_ids"`
	DecisionWeight    int8                                `yaml:"decision_weight" json:"decision_weight"`
	AuthInterface     grpc_client_service.ValidateRequest `yaml:"-" json:"-"`
}
