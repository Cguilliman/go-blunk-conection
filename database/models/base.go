package models

import (
    "database/sql"
    // "fmt"
    db "github.com/Cguilliman/test_blunk_db/database/base"
)

type Model interface {
    Scan(*sql.Rows) (Model, error)
    // ResponseConvert([]Model) ([]interface{})
}

// type ModelQuerySet interface {
//     // []*Model
//     string
//     MakeQuery(string)
// }

// func (self *ModelQuerySet) MakeQuery(query string) {
//     rows, err := db.Query(query)

// }

func Query(query string, responseChan chan *sql.Rows) error {
    database := db.GetDB()
    rows, err := database.Query(query)
    if err != nil {
        close(responseChan)
        return err
    }
    defer rows.Close()
    for rows.Next() {
        responseChan <- rows
    }
    close(responseChan)
    return nil
}

// TODO: implement creation/updating
