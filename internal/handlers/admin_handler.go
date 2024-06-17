package handlers

import (
	"fmt"

	"github.com/ReynirPY/library-managment-system/config"
	"github.com/ReynirPY/library-managment-system/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func InsertAdmin(user models.Admin) error {
	//take password from request and generate hash
	password := []byte(user.Password)
	hashPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error during hashing password %w", err)
	}
	//log.Println(string(hashPassword))

	//put hash inot user struct from request and incert to database
	user.Password = string(hashPassword)
	_, err = config.DB.NamedExec("INSERT INTO admins (username, password, email) VALUES (:username, :password, :email)", &user)
	if err != nil {
		return fmt.Errorf("error during insert admin %w", err)
	}

	return nil
}

func GetAdminByCredentials(email string) (models.Admin, error) {
	var admin models.Admin
	err := config.DB.Get(&admin, "SELECT * FROM admins WHERE email=$1", email)
	if err != nil {
		return models.Admin{}, fmt.Errorf("failed to get admin from db %w", err)
	}

	return admin, nil
}
