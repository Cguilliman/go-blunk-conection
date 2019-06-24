package main 

import (
    // "fmt"
    // "database/sql"
    "github.com/gin-gonic/gin"
    // "github.com/Cguilliman/chat/database/requests"
    // "github.com/Cguilliman/chat/database/models"
    db "github.com/Cguilliman/chat/database/base"
    "github.com/Cguilliman/chat/controllers"
)

func main() {
    db := db.Init()
    defer db.Close()
    // requests.GetMessage()
    // requests.GetRooms()
    // requests.MainUserRoomsList(1)

    // requests.GetMessages(1, 2)
    // requests.GetMessages(2, 2)
    // requests.GetMessages(3, 2)
    
    // requests.GetPerson(1)
    // requests.CreatePerson(&models.Person{
    //     Username:  "J2",
    //     FirstName: "fName2",
    //     LastName:  "lName3",
    // })
    // requests.UpdatePerson(5, &models.Person{
    //     Username:  "Jamesnew",
    //     FirstName: "fNamenew",
    //     LastName:  "lNamenew",
    // })
    engine := gin.Default()
    controllers.InitRoutings(engine)
    engine.Run("0.0.0.0:8000")
}
