package model

import "time"

// Account struct to illustrate full object
type Account struct {
	ID        int       `bson:"_id" json:"id" `
	Name      string    `bson:"name" json:"name"`
	Cpf       string    `bson:"cpf" json:"cpf"`
	Secret    string    `bson:"secret" json:"secret"`
	Balance   float64   `bson:"balance" json:"balance"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

// NewAccount struct to illustrate post body
type NewAccount struct {
	Name    string  `bson:"name" json:"name" validate:"required"`
	Cpf     string  `bson:"cpf" json:"cpf" validate:"required"`
	Secret  string  `bson:"secret" json:"secret" validate:"required"`
	Balance float64 `bson:"balance" json:"balance" validate:"required"`
}
