package services

import (
	"github.com/fadhlinw/backend-golang-test/domains"
	"github.com/fadhlinw/backend-golang-test/lib"
	"github.com/fadhlinw/backend-golang-test/repository"
	"gorm.io/gorm"
)

// RBACService service layer
type RBACService struct {
	env        lib.Env
	logger     lib.Logger
	repository repository.RBACRepository
}

// NewRBACService creates new rbac service
func NewRBACService(
	env lib.Env,
	logger lib.Logger,
	repository repository.RBACRepository) domains.RBACService {
	return RBACService{
		env:        env,
		logger:     logger,
		repository: repository,
	}
}

// WithTrx delegates transaction to repository database
func (s RBACService) WithTrx(trxHandle *gorm.DB) domains.RBACService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

// HasPermission check permission
func (s RBACService) HasPermission(roleId int, slug string, action string) (bool, error) {
	_, err := s.repository.GetPermission(roleId, slug, action)
	if err != nil {
		s.logger.Error("Error when checking permission by role_id: ", roleId)
		s.logger.Debug("Details: ", err.Error())
		return false, err
	}

	return true, nil
}
