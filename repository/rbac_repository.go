package repository

import (
	"github.com/fadhlinw/backend-golang-test/lib"
	"github.com/fadhlinw/backend-golang-test/models"
	"gorm.io/gorm"
)

// RBAC Repository database structure
type RBACRepository struct {
	lib.Database
	logger lib.Logger
}

// NewRBACRepository creates a new RBAC repository
func NewRBACRepository(db lib.Database, logger lib.Logger) RBACRepository {
	return RBACRepository{
		Database: db,
		logger:   logger,
	}
}

// WithTrx enables repository with transaction
func (r RBACRepository) WithTrx(trxHandle *gorm.DB) RBACRepository {
	if trxHandle == nil {
		r.logger.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}

// GetPermission
func (r RBACRepository) GetPermission(roleId int, slug string, action string) (*models.Permission, error) {
	var permission models.Permission
	return &permission, r.Database.Where("role_id = ? AND slug = ? AND action = ?", roleId, slug, action).First(&permission).Error
}
