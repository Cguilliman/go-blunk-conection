package requests

import (
    "fmt"
    "reflect"
    "github.com/Cguilliman/test_blunk_db/database/models"
    "github.com/Cguilliman/test_blunk_db/database/base"
)

// get current user
func GetPerson(id int) {
    database := base.GetDB()
    person := new(models.Person)
    err := database.QueryRow(fmt.Sprintf(`
        select 
            Person.ID, Person.Username, 
            Person.FirstName, Person.LastName 
        from Person 
        where Person.ID=%d
    `, id)).Scan(
        &person.ID, &person.Username, 
        &person.LastName, &person.FirstName,
    )
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(person)
}

// create person (registration)
func CreatePerson(person *models.Person) {
    elems := reflect.ValueOf(person).Elem()
    typeOf := elems.Type()

    for i := 0; i < elems.NumField(); i++ {
        fmt.Println(typeOf.Field(i).Name)
        fmt.Println(elems.Field(i))
    }
    // database := base.GetDB()
    // rows, err := database.Query(`
    //     insert into Person ()
    // `)
}

// update person
func UpdatePerson() {
}
