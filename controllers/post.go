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

func GetPostWithId(w http.ResponseWriter, r *http.Request) {
	postId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errStr := err.Error()
		log.Printf("[ERROR_1] %s\n", errStr)
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}
	var post models.Post
	db := models.GetDB()

	if res := db.Joins("Author").First(&post, postId); res.Error != nil {
		err := res.Error
		if err == gorm.ErrRecordNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(post)
}

func DeletePostWithId(w http.ResponseWriter, r *http.Request) {
	postId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errStr := err.Error()
		log.Printf("[ERROR_1] %s\n", errStr)
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}
	db := models.GetDB()
	res := db.Delete(&models.Post{}, postId)
	if res.Error != nil {
		errStr := res.Error.Error()
		log.Printf("[ERROR] %s\n", errStr)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UpdatePostWithId(w http.ResponseWriter, r *http.Request) {
	var p models.Post
	postId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errStr := err.Error()
		http.Error(w, errStr, http.StatusInternalServerError)
		log.Printf("[ERROR] %s\n", errStr)
		return
	}
	db := models.GetDB()
	json.NewDecoder(r.Body).Decode(&p)
	res := db.First(&models.Post{}, postId).Updates(models.Post{Title: p.Title, Content: p.Content})
	if res.Error != nil {
		errStr := res.Error.Error()
		log.Printf("[ERROR] %s\n", errStr)
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
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

// QUERY PARAMS:
// 	- newerFirst
//	- limit
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post

	newerPostsFirstQueryParam := r.URL.Query().Get("newerFirst")
	limitQueryParam, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		log.Printf("[limitQueryParamAtoiError] %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := models.GetDB()
	if newerPostsFirstQueryParam == "true" || newerPostsFirstQueryParam == "1" {
		db.Joins("Author").Limit(limitQueryParam).Order("created_at DESC").Find(&posts)
	} else {
		db.Joins("Author").Limit(limitQueryParam).Find(&posts)
	}

	json.NewEncoder(w).Encode(posts)
}
