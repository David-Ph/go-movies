package main

import (
	"net/http"
	"os"

	"moviesnow-backend/helper"

	"github.com/joho/godotenv"
)

func NewServer() *http.Server {

	return &http.Server{
		Addr: os.Getenv("HOST_URL"),
	}
}

func main() {
	err := godotenv.Load()
	helper.PanicIfError(err)

	server := InitializeServer()
	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
