package user

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var validate = validator.New()

func GetUser(db *gorm.DB) ([]UserResponse, error) {
	return Get(db)
}

func ShowUser(db *gorm.DB, id uint) (*User, error) {
	return FindById(db, id)
}

func DeleteUser(db *gorm.DB, id uint) error {
	return Delete(db, id)
}

func UpdateUser(db *gorm.DB, id uint, updated map[string]interface{}) error {
	if passwordRaw, ok := updated["password"]; ok {
		// Konversi interface{} ke string
		passwordStr, ok := passwordRaw.(string)
		if !ok {
			return fmt.Errorf("password must be a string")
		}

		// Hash password baru
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordStr), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		// Ganti value "password" dengan hasil hash
		updated["password"] = string(hashedPassword)
	}
	return Update(db, id, updated)
}

func StoreUser(db *gorm.DB, store UserStore) error {
	if err := validate.Struct(store); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(store.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := User{
		Username: store.Username,
		Email:    store.Email,
		Password: string(hashedPassword),
	}

	return Store(db, &user)

}
