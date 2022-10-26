package model

import (
	"github.com/google/uuid"
)

// User defines domain model and its json and db representations
type User struct {
	UID         uuid.UUID `db:"uid" json:"uid"`
	Email       string    `db:"email" json:"email"`
	Password    string    `db:"password" json:"-"`
	Name        string    `db:"name" json:"name"`
	ImageURL    string    `db:"image_url" json:"imageUrl"`
	Website     string    `db:"website" json:"website"`
	FirstName   string    `db:"firstname" json:"firstname"`
	LastName    string    `db:"lastname" json:"lastname"`
	Telegram    string    `db:"telegram" json:"telegram"`
	PhoneNumber string    `db:"phonenumber" json:"phonenumber"`
	Country     string    `db:"country" json:"country"`
	City        string    `db:"city" json:"city"`
	Location    string    `db:"location" json:"location"`
	Lang        string    `db:"lang" json:"lang"`
	Verified    bool      `db:"verified" json:"verified"`
}
