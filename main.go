package main 

import (
    // "fmt"
    // "database/sql"
    "github.com/gin-gonic/gin"
    "github.com/Cguilliman/test_blunk_db/database/requests"
    "github.com/Cguilliman/test_blunk_db/database/models"
    db "github.com/Cguilliman/test_blunk_db/database/base"
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
    //     Username:  "James1",
    //     FirstName: "fName2",
    //     LastName:  "lName3",
    // })
    requests.UpdatePerson(5, &models.Person{
        Username:  "Jamesnew",
        FirstName: "fNamenew",
        LastName:  "lNamenew",
    })
    engine := gin.Default()
    engine.Run("0.0.0.0:8000")
}
