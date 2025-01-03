package services

import (
	"github.com/fadhlinw/backend-golang-test/domains"
	"github.com/fadhlinw/backend-golang-test/lib"
	"github.com/fadhlinw/backend-golang-test/models"
	"github.com/fadhlinw/backend-golang-test/repository"

	"gorm.io/gorm"
)

type TokenStoreService struct {
	logger     lib.Logger
	repository repository.TokenStoreRepository
}

func NewTokenStoreService(logger lib.Logger, repository repository.TokenStoreRepository) domains.TokenStoreService {
	return TokenStoreService{
		logger:     logger,
		repository: repository,
	}
}

func (s TokenStoreService) WithTrx(trxHandle *gorm.DB) domains.TokenStoreService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

func (s TokenStoreService) CreateToken(tokenStore models.TokenStore) error {
	return s.repository.CreateToken(&tokenStore)
}

func (s TokenStoreService) DeleteToken(token string) error {
	return s.repository.DeleteToken(token)
}

func (s TokenStoreService) ValidateToken(token string) error {
	return s.repository.ValidateToken(token)
}
