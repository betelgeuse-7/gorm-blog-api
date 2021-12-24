package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/betelgeuse-7/gorm-blog-api/models"
)

func NewUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db := models.GetDB()
	if res := db.Create(&user); res.Error != nil {
		errStr := res.Error.Error()
		http.Error(w, errStr, http.StatusInternalServerError)
		log.Printf("[DB_ERROR] %s\n", errStr)
		return
	}

	log.Printf("added user with id %d\n", user.ID)

	w.WriteHeader(http.StatusCreated)
}
