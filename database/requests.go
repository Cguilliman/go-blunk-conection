package database

import (
    "fmt"
    // "reflect"
    "github.com/Cguilliman/test_blunk_db/database/models"
    "github.com/Cguilliman/test_blunk_db/database/base"
    // "database/sql"
)

func GetMessage() {
    rows, _ := base.GetDB().Query("SELECT * FROM Message")
    messages := make([]*models.Message, 0)
    for rows.Next() {
        message := new(models.Message)
        _ = rows.Scan(
            &message.ID, &message.Message, &message.IsRead,
            &message.RoomID, &message.FromID, &message.ToID,
        )
        messages = append(messages, message)
    }
    for _, obj := range messages {
        fmt.Println(obj.ID)
        fmt.Println(obj.ID)
    }

    // queryset := new(models.MessageQuerySet).MakeQuery(
    //     `SELECT * FROM Message`, 
    //     func(rows *sql.Rows) (*models.Message, error) {
    //         message := new(models.Message)
    //         message.Room = *new(models.Room)
    //         return message, rows.Scan(
    //             &message.ID, &message.Message, &message.IsRead,
    //             &message.RoomID, &message.FromID, &message.ToID,
    //             // &message.Room.Name,
    //         )
    //     },
    // )
    // if queryset.IsErr() {
    //     fmt.Println(queryset.Errors)
    // } else {
    //     for _, message := range queryset.Response {
    //         fmt.Println(message.ID)
    //         fmt.Println(message.Message)
    //     }
    // }

    // response, err := models.Query(
    //     func(rows *sql.Rows) (models.Model, error) {
    //         message := new(models.Message)
    //         message.Room = *new(models.Room)
    //         return message, rows.Scan(
    //             &message.ID, &message.Message, &message.IsRead,
    //             &message.RoomID, &message.FromID, &message.ToID,
    //             &message.Room.Name,
    //         )
    //     },
    //     `SELECT 
    //         msg.ID, msg.Message, msg.IsRead, 
    //         msg.RoomID, msg.FromID, msg.ToID,
    //         Room.Name as RoomName
    //     FROM Message as msg 
    //     INNER JOIN Room 
    //     ON msg.RoomID = Room.ID`, 
    //     db,  // TODO remove
    // )
    // if err != nil {
    //     fmt.Println(err)
    // }
    // messages := models.ResponseConvert(response)
    // fmt.Println(messages[0].ID)
}
