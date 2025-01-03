package domains

import "gorm.io/gorm"

type RBACService interface {
	WithTrx(trxHandle *gorm.DB) RBACService
	HasPermission(roleId int, resource string, action string) (bool, error)
}
