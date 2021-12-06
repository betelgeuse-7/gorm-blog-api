package routes

import (
	"net/http"

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

	return r
}
