package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/betelgeuse-7/gorm-blog-api/models"
	"github.com/go-chi/chi"
)

func GetPostWithId(w http.ResponseWriter, r *http.Request) {
	var post models.Post

	postId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errStr := err.Error()
		log.Printf("[ERROR_1] %s\n", errStr)
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}
	db := models.GetDB()
	//db.Preload("Author").First(&post, postId)
	db.First(&post, postId)
	json.NewEncoder(w).Encode(post)
}

func DeletePostWithId(w http.ResponseWriter, r *http.Request) {

}

func UpdatePostWithId(w http.ResponseWriter, r *http.Request) {

}

func NewPost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		errStr := err.Error()
		log.Printf("[ERROR] %s\n", errStr)
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}
	db := models.GetDB()
	if res := db.Create(&post); res.Error != nil {
		errStr := res.Error.Error()
		log.Printf("[ERROR] %s\n", errStr)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
