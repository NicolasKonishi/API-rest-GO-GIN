package controller

import (
	"02Alura/database"
	"02Alura/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAlunos(c *gin.Context) {
	//c.JSON(200, models.Alunos)
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)

}

func Hello(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz : ": "Olá " + nome + ", tudo bem?"})
}

func CreateAlunos(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func GetAlunoByID(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	c.JSON(http.StatusOK, aluno)
	//if err := database.DB.Where("id = ?", id).First(&aluno).Error; err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	//	return
	//}
}
