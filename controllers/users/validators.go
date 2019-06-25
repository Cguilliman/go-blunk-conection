package users

import (
    // "fmt"
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "golang.org/x/crypto/bcrypt"

    "github.com/Cguilliman/chat/database/models"
    "github.com/Cguilliman/chat/database/requests"
)

type UserRegistrationValidator struct {
    Person struct {
        Username  string `form:"username"  json:"username"  binding:"required"` 
        FirstName string `form:"firstname" json:"firstname" binding:"required"` 
        LastName  string `form:"lastname"  json:"lastname"  binding:"required"` 
        Password  string `form:"password"  json:"password"  binding:"required"` 
    } `json:"person"`
    personModel models.Person `json:"-"`
}

// TODO: remove to common package
func ConvertPassword(password []byte) (string, error) {
    hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
    return string(hash), err
}

func Bind(c *gin.Context, obj interface{}) error {
    b := binding.Default(c.Request.Method, c.ContentType())
    return c.ShouldBindWith(obj, b)
}

type ValidationError struct {
    s string
}

func (self *ValidationError) Error() string {
    return self.s
}
// end remove

func (self *UserRegistrationValidator) Bind(c *gin.Context) error {
    var err error
    if err = Bind(c, self); err != nil {
        return err
    }
    if requests.CheckPerson(self.Person.Username) {
        return &ValidationError{"Sorry. User with current username already exists."}
    }  

    self.personModel.Username = self.Person.Username
    self.personModel.FirstName = self.Person.FirstName
    self.personModel.LastName = self.Person.LastName
    self.personModel.Password, err = ConvertPassword([]byte(self.Person.Password))
    return err
}

func (self *UserRegistrationValidator) Register() (int64, error) {
    return requests.CreatePerson(&self.personModel)
}

func NewRegistrationValidator() UserRegistrationValidator {
    return UserRegistrationValidator{}
}

type LoginValidator struct {
    Person struct {
        Username string `form:"username" json:"username" binding:"required"`
        Password string `form:"password" json:"password" binding:"required"`
    }
    // personModel models.Person
}

func (self *LoginValidator) Bind(c *gin.Context) error {
    if err := Bind(c, self); err != nil {
        return err
    }
    return nil
}

func CheckPassword(modelPsw, password []byte) error {
    return bcrypt.CompareHashAndPassword(modelPsw, password)
}

func (self *LoginValidator) Login() (models.Person, error) {
    person, modelPsw, err := requests.Login(self.Person.Username)
    if err != nil {
        return person, err
    }
    return person, CheckPassword(
        []byte(modelPsw), 
        []byte(self.Person.Password),
    )
}

func NewLoginValidator() LoginValidator {
    validator := LoginValidator{}
    return validator
}
