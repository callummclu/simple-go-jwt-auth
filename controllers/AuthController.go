package controllers

import (
	"go-react-auth/configs"
	"go-react-auth/services"
)

func AuthController() {
	api := (configs.Router).Group("auth")
	{
		api.POST("login", services.LogUserIn)
		api.GET(":token", services.CheckUserStatus)
	}
}
