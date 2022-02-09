package entity

import (
	"github.com/felipeagger/go-boilerplate/pkg/utils"
	"time"

	"gorm.io/gorm"
)

const dateLayout = "2006/01/02"

type User struct {
	ID        int64  `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"type:varchar(200)" json:"name"`
	Email     string `gorm:"type:varchar(100);not null;unique_index" json:"email"`
	BirthDate time.Time
	Password  string `gorm:"type:varchar(255)" json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

//NewUser create a new user
func NewUser(name, email, password, birthdate string) (*User, error) {

	date, _ := time.Parse(dateLayout, birthdate)
	birthDate := date

	hashPassword, err := utils.GenerateHashPassword(password)
	if err != nil {
		return nil, ErrGeneratePassword
	}

	usr := &User{
		ID:        utils.GeneratedID(),
		Name:      name,
		Email:     email,
		BirthDate: birthDate,
		Password:  hashPassword,
		CreatedAt: time.Now(),
	}

	err = usr.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return usr, nil
}

func (u *User) UpdatePassword(password string) error {

	hashPassword, err := utils.GenerateHashPassword(password)
	if err != nil {
		return ErrGeneratePassword
	}

	u.Password = hashPassword
	return nil
}

//Validate validate user
func (u *User) Validate() error {
	if u.Name == "" || u.Email == "" || u.Password == "" {
		return ErrInvalidEntity
	}
	return nil
}

func (u *User) ValidatePassword(password string) bool {

	return utils.CheckPasswordHash(password, u.Password)
}