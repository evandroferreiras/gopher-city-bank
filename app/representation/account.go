package representation

import (
	"time"

	"github.com/evandroferreiras/gopher-city-bank/app/model"
)

// NewAccount struct to illustrate post body
type NewAccount struct {
	Name    string  `bson:"name" json:"name" validate:"required"`
	Cpf     string  `bson:"cpf" json:"cpf" validate:"required"`
	Secret  string  `bson:"secret" json:"secret" validate:"required"`
	Balance float64 `bson:"balance" json:"balance" validate:"gte=0"`
}

// AccountCreated struct to illustrated created account
type AccountCreated struct {
	ID        int       `bson:"_id" json:"id" `
	Name      string    `bson:"name" json:"name"`
	Cpf       string    `bson:"cpf" json:"cpf"`
	Balance   float64   `bson:"balance" json:"balance"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

// ToModel converts NewAccount representation struct to Account Model
func (n *NewAccount) ToModel() model.Account {
	return model.Account{
		ID:        0,
		Name:      n.Name,
		Cpf:       n.Cpf,
		Secret:    n.Secret,
		Balance:   n.Balance,
		CreatedAt: time.Time{},
	}
}

// ModelToAccountCreated converts Account model to AccountCreated representation
func ModelToAccountCreated(a *model.Account) *AccountCreated {
	return &AccountCreated{
		ID:        a.ID,
		Name:      a.Name,
		Cpf:       a.Cpf,
		Balance:   a.Balance,
		CreatedAt: a.CreatedAt,
	}
}
