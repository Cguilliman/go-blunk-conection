package requests

import (
    "fmt"
    "reflect"
    "strings"
    "github.com/Cguilliman/chat/database/models"
    "github.com/Cguilliman/chat/database/base"
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

// get current user
func CheckPerson(username string) bool {
    var id int
    database := base.GetDB()
    err := database.QueryRow(fmt.Sprintf(`
        select 
            Person.ID
        from Person 
        where Person.Username="%s"
    `, username)).Scan(&id)
    if err != nil {
        return false
    }
    return true
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
        return fmt.Sprintf("'%s'", value.(string)), true
    case bool:
        if value.(bool) {
            return "1", true
        }
        return "0", true
    default:
        return "", false
    }
}

func ConvertPersonToPush(person *models.Person) ([]string, []string) {
    var fields, values []string

    elems := reflect.ValueOf(person).Elem()
    typeOf := elems.Type()

    for i := 0; i < elems.NumField(); i++ {
        value, ok := Test(elems.Field(i).Interface())
        if ok && typeOf.Field(i).Name != "ID" {
            values = append(values, value)
            fields = append(fields, typeOf.Field(i).Name)
        }
    }
    return fields, values
}

// create person (registration)
func CreatePerson(person *models.Person) (int64, error) {
    fields, values := ConvertPersonToPush(person)
    var id int64

    insertFields := strings.Join(fields, ",")
    insertValues := strings.Join(values, ",")
    
    database := base.GetDB()
    res, err := database.Exec(fmt.Sprintf(`
        insert into Person(%s) 
        values (%s)
    `, insertFields, insertValues))
    if err != nil{
        return id, err
    } else {
        id, err = res.LastInsertId()        
        if err != nil {
            return id, err
        }
    }
    return id, nil
}


// update person
func UpdatePerson(id int, person *models.Person) {
    fields, values := ConvertPersonToPush(person)
    var set []string
    for i := 0; i < len(fields); i++ {
        set = append(set, fmt.Sprintf("%s = %s", fields[i], values[i]))
    }
    setValue := strings.Join(set, ",")

    database := base.GetDB()
    _, err := database.Exec(fmt.Sprintf(`
        update Person 
        set %s
        where Person.ID = %d
    `, setValue, id))
    if err != nil {
        fmt.Println(err)
    }
}
