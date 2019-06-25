package users

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "github.com/Cguilliman/chat/controllers/middlewares"
    "github.com/Cguilliman/chat/shared"
)

func Registration(c *gin.Context) {
    validator := NewRegistrationValidator()
    if err := validator.Bind(c); err != nil {
        c.JSON(http.StatusUnprocessableEntity, gin.H{
            "errors": err,
        })
        return
    }
    id, err := validator.Register()
    if err != nil {
        c.JSON(http.StatusUnprocessableEntity, gin.H{
            "errors": err,
        })
        return
    }
    c.JSON(http.StatusCreated, gin.H{"userID": id})
}


func Login(c *gin.Context) {
    validator := NewLoginValidator()
    if err := validator.Bind(c); err != nil {
        c.JSON(http.StatusUnprocessableEntity, gin.H{
            "errors": err,
        })
        return
    }
    person, err := validator.Login()
    if err != nil {
        // TODO: edit to custom error
        c.JSON(http.StatusUnprocessableEntity, gin.H{
            "errors": err,
        })
        return
    }
    middlewares.UpdateContext(c, person.ID)
    c.JSON(http.StatusOK, gin.H{
        "person": person, 
        "token": shared.GenToken(person.ID),
    })
}
