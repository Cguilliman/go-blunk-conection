package main 

import (
    "github.com/gin-gonic/gin"
    db "github.com/Cguilliman/chat/database/base"
    "github.com/Cguilliman/chat/controllers"
    "github.com/Cguilliman/chat/settings"
)

func main() {
    db := db.Init()
    defer db.Close()
    engine := gin.Default()
    engine.LoadHTMLGlob(settings.Static)

    controllers.InitRoutings(engine)
    engine.Run(settings.Run)
}
