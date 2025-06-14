package auth

import (
	"fmt"
	user "pbl409-dashboard/pkg/users"
	"pbl409-dashboard/pkg/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtSecret = utils.GetJWTSecret()

func Login(db *gorm.DB, input user.LoginDTO) (string, error) {
	foundUser, err := user.FindByUsername(db, input.Username)
	if err != nil {
		return "", fmt.Errorf("invalid username or password")
	}
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(input.Password))
	if err != nil {
		return "", fmt.Errorf("invalid username or password")
	}

	claims := jwt.MapClaims{
		"sub":      fmt.Sprintf("%d", foundUser.ID),
		"user_id":  foundUser.ID,
		"username": foundUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return signedToken, nil
}
