package models

import "database/sql"

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

func ScanMessage(rows *sql.Rows) (Model, error) {
    return new(Message).Scan(rows)
}

func Empty() []*Message {
    return make([]Message, 0)
}

func ConvertOne(obj Model) *Message {
    return obj.(*Message)
}

func ResponseConvert(response []Model) []*interface{} {
    var messages []*Message
    for _, obj := range response {
        messages = append(messages, obj.(*Message))
    }
    return messages
}
