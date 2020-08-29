package infra

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type server struct {
	router *gin.Engine
	addr   string
}

func (s *server) init() {
	router := gin.Default()
	mapRoutes(router)
	s.router = router

	host := os.Getenv("HOST")
	if len(host) <= 0 {
		host = "127.0.0.1"
	}
	port := os.Getenv("PORT")
	if len(port) <= 0 {
		port = "3000"
	}
	s.addr = fmt.Sprintf("%s:%s", host, port)
}

func (s *server) Start() {
	log.Fatal(router.Run(addr))
}
