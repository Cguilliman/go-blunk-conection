package controllers

import (
    "github.com/gin-gonic/gin"

    "github.com/Cguilliman/chat/controllers/users"
    "github.com/Cguilliman/chat/controllers/middlewares"
    "github.com/Cguilliman/chat/controllers/pages"
)

func InitRoutings(engine *gin.Engine) {
    render := engine.Group("")
    render.GET("", pages.Main)
    
    v1 := engine.Group("api/v1")
    v1.POST("users/registration", users.Registration)
    v1.POST("users/login", users.Login)
    v1.Use(middlewares.AuthMiddleware(false))
    v1.GET("users/receive", users.Receive)
}
