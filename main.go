package main

import (
	"moviesnow-backend/helper"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	helper.PanicIfError(err)

	server := InitializeServer()
	err = server.Start(os.Getenv("HOST_URL"))
	helper.PanicIfError(err)
}
