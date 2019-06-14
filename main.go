package main 

import (
    "fmt"
    "database/sql"
    "github.com/gin-gonic/gin"
    "github.com/Cguilliman/test_blunk_db/database"
    _ "github.com/mattn/go-sqlite3"
)

func Init() *sql.DB {
    db, err := sql.Open("sqlite3", "./../test_blunk_db.db")
    if err != nil {
        fmt.Println(err)
    }
    return db
}

func main() {
    db := Init()
    defer db.Close()
    database.GetMessage(db)
    engine := gin.Default()
    engine.Run("0.0.0.0:8000")
}
