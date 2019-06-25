package users

import (
    // "fmt"
    "github.com/gin-gonic/gin"

    "github.com/Cguilliman/chat/shared"
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

// end remove

func (self *UserRegistrationValidator) Bind(c *gin.Context) error {
    var err error
    if err = shared.Bind(c, self); err != nil {
        return err
    }
    if requests.CheckPerson(self.Person.Username) {
        return &shared.ValidationError{
            "Sorry. User with current username already exists.",
        }
    }  

    self.personModel.Username = self.Person.Username
    self.personModel.FirstName = self.Person.FirstName
    self.personModel.LastName = self.Person.LastName
    self.personModel.Password, err = shared.ConvertPassword([]byte(self.Person.Password))
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
}

func (self *LoginValidator) Bind(c *gin.Context) error {
    if err := shared.Bind(c, self); err != nil {
        return err
    }
    return nil
}

func (self *LoginValidator) Login() (models.Person, error) {
    person, modelPsw, err := requests.Login(self.Person.Username)
    if err != nil {
        return person, err
    }
    return person, shared.CheckPassword(
        []byte(modelPsw), 
        []byte(self.Person.Password),
    )
}

func NewLoginValidator() LoginValidator {
    validator := LoginValidator{}
    return validator
}
