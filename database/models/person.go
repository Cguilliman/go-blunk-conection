package models

import "database/sql"

type Person struct {
    ID           uint
    Username     string
    FirstName    string
    LastName     string
    Password     string
    SendMessages []*Message
    GetMessages  []*Message
    Rooms        []*Room
}

func (self Person) Scan(rows *sql.Rows) (Model, error) {
    return self, rows.Scan(
        &self.ID, &self.Username, 
        &self.FirstName, &self.LastName, 
        &self.Password,
    )
}

type PersonQuerySet struct {
    Persons []*Person
    Errors  []error
    Qeury   string
}

func (self *PersonQuerySet) IsErr() bool {
    return len(self.Errors) > 0
}

func (self *PersonQuerySet) WriteOne(obj interface{}) {
    self.Persons = append(self.Persons, obj.(*Person))
}

func (self *PersonQuerySet) AddErr(err error) {
    self.Errors = append(self.Errors, err)
}

func (self *PersonQuerySet) NewQuery(query string, scanOne func(*sql.Rows)(interface{}, error)) *PersonQuerySet {
    if scanOne == nil {
        scanOne = func(rows *sql.Rows)(interface{}, error) {
            return new(Person).Scan(rows)
        }
    }
    res := Query(query, self, scanOne)
    self = res.(*PersonQuerySet)
    return self
}

