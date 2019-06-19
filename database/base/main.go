package base

import (
    "fmt"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() *sql.DB {
    db, err := sql.Open("sqlite3", "./../test_blunk_db.db")
    if err != nil {
        fmt.Println(err)
    }
    DB = db
    return DB
}

func GetDB() *sql.DB {
    return DB
}
