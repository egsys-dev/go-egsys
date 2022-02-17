package main

import (
	"github.com/fabiom2211/go-egsys/database"
	"github.com/fabiom2211/go-egsys/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
