package models

import (
    "database/sql"
)

type Message struct {
    ID      uint
    Message string
    IsRead  bool
    RoomID  int
    FromID  int
    ToID    int
    Room    Room
}

// implement conversion in Message object
func (self *Message) Scan(rows *sql.Rows) (Model, error) {
    return self, rows.Scan(          // magic ))
        &self.ID, &self.Message, 
        &self.IsRead, &self.RoomID,
        &self.FromID, &self.ToID,
    )
}

type MessageQuerySet struct {
    Response []*Message
    Errors   []error
    Query    string
}

func (self *MessageQuerySet) IsErr() bool {
    return len(self.Errors) > 0
}

func (self *MessageQuerySet) MakeQuery(query string, scanOne func(*sql.Rows)(*Message, error)) *MessageQuerySet {
    self.Response = make([]*Message, 0) 
    messageChan := make(chan *sql.Rows)

    if scanOne == nil {
        scanOne = func(rows *sql.Rows)(*Message, error) {
            models, err := new(Message).Scan(rows)
            return models.(*Message), err
        }
    }

    go func() {
        if err := Query(query, messageChan); err != nil {
            self.Errors = append(self.Errors, err)
        }
    }()
    for rows := range messageChan {
        message, err := scanOne(rows)
        if err != nil {
            close(messageChan)
            self.Errors = append(self.Errors, err)
            return self
        }
        self.Response = append(self.Response, message)
    }

    return self
}
