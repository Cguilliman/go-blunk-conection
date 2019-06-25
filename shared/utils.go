package shared

import (
    "time"
    "github.com/dgrijalva/jwt-go"
    "github.com/Cguilliman/chat/settings"
)

func GenToken(id uint) string {
    jwt_token := jwt.New(jwt.GetSigningMethod("HS256"))
    // Set some claims
    jwt_token.Claims = jwt.MapClaims{
        "id":  id,
        "exp": time.Now().Add(time.Hour * 24).Unix(),
    }
    // Sign and get the complete encoded token as a string
    token, _ := jwt_token.SignedString([]byte(settings.Secret))
    return token
}
