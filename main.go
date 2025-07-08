package main

import (
	"02Alura/database"
	"02Alura/models"
	"02Alura/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Nicolas", CPF: "123232313", RG: "123232313"},
		{Nome: "Lucas", CPF: "121212121", RG: "1111111111"},
	}
	routes.HandleRequest()
}
