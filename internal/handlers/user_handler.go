package handlers

import (
	"fmt"
	"log"

	"github.com/ReynirPY/library-managment-system/config"
	"github.com/ReynirPY/library-managment-system/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func InsertUser(user models.User) error {
	//take password from request and generate hash
	password := []byte(user.Password)
	hashPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error during hashing password %w", err)
	}
	log.Println(string(hashPassword))

	//put hash inot user struct from request and incert to database
	user.Password = string(hashPassword)
	_, err = config.DB.NamedExec("INSERT INTO users (username, password, email) VALUES (:username, :password, :email)", &user)
	if err != nil {
		return fmt.Errorf("error during insert user %w", err)
	}

	return nil
}

func GetUserByCredentials(email string) (models.User, error) {
	var user models.User
	err := config.DB.Get(&user, "SELECT * FROM users WHERE email=$1", email)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to get user from db %w", err)
	}

	return user, nil
}

func GetUserByID(id int) (models.User, error) {
	var user models.User
	err := config.DB.Get(&user, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to get user from db %w", err)
	}

	return user, nil
}
