package controllers

import (
	authentication_service "APImod/src/core/services/authentication-service"
	"APImod/src/presentation-layer/contracts/authentication"
	"github.com/gin-gonic/gin"
)

type ControllerAuthResult struct{}

func (cAr *ControllerAuthResult) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request authentication.RegisterRequest
		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid request format" + err.Error()})
			return
		}
		authResult := &authentication_service.AuthenticationResult{}
		result, err := authResult.Register(request)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err})
			return
		}
		if result {
			ctx.JSON(200, gin.H{"message": "Регистрация успешна"})
		} else {
			ctx.JSON(400, gin.H{"message": err})
		}
	}
}

func (cAr *ControllerAuthResult) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request authentication.LoginRequest
		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid request format" + err.Error()})
			return
		}
		authResult := &authentication_service.AuthenticationResult{}
		result, err := authResult.Login(request)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if result {
			ctx.JSON(200, gin.H{"message": "Авторизация успешна"})
		} else {
			ctx.JSON(400, gin.H{"message": err})
		}
	}
}
