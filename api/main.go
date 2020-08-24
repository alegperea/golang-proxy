package main

import (

	controllers "github.com/aperea/go-mlproxy/api/controllers"
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	r.GET("/getrule", controllers.RulesHandler)
	r.Run(":9090")
}
