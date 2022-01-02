package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/betelgeuse-7/gorm-blog-api/models"
	"github.com/go-chi/chi"
	"gorm.io/gorm"
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

	w.WriteHeader(http.StatusCreated)
}

func GetUserWithId(w http.ResponseWriter, r *http.Request) {
	var user models.User
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errStr := err.Error()
		http.Error(w, errStr, http.StatusInternalServerError)
		log.Printf("[ERROR] %s\n", errStr)
		return
	}
	db := models.GetDB()
	res := db.First(&user, userId)
	if err := res.Error; err != nil {
		errStr := err.Error()
		if err == gorm.ErrRecordNotFound {
			http.Error(w, errStr, http.StatusNotFound)
		} else {
			http.Error(w, errStr, http.StatusInternalServerError)
		}
		log.Printf("[ERROR] %s\n", errStr)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func DeleteUserWithId(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errStr := err.Error()
		http.Error(w, errStr, http.StatusInternalServerError)
		log.Printf("[ERROR] %s\n", errStr)
		return
	}
	db := models.GetDB()
	res := db.Debug().Delete(&models.User{}, userId)
	if err := res.Error; err != nil {
		errStr := res.Error.Error()
		http.Error(w, errStr, http.StatusBadRequest)
		log.Printf("[ERROR] %s\n", errStr)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UpdateUserWithId(w http.ResponseWriter, r *http.Request) {
	var u models.User
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errStr := err.Error()
		http.Error(w, errStr, http.StatusInternalServerError)
		log.Printf("[ERROR] %s\n", errStr)
		return
	}
	db := models.GetDB()
	// read body (new username)
	json.NewDecoder(r.Body).Decode(&u)
	// i will not check if the new username is the same as the old one
	res := db.First(&models.User{}, userId).Update("username", u.Username)
	if err := res.Error; err != nil {
		errStr := err.Error()
		http.Error(w, errStr, http.StatusBadRequest)
		log.Printf("[ERROR] %s\n", errStr)
		return
	}
	w.WriteHeader(http.StatusOK)
}
