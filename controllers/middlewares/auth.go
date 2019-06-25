package middlewares

import (
    "strings"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
    "github.com/dgrijalva/jwt-go/request"

    "github.com/Cguilliman/chat/database/requests"
    "github.com/Cguilliman/chat/database/models"
    "github.com/Cguilliman/chat/settings"
)

func stripBearerPrefixFromTokenString(token string) (string, error) {
    if len(token) > 5 && strings.ToUpper(token[0:6]) == "TOKEN " {
        return token[6:], nil
    }
    return token, nil
}

var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
    request.HeaderExtractor{"Authorization"},
    stripBearerPrefixFromTokenString,
}

var MyAuth2Extractor = &request.MultiExtractor{
    AuthorizationHeaderExtractor,
    request.ArgumentExtractor{"access_token"},
}

func UpdateContext(c *gin.Context, currentUser uint) {
    var userModel models.Person
    if currentUser != 0 {
        userModel, _ = requests.GetPerson(currentUser)
    }
    c.Set("user_id", currentUser)
    c.Set("user", userModel)
}

func AuthMiddleware(auto401 bool) gin.HandlerFunc {
    return func(c *gin.Context) {
        UpdateContext(c, 0)
        token, err := request.ParseFromRequest(
            c.Request,
            MyAuth2Extractor,
            func(token *jwt.Token) (interface{}, error) {
                b := ([]byte(settings.Secret))
                return b, nil
            },
        )

        if err != nil {
            if auto401 {
                c.AbortWithError(http.StatusUnprocessableEntity, err)
            }
        }
        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            UserID := uint(claims["id"].(float64))
            UpdateContext(c, UserID)
        }
    }
}