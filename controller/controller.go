package controller

import "github.com/gin-gonic/gin"

func GetAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		//"message": "Hello World",
		//retorna uma mensagem utilizando o gin
		"id":   "1",
		"nome": "nicolas",
	})
}

func Hello(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz : ": "Olá " + nome + ", tudo bem?"})
}
