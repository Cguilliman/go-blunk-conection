package main 

import (
    // "fmt"
    // "database/sql"
    "github.com/gin-gonic/gin"
    "github.com/Cguilliman/test_blunk_db/database"
    db "github.com/Cguilliman/test_blunk_db/database/base"
)

func main() {
    db := db.Init()
    defer db.Close()
    database.GetMessage()
    database.GetRooms()
    engine := gin.Default()
    engine.Run("0.0.0.0:8000")
}
