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

// func ConvertInString(value interface{}) string {
//     switch t := value.(type) {
//     case int, uint:
//         return ""
//     case string:
//         return value
//     case bool:
//         if value {
//             return "1"
//         }
//         return "0"
//     }
// }

func Test(value interface{}) (string, bool) {
    switch value.(type) {
    case int: 
        return fmt.Sprint(value), true
    case uint:
        return fmt.Sprint(value), true
    case string:
        if value.(string) == "" {
            return "", false
        }
        return value.(string), true
    case bool:
        if value.(bool) {
            return "1", true
        }
        return "0", true
    default:
        return "", false
    }
}

// create person (registration)
func CreatePerson(person *models.Person) {
    var fields, values []string

    elems := reflect.ValueOf(person).Elem()
    typeOf := elems.Type()

    for i := 0; i < elems.NumField(); i++ {
        value, ok := Test(elems.Field(i).Interface())
        if ok && typeOf.Field(i).Name != "ID" {
            fields = append(fields, value)
            values = append(values, typeOf.Field(i).Name)
        }
        // fmt.Println(elems.Field(i).String())
        // fmt.Println(reflect.TypeOf(elems.Field(i).Interface()))
    }
    fmt.Println(fields)
    fmt.Println(values)
    // database := base.GetDB()
    // rows, err := database.Query(`
    //     insert into Person ()
    // `)
}

func PushToPerson(fields, values []string) {
    insertFields := strings.Join(fields, ",")
    insertValues := strings.Join(values, ",")
    database := base.GetDB()
    rows, err := database.Query(`
        insert into Person(%s)
    `)
}

// update person
func UpdatePerson() {
}
