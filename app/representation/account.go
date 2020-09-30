package representation

import (
	"time"

	"github.com/evandroferreiras/gopher-city-bank/app/model"
)

// NewAccountBody struct to illustrate post body
type NewAccountBody struct {
	Name    string  `bson:"name" json:"name" validate:"required"`
	Cpf     string  `bson:"cpf" json:"cpf" validate:"required"`
	Secret  string  `bson:"secret" json:"secret" validate:"required"`
	Balance float64 `bson:"balance" json:"balance" validate:"gte=0"`
}

// AccountResponse struct to illustrated account response
type AccountResponse struct {
	ID        int       `bson:"_id" json:"id" `
	Name      string    `bson:"name" json:"name"`
	Cpf       string    `bson:"cpf" json:"cpf"`
	Balance   float64   `bson:"balance" json:"balance"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

// AccountsList struct to illustrate list of accounts
type AccountsList struct {
	Accounts []AccountResponse `bson:"accounts" json:"accounts"`
	Count    int               `bson:"count" json:"count"`
}

// ToModel converts NewAccountBody representation struct to Account Model
func (n *NewAccountBody) ToModel() model.Account {
	return model.Account{
		ID:        0,
		Name:      n.Name,
		Cpf:       n.Cpf,
		Secret:    n.Secret,
		Balance:   n.Balance,
		CreatedAt: time.Time{},
	}
}

// ModelToAccountCreated converts Account model to AccountResponse representation
func ModelToAccountCreated(a *model.Account) *AccountResponse {
	return &AccountResponse{
		ID:        a.ID,
		Name:      a.Name,
		Cpf:       a.Cpf,
		Balance:   a.Balance,
		CreatedAt: a.CreatedAt,
	}
}
