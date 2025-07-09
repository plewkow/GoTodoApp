package dto

import (
	"draft-zadania-1/models"
)

func ToUser(dto UserDTO) models.User {
	return models.User{
		Username: dto.Username,
		Email:    dto.Email,
	}
}

func ToResponseUserDTO(user models.User) ResponseUserDTO {
	return ResponseUserDTO{
		Id:       user.Id.String(),
		Username: user.Username,
		Email:    user.Email,
	}
}

func ToResponseUserDTOs(users []models.User) []ResponseUserDTO {
	var dtos []ResponseUserDTO
	for _, t := range users {
		dtos = append(dtos, ToResponseUserDTO(t))
	}
	return dtos
}
