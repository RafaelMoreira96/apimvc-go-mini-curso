package main

import (
	"catalogolivros/database"
	"catalogolivros/server"
)

func main() {
	database.Connect()

	server := server.NewServer()

	server.Run()
}
