package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/betelgeuse-7/gorm-blog-api/routes"
)

func main() {
	portFlag := flag.Int("port", 8000, "port to listen to")
	flag.Parse()
	port := fmt.Sprintf(":%d", *portFlag)

	routes := routes.Routes()

	log.Println("========================================")
	log.Printf("Server started listening on 127.0.0.1%s\n", port)
	log.Println("========================================")
	http.ListenAndServe(port, routes)
}
