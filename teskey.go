package main

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

const key = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDk3MDAzMTYsImlhdCI6MTc0OTYxMzkxNiwidXNlcl9pZCI6MiwidXNlcm5hbWUiOiJ1ZGluQWhheSJ9.atnAGzoVR4tJK9YKescAqEAO7uhpRQa9_CI0ygbm-_8"

func PrintKey(k string) {
	fmt.Print(key)
}

func maina() {
	token, err := jwt.Parse(key, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		fmt.Println("Token error:", err)
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("Token valid, claims:", claims)
	} else {
		fmt.Println("Token invalid")
	}
}
