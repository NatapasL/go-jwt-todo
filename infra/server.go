package infra

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	mapRoutes(router)

	host := os.Getenv("HOST")
	if len(host) <= 0 {
		host = "127.0.0.1"
	}

	port := os.Getenv("PORT")
	if len(port) <= 0 {
		port = "3000"
	}

	addr := fmt.Sprintf("%s:%s", host, port)
	log.Fatal(router.Run(addr))
}
