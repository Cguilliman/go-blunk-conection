package requests

import (
    "fmt"
    // "database/sql"
    "github.com/Cguilliman/chat/database/models"
    "github.com/Cguilliman/chat/database/base"
)

type Pagination struct {
    PerPage     int
    Amount      int
    CurrentPage int
}

type PaginatedMessages struct {
    Messages   []*models.Message
    Pagination *Pagination
}

func GetMessages(page, paginateBy int) {
    database := base.GetDB()
    // Get general messages amount
    rows, err := database.Query(`select count(*) as MessageCount from Message;`)
    if err != nil {
        fmt.Println(err)
    }
    rows.Next()
    var messageCount int
    err = rows.Scan(&messageCount)
    if err != nil {
        fmt.Println(err)
    }

    // Get messages with pagination
    rows, err = database.Query(fmt.Sprintf(`
        select 
            Message.ID, Message.Message, 
            Message.CreatedAt, Person.Username,
            Person.LastName, Person.FirstName
        from Message
            inner join Person on Message.FromID=Person.ID
        limit %d, %d
    `, paginateBy*page-paginateBy, paginateBy))
    if err != nil {
        fmt.Println(err)
    }
    response := new(PaginatedMessages)
    for rows.Next() {
        message := new(models.Message)
        message.From = new(models.Person)
        err := rows.Scan(
            &message.ID, &message.Message,
            &message.CreatedAt, &message.From.Username,
            &message.From.LastName, &message.From.FirstName,
        )
        if err != nil {
            fmt.Println(err)
        }
        response.Messages = append(response.Messages, message)
    }
    response.Pagination = &Pagination{
        PerPage:     paginateBy,
        Amount:      messageCount,
        CurrentPage: page,
    }
    fmt.Println(response.Pagination)
    for _, message := range response.Messages {
        fmt.Println(message)
    }
}

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
