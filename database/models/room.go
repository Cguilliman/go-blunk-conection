package models

import "database/sql"


type Room struct {
    ID   uint
    Name string
}

func (self *Room) Scan(rows *sql.Rows) (Model, error) {
    return self, rows.Scan(&self.ID, &self.Name)
}
