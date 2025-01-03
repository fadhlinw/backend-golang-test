package domains

import (
	"github.com/fadhlinw/backend-golang-test/models"
	"gorm.io/gorm"
)

type TokenStoreService interface {
	WithTrx(trxHandle *gorm.DB) TokenStoreService
	CreateToken(tokenStore models.TokenStore) error
	DeleteToken(token string) error
	ValidateToken(token string) error
}
