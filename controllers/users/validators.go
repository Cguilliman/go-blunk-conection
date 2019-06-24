package users

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "golang.org/x/crypto/bcrypt"

    "github.com/Cguilliman/chat/database/models"
    "github.com/Cguilliman/chat/database/requests"
)

type UserRegistrationValidator struct {
    Person struct {
        Username  string `from:"username"  json:"username"  binding:"exists,max=255"` 
        FirstName string `from:"firstname" json:"firstname" binding:"exists,max=255"` 
        LastName  string `from:"lastname"  json:"lastname"  binding:"exists,max=255"` 
        Password  string `from:"password"  json:"password"  binding:"exists,max=255"` 
    } `json:"person"`
    personModel models.Person `json:"-"`
}

// TODO: remove to common package
func ConvertPassword(password []byte) (string, error) {
    hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
    // bcrypt.CompareHashAndPassword(model.Password string, password []byte) - return error
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
        fmt.Println(err)
        return err
    }
    if requests.CheckPerson(self.Person.Username) {
        return &ValidationError{"Sorry. User with current username already exists."}
    }  
    fmt.Println(self.Person)

    self.personModel.Username = self.Person.Username
    self.personModel.FirstName = self.Person.FirstName
    self.personModel.LastName = self.Person.LastName
    self.personModel.Password, err = ConvertPassword([]byte(self.Person.Password))
    return err
}

func (self *UserRegistrationValidator) Register() (int64, error) {
    fmt.Println("SCHEMA =============== ",self.personModel)
    return requests.CreatePerson(&self.personModel)
}

func NewRegistrationValidator() UserRegistrationValidator {
    return UserRegistrationValidator{}
}
