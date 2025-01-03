package domains

import (
	"github.com/fadhlinw/backend-golang-test/dto"
	"github.com/fadhlinw/backend-golang-test/utils"
	"gorm.io/gorm"
)

type UserService interface {
	WithTrx(trxHandle *gorm.DB) UserService
	GetOneUserById(id int) (dto dto.UserResponseDto, err error)
	GetOneUserByEmail(email string) (dto dto.UserResponseDto, err error)
	GetAllUser(searchQuery string, pagination utils.Pagination) (utils.Pagination, error)
	CreateUser(createUserDto dto.CreateUserRequest) error
	UpdateUser(id int, updateUserDto dto.CreateUserRequest) error
	DeleteUser(id int) error
}
