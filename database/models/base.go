package models

import (
    "database/sql"
)

type Model interface {
    Scan(*sql.Rows) (Model, error)
    // ResponseConvert([]Model) ([]interface{})
}

func Query(scanOne func(*sql.Rows)(Model, error), query string, db *sql.DB) ([]Model, error) {
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    response_set := make([]Model, 0)
    for rows.Next() {
        obj, err := scanOne(rows)
        if err != nil {
            return nil, err
        }
        response_set = append(response_set, obj)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }
    return response_set, nil
}

// TODO: implement creation/updating
