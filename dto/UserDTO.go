package dto

type CreateUserDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type ResponseUserDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UpdateUserDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
