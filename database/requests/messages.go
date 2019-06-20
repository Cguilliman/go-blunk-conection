package requests

import (
    "fmt"
    "database/sql"
    "github.com/Cguilliman/test_blunk_db/database/models"
    "github.com/Cguilliman/test_blunk_db/database/base"
)

// func GetMessage(page, paginateBy int) {
//     database := base.GetDB()
//     rows, err := database.Query(`
//     `)
// }

// func GetMessage() {
//     queryset := new(models.MessageQuerySet).NewQuery(
//         `
//         SELECT 
//             msg.*, Room.Name as RoomName
//         FROM Message as msg
//         INNER JOIN Room
//         ON msg.RoomID = Room.ID`, 
//         func(rows *sql.Rows) (interface{}, error) {
//             message := new(models.Message)
//             message.Room = new(models.Room)
//             return message, rows.Scan(
//                 &message.ID, &message.Message, &message.IsRead,
//                 &message.RoomID, &message.FromID, &message.ToID,
//                 &message.Room.Name,
//             )
//         },
//     )
//     for _, message := range queryset.Response {
//         fmt.Println(message.Message)
//     }
// }
