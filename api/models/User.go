package models

import (
	"errors"
	"html"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	sto "github.com/igorariza/Dockerized-Golang_API-MySql-React.js/internal/storage"
)

type User struct {
	gorm.Model
	Name      string `gorm:"size:255" json:"name"`
	Email     string `gorm:"size:255" json:"email"`
	Password  string `gorm:"size:100;" json:"password"`
	Address   string `gorm:"size:100;" json:"address"`
	Birthdate string `gorm:"size:100;" json:"birthdate"`
	City      string `gorm:"size:100;" json:"city"`
}

type UserResponse struct {
	gorm.Model
	Name      string `gorm:"size:255" json:"name"`
	Email     string `gorm:"size:255" json:"email"`
	Address   string `gorm:"size:100;" json:"address"`
	Birthdate string `gorm:"size:100;" json:"birthdate"`
	City      string `gorm:"size:100;" json:"city"`
}

type LoginUser struct {
	ID       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginGateway interface {
	LoginUser(db *gorm.DB) (*User, error)
}

type LoginInDB struct{}

func (c *User) LoginUser(db *gorm.DB) (*LoginUser, error) {
	return c.loginUserDB(db)
}

func (u *User) Prepare() {
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Password = html.EscapeString(strings.TrimSpace(u.Password))
	u.Address = html.EscapeString(strings.TrimSpace(u.Address))
	u.Birthdate = html.EscapeString(strings.TrimSpace(u.Birthdate))
	u.City = html.EscapeString(strings.TrimSpace(u.City))
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if u.Birthdate == "" {
			return errors.New("Required Birthdate")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Address == "" {
			return errors.New("Required Address")
		}
		if u.City == "" {
			return errors.New("Required City")
		}

		return nil

	default:
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if u.Birthdate == "" {
			return errors.New("Required Birthdate")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Address == "" {
			return errors.New("Required Address")
		}
		if u.City == "" {
			return errors.New("Required City")
		}
		return nil
	}
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	u.Password, _ = sto.EncryptPassword(u.Password)
	var err error

	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return &User{
		Name:      u.Name,
		Email:     u.Email,
		Address:   u.Address,
		Birthdate: u.Birthdate,
		City:      u.City,
	}, nil
}

func (u *User) FindAllusers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}
func (u *User) List(db *gorm.DB, pagination Pagination) (*Pagination, error) {
	users := []User{}
	db.Scopes(paginate(users, &pagination, db)).Find(&users)
	pagination.Rows = users

	return &pagination, nil
}

func (u *User) loginUserDB(db *gorm.DB) (*LoginUser, error) {
	var user LoginUser
	err := db.Table("users").Where("email = ?", u.Email).First(&user).Error
	if err != nil {
		log.Printf("cannot fetch user email loginUserDB")
		return nil, err
	}

	passwordBytes := []byte(u.Password)
	passwordBD := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		log.Printf("error en la contraseña ")
		return nil, err
	}
	return &LoginUser{
		Email:    user.Email,
		Password: "",
	}, nil
}

// ChequeoYaExisteUsuario recibe email y chequea en la BD si existe
func (u *User) ChequeoYaExisteUsuario(db *gorm.DB, email string) (bool, error) {
	var user User
	var err error
	err = db.Table("users").Where("email = ?", u.Email).First(&user).Error
	if err != nil {
		log.Printf("cannot fetch user email ChequeoYaExisteUsuario")
		return false, err
	}
	if user.Email != "" {
		return false, nil
	}

	return true, err
}
