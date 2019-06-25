package shared

import (
    "time"

    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "github.com/dgrijalva/jwt-go"
    "golang.org/x/crypto/bcrypt"

    "github.com/Cguilliman/chat/settings"
)

type ValidationError struct {
    Message string
}

func (self *ValidationError) Error() string {
    return self.Message
}

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

func ConvertPassword(password []byte) (string, error) {
    hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
    return string(hash), err
}

func CheckPassword(modelPsw, password []byte) error {
    return bcrypt.CompareHashAndPassword(modelPsw, password)
}

func Bind(c *gin.Context, obj interface{}) error {
    b := binding.Default(c.Request.Method, c.ContentType())
    return c.ShouldBindWith(obj, b)
}
