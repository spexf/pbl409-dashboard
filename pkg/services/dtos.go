package service

type ServiceStore struct {
	Name     string `json:"name" validate:"required"`
	Host     string `json:"host" validate:"required"`
	Type     string `json:"type" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
