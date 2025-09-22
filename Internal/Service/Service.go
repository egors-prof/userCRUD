package Service

import (
	"CSR/Internal/contracts"
)

type Service struct {
	repository contracts.RepositoryI
}

func NewService(repository contracts.ServiceI) *Service {
	return &Service{repository: repository}
}
