package mapper

import (
	"github.com/fadhlinw/backend-golang-test/dto"
	"github.com/fadhlinw/backend-golang-test/models"
	"github.com/fadhlinw/backend-golang-test/utils"
)

func ToUsersResponseDto(users []models.User) []dto.UserResponseDto {
	user := make([]dto.UserResponseDto, len(users))
	for i, item := range users {
		user[i] = ToUserResponseDto(item)
	}
	return user
}

func ToUserResponseDto(user models.User) dto.UserResponseDto {
	return dto.UserResponseDto{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		Username:  user.Username,
		Role:      user.Role,
		RoleID:    user.RoleID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserModel(user dto.CreateUserRequest) (models.User, error) {
	password, err := utils.HashPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}

	if user.RoleID == 1 {
		user.Role = "owner"
	} else if user.RoleID == 2 {
		user.Role = "librarian"
	} else if user.RoleID == 3 {
		user.Role = "visitor"
	}

	return models.User{
		FullName: user.FullName,
		Email:    user.Email,
		Username: user.Username,
		Password: password,
		Role:     user.Role,
		RoleID:   user.RoleID,
	}, nil
}
