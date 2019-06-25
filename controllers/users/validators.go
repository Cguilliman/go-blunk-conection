package users

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "golang.org/x/crypto/bcrypt"

    "github.com/Cguilliman/chat/database/models"
    "github.com/Cguilliman/chat/database/requests"
)

type UserRegistrationValidator struct {
    Person struct {
        Username  string `form:"username"  json:"username"  binding:"exists"` 
        FirstName string `form:"firstname" json:"firstname" binding:"exists"` 
        LastName  string `form:"lastname"  json:"lastname"  binding:"exists"` 
        Password  string `form:"password"  json:"password"  binding:"exists"` 
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
