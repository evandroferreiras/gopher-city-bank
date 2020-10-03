package transfer

import "github.com/evandroferreiras/gopher-city-bank/app/model"

const accountOriginID = "800"
const accountDestinationID = "189"
const amount float64 = 500

var accountOriginReturned = model.Account{
	ID:      accountOriginID,
	Name:    "Bruce Wayne",
	Cpf:     "01595995555",
	Balance: amount,
}

var accountDestinationReturned = model.Account{
	ID:      accountDestinationID,
	Name:    "Damian Wayne",
	Cpf:     "01595995666",
	Balance: 0.2,
}
