package models

import (
    "database/sql"
    // "fmt"
)

// implement `Model`
type Message struct {
    ID      uint
    Message string
    IsRead  bool
    RoomID  int
    FromID  int
    ToID    int
    Room    *Room
}

// implement conversion in Message object
func (self *Message) Scan(rows *sql.Rows) (interface{}, error) {
    err := rows.Scan(          // magic ))
        &self.ID, &self.Message, 
        &self.IsRead, &self.RoomID,
        &self.FromID, &self.ToID,
    )
    return self, err
}

// implement `ModelQuerySet`
type MessageQuerySet struct {
    Response []*Message
    Errors   []error
    Query    string
}

func (self *MessageQuerySet) IsErr() bool {
    return len(self.Errors) > 0
}

func (self *MessageQuerySet) WriteOne(obj interface{}) {
    self.Response = append(self.Response, obj.(*Message))
}

func (self *MessageQuerySet) AddErr(err error) {
    self.Errors = append(self.Errors, err)
}

func (self *MessageQuerySet) NewQuery(query string, scanOne func(*sql.Rows)(interface{}, error)) *MessageQuerySet {
    if scanOne == nil {
        scanOne = func(rows *sql.Rows)(interface{}, error) {
            return new(Message).Scan(rows)
        }
    }
    res := Query(query, self, scanOne)
    self = res.(*MessageQuerySet)
    return self
}
