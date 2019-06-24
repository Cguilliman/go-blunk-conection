package controllers

import (
    "github.com/gin-gonic/gin"

    "github.com/Cguilliman/chat/controllers/users"
)

func InitRoutings(engine *gin.Engine) {
    v1 := engine.Group("api/v1")
    v1.POST("users/registration", users.Registration)
}
