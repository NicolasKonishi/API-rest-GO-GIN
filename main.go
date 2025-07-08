package main

import (
	"02Alura/database"
	"02Alura/routes"
)

func main() {
	database.ConectaComBancoDeDados()

	routes.HandleRequest()
}
