package domains

import "github.com/fadhlinw/backend-golang-test/models"

type OTPService interface {
	Create(userId int, code string) error
	UpdateById(id int, isUsed bool) error
	GetByCode(userId int, code string) (*models.Otp, error)
	GetByUserIdAndIsUsed(userId int, isUsed bool) (*models.Otp, error)
}
