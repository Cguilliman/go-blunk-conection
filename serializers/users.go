package serializers

import (
    "github.com/gin-gonic/gin"
    "github.com/Cguilliman/chat/shared"
    "github.com/Cguilliman/chat/database/models"
)

type PersonSerializer struct {
    C *gin.Context
}

type PersonStruct struct {
    ID        int    `json:"id"`
    Username  string `json:"username"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
}

type PersonResponse struct {
    Person PersonStruct `json:"person"`
    Token  string       `json:"token"`
}

func (self *PersonSerializer) Response(person models.Person, isToken bool) PersonResponse {
    var token string
    if isToken {
        token = shared.GenToken(person.ID)
    }
    return PersonResponse{
        Person: PersonStruct{
            ID:        int(person.ID),
            Username:  person.Username,
            FirstName: person.FirstName,
            LastName:  person.LastName,
        },
        Token: token,
    }
}
