package users

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"

    // "github.com/Cguilliman/post-it-note/common"
    // "github.com/Cguilliman/post-it-note/controllers/middlewares"
    // "github.com/Cguilliman/post-it-note/models"
    // "github.com/Cguilliman/post-it-note/serializers"
)

func Registration(c *gin.Context) {
    validator := NewRegistrationValidator()
    if err := validator.Bind(c); err != nil {
        fmt.Println("error:::---", err)
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
