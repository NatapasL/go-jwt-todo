package main

import (
	"os"

	"go-jwt-todo/infra"
)

func init() {
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	os.Setenv("REFRESH_SECRET", "mkjdsfjklsdj")
}

func main() {
	infra.StartServer()
}
