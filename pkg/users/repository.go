package user

import "gorm.io/gorm"

func Get(db *gorm.DB) ([]UserResponse, error) {
	var users []User

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	var userResponse []UserResponse

	for _, u := range users {
		userResponse = append(userResponse, UserResponse{
			ID:       int(u.ID),
			Username: u.Username,
			Email:    u.Email,
		})
	}

	return userResponse, nil
}

func FindByUsername(db *gorm.DB, username string) (*User, error) {
	var user User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func FindById(db *gorm.DB, id uint) (*User, error) {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func Delete(db *gorm.DB, id uint) error {
	var user User

	if err := db.First(&user, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&user, id).Error; err != nil {
		return err
	}

	return nil

}
func Update(db *gorm.DB, id uint, updated map[string]interface{}) error {
	var user User

	if err := db.First(&user, id).Error; err != nil {
		return err
	}

	if err := db.Model(&user).Updates(updated).Error; err != nil {
		return err
	}

	return nil
}
func Store(db *gorm.DB, user *User) error {
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}
