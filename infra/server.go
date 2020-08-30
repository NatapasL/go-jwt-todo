package infra

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	ConnectToDb()
	db := GetDB()
	defer db.Close()

	router := gin.Default()
	mapRoutes(router)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	addr := fmt.Sprintf("%s:%s", host, port)
	log.Fatal(router.Run(addr))
}
