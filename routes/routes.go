package routes

import (
	"net/http"

	"github.com/betelgeuse-7/gorm-blog-api/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var r *chi.Mux = chi.NewRouter()

func Routes() *chi.Mux {
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/html")
		w.Write([]byte("<h1>Hello</h1>"))
	})

	r.Post("/api/user/new", controllers.NewUser)

	return r
}
