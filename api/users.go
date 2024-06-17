package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ReynirPY/library-managment-system/internal/auth"
	"github.com/ReynirPY/library-managment-system/internal/handlers"
	"github.com/ReynirPY/library-managment-system/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func UserRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "wrong input", http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	err = handlers.InsertUser(user)
	if err != nil {
		http.Error(w, "error during create user", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials models.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "wrong input", http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	log.Println(credentials)

	user, err := handlers.GetUserByCredentials(credentials.Login)
	if err != nil {
		http.Error(w, "error during getting user in db", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("logged in")

	tokenString, err := auth.CreateToken(user.ID, user.Username, "admin")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "failed to create token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})

}
