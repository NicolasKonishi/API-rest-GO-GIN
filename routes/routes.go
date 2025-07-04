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
	r.Run()
}
