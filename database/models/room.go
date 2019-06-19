package models

import "database/sql"


// implement `Model`
type Room struct {
    ID       uint
    Name     string
    Messages []Message
}

func (self *Room) Scan(rows *sql.Rows) (Model, error) {
    return self, rows.Scan(&self.ID, &self.Name)
}

// implement `ModelQuerySet`
type RoomQuerySet struct {
    Rooms  []*Room
    Errors []error
    Query  string
}

func (self *RoomQuerySet) IsErr() bool {
    return len(self.Errors) > 0
}

func (self *RoomQuerySet) WriteOne(obj interface{}) {
    self.Rooms = append(self.Rooms, obj.(*Room))
}

func (self *RoomQuerySet) AddErr(err error) {
    self.Errors = append(self.Errors, err)
}

func (self *RoomQuerySet) NewQuery(query string, scanOne func(*sql.Rows)(interface{}, error)) *RoomQuerySet {
    if scanOne == nil {
        scanOne = func(rows *sql.Rows)(interface{}, error) {
            return new(Room).Scan(rows)
        }
    }
    res := Query(query, self, scanOne)
    self = res.(*RoomQuerySet)
    return self
}
