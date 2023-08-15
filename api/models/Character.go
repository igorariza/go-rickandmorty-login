package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type Character struct {
	gorm.Model
	Name string `gorm:"size:255" json:"name"`
	Status  string `gorm:"size:255" json:"status"`
	Species  string `gorm:"size:100;" json:"species"`
	Gender    string `gorm:"size:100;" json:"gender"`
	Image      string `gorm:"size:100;" json:"image"`
	Created      string `gorm:"size:100;" json:"created"`
}


func (u *Character) Prepare() {
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Status = html.EscapeString(strings.TrimSpace(u.Status))
	u.Species = html.EscapeString(strings.TrimSpace(u.Species))
	u.Gender = html.EscapeString(strings.TrimSpace(u.Gender))
	u.Image = html.EscapeString(strings.TrimSpace(u.Image))
	u.Created = html.EscapeString(strings.TrimSpace(u.Created))
}

func (u *Character) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Name == "" {
			return errors.New("Required Name")
		}
		
		return nil

	default:
		if u.Name == "" {
			return errors.New("Required Name")
		}
		return nil
	}
}

func (u *Character) SaveCharacter(db *gorm.DB) (*Character, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Character{}, err
	}
	return u, nil
}

func (u *Character) FindAllcharacters(db *gorm.DB) (*[]Character, error) {
	var err error
	characters := []Character{}
	err = db.Debug().Model(&Character{}).Find(&characters).Error
	if err != nil {
		return &[]Character{}, err
	}
	return &characters, err
}

func (u *Character) FindCharacterByID(db *gorm.DB, uid uint32) (*Character, error) {
	var err error
	err = db.Debug().Model(Character{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Character{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Character{}, errors.New("User Not Found")
	}
	return u, err
}
func (u *Character) List(db *gorm.DB, pagination Pagination) (*Pagination, error) {
	users := []Character{}
	db.Scopes(paginate(users, &pagination, db)).Find(&users)
	pagination.Rows = users

	return &pagination, nil
}
