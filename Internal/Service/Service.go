package Service

import "CSR/Internal/Repository"

type Service struct {
	repository *Repository.Repository
}

func NewService(repository *Repository.Repository) *Service {
	return &Service{repository: repository}
}
