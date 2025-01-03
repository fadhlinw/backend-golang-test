package services

import (
	"net/http"
	"strings"

	"github.com/fadhlinw/backend-golang-test/constants"
	"github.com/fadhlinw/backend-golang-test/domains"
	"github.com/fadhlinw/backend-golang-test/dto"
	httperror "github.com/fadhlinw/backend-golang-test/error"
	"github.com/fadhlinw/backend-golang-test/lib"
	"github.com/fadhlinw/backend-golang-test/mapper"
	"github.com/fadhlinw/backend-golang-test/repository"
	"github.com/fadhlinw/backend-golang-test/utils"

	"gorm.io/gorm"
)

// UserService handles business logic for users
type UserService struct {
	logger     lib.Logger
	repository repository.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(logger lib.Logger, repository repository.UserRepository) domains.UserService {
	return UserService{
		logger:     logger,
		repository: repository,
	}
}

// WithTrx sets transaction for the repository
func (s UserService) WithTrx(trxHandle *gorm.DB) domains.UserService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

func (s UserService) GetOneUserById(id int) (dto dto.UserResponseDto, err error) {
	user, err := s.repository.GetByID(id)
	if err != nil {
		return dto, httperror.NewHttpError(constants.ERROR_USER_NOT_FOUND, "", http.StatusNotFound)
	}
	return mapper.ToUserResponseDto(*user), nil
}

func (s UserService) GetOneUserByEmail(email string) (dto dto.UserResponseDto, err error) {
	user, err := s.repository.GetByEmail(email)
	if err != nil {
		return dto, httperror.NewHttpError(constants.ERROR_USER_NOT_FOUND, "", http.StatusNotFound)
	}
	return mapper.ToUserResponseDto(*user), nil
}

func (s UserService) GetOneUserByUsername(username string) (dto dto.UserResponseDto, err error) {
	user, err := s.repository.GetByUsername(username)
	if err != nil {
		return dto, httperror.NewHttpError(constants.ERROR_USER_NOT_FOUND, "", http.StatusNotFound)
	}
	return mapper.ToUserResponseDto(*user), nil
}

func (s UserService) GetAllUser(searchQuery string, pagination utils.Pagination) (utils.Pagination, error) {

	query, users := s.repository.GetAll(searchQuery)

	err := query.Scopes(paginate(&users, &pagination, query)).
		Order(pagination.GetSort()).Find(&users).Error
	if err != nil {
		return pagination, httperror.NewHttpError(constants.ERROR_GETTING_USER, "", http.StatusInternalServerError)
	}

	pagination.Rows = mapper.ToUsersResponseDto(users)
	return pagination, nil
}

func (s UserService) CreateUser(createUserDto dto.CreateUserRequest) error {
	createUserDto.Email = strings.ToLower(createUserDto.Email)
	createUserDto.Username = strings.ToLower(createUserDto.Username)
	// password := utils.GenerateRandomString(12)
	user, err := mapper.ToUserModel(createUserDto)
	if err != nil {
		s.logger.Error("Error when creating customer: ", err)
		return err
	}
	// user.Password = password

	// sendEmailRequest := dto.SendEmailRequestDto{
	// 	To:      createUserDto.Email,
	// 	Subject: "Hello",
	// 	Body:    "Your password is: " + password,
	// }

	// err = s.smtpClient.SendEmail(sendEmailRequest)
	// if err != nil {
	// 	return httperror.NewHttpError(constants.ERROR_SENDING_EMAIL, "", http.StatusInternalServerError)
	// }
	return s.repository.Create(&user)
}

func (s UserService) UpdateUser(id int, updateUserDto dto.CreateUserRequest) error {
	user, err := mapper.ToUserModel(updateUserDto)
	if err != nil {
		s.logger.Error("Error when creating customer: ", err)
		return err
	}
	return s.repository.Update(id, &user)
}

func (s UserService) DeleteUser(id int) error {
	return s.repository.Delete(id)
}