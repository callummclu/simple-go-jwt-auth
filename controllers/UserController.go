package controllers

import (
	"go-react-auth/configs"
	"go-react-auth/services"
)

func UserController() {
	api := (configs.Router).Group("user")
	{
		api.POST("", services.CreateNewUser)
	}
}
