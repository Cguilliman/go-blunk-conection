package requests

import (
    "fmt"
    "reflect"
    "strings"
    // "strconv"
    "github.com/Cguilliman/chat/database/models"
    "github.com/Cguilliman/chat/database/base"
)

type RoomWithMessageCount struct {
    Room         *models.Room
    MessageCount int
    LastMessage  *models.Message
}

func MainUserRoomsList(userID int) {
    database := base.GetDB()
    rows, err := database.Query(fmt.Sprintf(`
        select 
            Room.ID, Room.Name, count(Message.ID) AS MessageCount
        from Room
            left join 
                RoomPerson 
            on Room.ID=RoomPerson.RoomID
            left outer join 
                Message 
            on Room.ID = Message.RoomID and Message.ToID=1
        where RoomPerson.PersonID = %d
        group by Room.ID, Room.Name
    `, userID))
    if err != nil {
        fmt.Println(err)
    }
    rooms := make([]*RoomWithMessageCount, 0)
    roomsIndx := make(map[uint]*RoomWithMessageCount, 0)
    for rows.Next() {
        responseObj := new(RoomWithMessageCount)
        responseObj.Room = new(models.Room)
        err := rows.Scan(
            &responseObj.Room.ID, &responseObj.Room.Name,
            &responseObj.MessageCount,
        )
        if err != nil {
            fmt.Println(err)
        }
        rooms = append(rooms, responseObj)
        roomsIndx[responseObj.Room.ID] = responseObj
    }

    var strIndx []string
    for _, value := range reflect.ValueOf(roomsIndx).MapKeys() {
        strIndx = append(strIndx, fmt.Sprint(value.Uint()))
    }
    filter := "("+strings.Join(strIndx, ",")+")"
    fmt.Println(filter)

    rows, err = database.Query(fmt.Sprintf(`
        select 
            Message.ID, Message.Message, max(Message.CreatedAt) as CreatedAt, 
            Message.RoomID, Message.FromID, Person.ID as PersonID, 
            Person.Username, Person.FirstName, Person.LastName
        from Message 
        inner join Person on Person.ID=Message.FromID
        where Message.RoomID in %s 
        group by Message.RoomID
    `, filter))
    for rows.Next() {
        message := new(models.Message)
        message.From = new(models.Person)
        err := rows.Scan(
            &message.ID, &message.Message, &message.CreatedAt,
            &message.RoomID, &message.FromID, &message.From.ID, 
            &message.From.Username, &message.From.FirstName, 
            &message.From.LastName,
        )
        if err != nil {
            fmt.Println(err)
        }

        roomsIndx[uint(message.RoomID)].LastMessage = message
    }

    for _, room := range rooms {
        fmt.Println(
            room.Room, 
            room.LastMessage, 
            room.LastMessage.From, 
            room.MessageCount,
        )
    }
}

func GetRooms() {
    database := base.GetDB()
    rows, _ := database.Query(`
        select 
            Room.*, Message.Message, 
            Message.IsRead, Message.CreatedAt
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
            &message.CreatedAt,
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
            fmt.Println(
                message.Message, 
                message.IsRead, 
                message.CreatedAt,
            )
        }
    }
}
