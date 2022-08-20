package main

import (
	"go-react-auth/configs"
	"go-react-auth/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine = gin.Default()

func main() {

	controllers.AuthController()
	controllers.UserController()

	port, err := configs.EnvPORT()

	if err != nil {
		log.Fatal("Bad Port")
	}

	configs.Router.Run(port)
}
