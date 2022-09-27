package entities

import "time"

type User struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
