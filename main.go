package main

import (
	"os"

	"github.com/NatapasL/go-jwt-todo/infra"
)

func init() {
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	os.Setenv("REFRESH_SECRET", "mkjdsfjklsdj")
}

func main() {
	infra.StartServer()
}
