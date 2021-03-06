package model

import (
	"time"
)

// Account struct to illustrate database object
type Account struct {
	ID        string    `bson:"_id" json:"id" gorm:"type:varchar(36);primaryKey"`
	Name      string    `bson:"name" json:"name"`
	Cpf       string    `bson:"cpf" json:"cpf" gorm:"type:varchar(11);unique"`
	Secret    string    `bson:"secret" json:"secret"`
	Balance   float64   `bson:"balance" json:"balance"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

// EmptyAccount is a empty account model used to comparing
var EmptyAccount = Account{}
