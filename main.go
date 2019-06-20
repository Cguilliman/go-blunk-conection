package main 

import (
    // "fmt"
    // "database/sql"
    "github.com/gin-gonic/gin"
    "github.com/Cguilliman/test_blunk_db/database/requests"
    db "github.com/Cguilliman/test_blunk_db/database/base"
)

func main() {
    db := db.Init()
    defer db.Close()
    // requests.GetMessage()
    // requests.GetRooms()
    // requests.MainUserRoomsList(1)

    requests.GetMessages(1, 2)
    requests.GetMessages(2, 2)
    requests.GetMessages(3, 2)
    
    engine := gin.Default()
    engine.Run("0.0.0.0:8000")
}
