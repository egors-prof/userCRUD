package service

import (
	"CSR/internal/contracts"
)

type Service struct {
	repository contracts.RepositoryI
	cache contracts.CacheI
}

func NewService(repository contracts.ServiceI,cache contracts.CacheI) *Service {
	return &Service{repository: repository,cache: cache}
}
