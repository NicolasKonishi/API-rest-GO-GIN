package routes

import (
	"02Alura/controller"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	//r.GET("/", func(c *gin.Context) {})
	r.GET("/alunos", controller.GetAlunos)
	r.GET("/:nome", controller.Hello)
	r.POST("/alunos", controller.CreateAlunos)
	r.GET("/alunos/:id", controller.GetAlunoByID)
	r.DELETE("/alunos/:id", controller.DeleteAlunos)
	r.PATCH("/alunos/:id", controller.EditAlunos)
	r.GET("/alunos/cpf/:cpf", controller.GetAlunoByCpf)

	r.Run()
}
