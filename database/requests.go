package database

import (
    "fmt"
    "github.com/Cguilliman/test_blunk_db/database/models"
    "github.com/Cguilliman/test_blunk_db/database/base"
    "database/sql"
)

func GetMessage() {
    queryset := new(models.MessageQuerySet).NewQuery(
        `
        SELECT 
            msg.*, Room.Name as RoomName
        FROM Message as msg
        INNER JOIN Room
        ON msg.RoomID = Room.ID`, 
        func(rows *sql.Rows) (interface{}, error) {
            message := new(models.Message)
            message.Room = new(models.Room)
            return message, rows.Scan(
                &message.ID, &message.Message, &message.IsRead,
                &message.RoomID, &message.FromID, &message.ToID,
                &message.Room.Name,
            )
        },
    )
    for _, message := range queryset.Response {
        fmt.Println(message.Message)
    }
}

func GetRooms() {
    database := base.GetDB()
    rows, _ := database.Query(`
        select 
            Room.*, Message.Message, 
            Message.IsRead 
        from 
            Room left join Message 
        on 
            Message.RoomID = Room.ID
    `)
    defer rows.Close()
    rooms := make([]*models.Room, 0)
    roomsIndx := make(map[uint]*models.Room, 0)
    for rows.Next() {
        room := new(models.Room)
        message := new(models.Message)
        err := rows.Scan(
            &room.ID, &room.Name,
            &message.Message, &message.IsRead,
        )
        if err != nil {
            fmt.Println(err)
        }

        if _room, ok := roomsIndx[room.ID]; ok {
            room = _room
        } else {
            rooms = append(rooms, room)
            roomsIndx[room.ID] = room
        }

        room.Messages = append(room.Messages, message)
    }
    for _, room := range rooms {
        fmt.Println("-------")
        fmt.Println(room.Name)
        for _, message := range room.Messages {
            fmt.Println(message.Message, message.IsRead)
        }
    }
    // fmt.Println(roomsIndx)
}
