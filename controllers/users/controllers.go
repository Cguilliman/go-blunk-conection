package users

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "github.com/Cguilliman/chat/controllers/middlewares"
    "github.com/Cguilliman/chat/serializers"
    "github.com/Cguilliman/chat/database/models"
)

func Registration(c *gin.Context) {
    validator := NewRegistrationValidator()
    if err := validator.Bind(c); err != nil {
        c.JSON(http.StatusUnprocessableEntity, gin.H{
            "errors": err,
        })
        return
    }
    _, err := validator.Register()
    if err != nil {
        c.JSON(http.StatusUnprocessableEntity, gin.H{
            "errors": err,
        })
        return
    }
    c.JSON(http.StatusCreated, gin.H{"Message": "Success. Approve credentials."})
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
    serializer := serializers.PersonSerializer{c}
    c.JSON(http.StatusOK, gin.H{
        "response": serializer.Response(person, true),
    })
}

func Receive(c *gin.Context) {
    person := c.MustGet("user").(models.Person)
    serializer := serializers.PersonSerializer{c}
    c.JSON(http.StatusOK, gin.H{
        "response": serializer.Response(person, false),
    })
}
