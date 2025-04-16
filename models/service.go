package models

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Name        string
	Host        string
	Type        string
	Username    string
	Password    string
	ActiveToken string
}
