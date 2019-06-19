package database

import (
    "fmt"
    // "reflect"
    "github.com/Cguilliman/test_blunk_db/database/models"
    "database/sql"
)

func GetMessage(db *sql.DB) {
    response, err := models.Query(
        func(rows *sql.Rows) (models.Model, error) {
            message := new(models.Message)
            message.Room = *new(models.Room)
            return message, rows.Scan(
                &message.ID, &message.Message, &message.IsRead,
                &message.RoomID, &message.FromID, &message.ToID,
                &message.Room.Name,
            )
        },
        `SELECT 
            msg.ID, msg.Message, msg.IsRead, 
            msg.RoomID, msg.FromID, msg.ToID,
            Room.Name as RoomName
        FROM Message as msg 
        INNER JOIN Room 
        ON msg.RoomID = Room.ID`, 
        db,  // TODO remove
    )
    if err != nil {
        fmt.Println(err)
    }
    messages := models.ResponseConvert(response)
    fmt.Println(messages[0].ID)
}
