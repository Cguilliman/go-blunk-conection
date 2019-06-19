package models

import (
    "database/sql"
    db "github.com/Cguilliman/test_blunk_db/database/base"
)

type Model interface {
    Scan(*sql.Rows) (Model, error)
}

type ModelQuerySet interface {
    AddErr(error)
    WriteOne(interface{})
    IsErr() bool
}

func Query(query string, queryset ModelQuerySet, scan func(*sql.Rows) (interface{}, error)) ModelQuerySet {
    responseChan := make(chan *sql.Rows)
    database := db.GetDB()
    rows, err := database.Query(query)
    if err != nil {
        queryset.AddErr(err)
        return queryset
    }
    defer rows.Close()

    for rows.Next() {
        obj, err := scan(rows)
        if err != nil {
            close(responseChan)
            queryset.AddErr(err)
            break
        }
        queryset.WriteOne(obj)
    }
    return queryset
}

// TODO: implement creation/updating
