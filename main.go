package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/betelgeuse-7/gorm-blog-api/models"
	"github.com/betelgeuse-7/gorm-blog-api/routes"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	portFlag := flag.Int("port", 8000, "port to listen to")
	flag.Parse()
	port := fmt.Sprintf(":%d", *portFlag)

	pg := Postgres{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}

	db, err := pg.Open()
	if err != nil {
		log.Printf("[FAIL] error while connecting to db (%s): %s\n", pg.DSN(), err.Error())
		log.Println("shutting down the server...")
		return
	}

	db.AutoMigrate(&models.Post{}, &models.User{})
	models.RegisterDB(db)

	routes := routes.Routes()

	log.Printf(`
		========================================
		Server started listening on 127.0.0.1%s
		========================================`, port)
	http.ListenAndServe(port, routes)
}
