package routes

import (
	"github.com/betelgeuse-7/gorm-blog-api/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var r *chi.Mux = chi.NewRouter()

func Routes() *chi.Mux {
	r.Use(middleware.Logger)

	r.Get("/api/user/{id}", controllers.GetUserWithId)
	r.Delete("/api/user/{id}", controllers.DeleteUserWithId)
	r.Put("/api/user/{id}", controllers.UpdateUserWithId)
	r.Post("/api/user/new", controllers.NewUser)

	return r
}
