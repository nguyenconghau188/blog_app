package main

import (
	"blogs/db"
	"blogs/internal/router"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.InitDB()

	router.RegisterUserRoutes(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
