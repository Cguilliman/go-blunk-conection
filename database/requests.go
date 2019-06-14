package database

import (
    "fmt"
    "database/sql"
)

type Message struct {
    ID      uint
    Message string
    IsRead  bool
    RoomID  int
    FromID  int
    ToID    int
}


func GetMessage(db *sql.DB) {
    rows, err := db.Query("SELECT * FROM Message")
    
    if err != nil {
        fmt.Println(err)
    }
    defer rows.Close()
    messages := make([]*Message, 0)
    for rows.Next() {
        message := new(Message)
        err := rows.Scan(
            &message.ID, &message.Message, 
            &message.IsRead, &message.RoomID, 
            &message.FromID, &message.ToID,
        )
        if err != nil {
            fmt.Println(err)
        }
        messages = append(messages, message)
    }
    if err := rows.Err(); err != nil {
        fmt.Println(err)
    }
    fmt.Println(messages[0])
}
