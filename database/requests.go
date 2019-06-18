package database

import (
    "fmt"
    "github.com/Cguilliman/test_blunk_db/database/models"
    "database/sql"
)

func GetMessage(db *sql.DB) {
    response, err := models.Query(
        models.ScanMessage, 
        `SELECT 
            ID, Message, IsRead, 
            RoomID, FromID, ToID 
        FROM Message`, 
        db,  // TODO remove
    )
    if err != nil {
        fmt.Println(err)
    }
    message := response[0].(*models.Message)
    fmt.Println(message.ID)
}
