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
