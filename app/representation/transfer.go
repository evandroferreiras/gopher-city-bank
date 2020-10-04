package representation

import (
	"time"

	"github.com/evandroferreiras/gopher-city-bank/app/model"
)

// TransferBody struct to illustrate post transfer
type TransferBody struct {
	AccountDestinationID string  `bson:"account_destination_id" json:"account_destination_id" validate:"required"`
	Amount               float64 `bson:"amount" json:"amount" validate:"required"`
}

// TransferWithDrawResponse struct to illustrate a withdraw transfer response
type TransferWithDrawResponse struct {
	AccountDestinationID string    `bson:"account_destination_id" json:"account_destination_id"`
	Amount               float64   `bson:"amount" json:"amount"`
	Date                 time.Time `bson:"date" json:"date"`
}

// TransferDepositResponse struct to illustrate a deposit transfer response
type TransferDepositResponse struct {
	AccountOriginID string    `bson:"account_origin_id" json:"account_origin_id"`
	Amount          float64   `bson:"amount" json:"amount"`
	Date            time.Time `bson:"date" json:"date"`
}

// TransferListResponse struct to illustrate a transfer list response
type TransferListResponse struct {
	Withdraws []TransferWithDrawResponse `bson:"withdraws" json:"withdraws"`
	Deposits  []TransferDepositResponse  `bson:"deposits" json:"deposits"`
}

// NewTransferListResponse create a transfer list response based on transfers models
func NewTransferListResponse(withdraws []model.Transfer, deposits []model.Transfer) TransferListResponse {

	withdrawsResponse := make([]TransferWithDrawResponse, 0)
	depositsResponse := make([]TransferDepositResponse, 0)

	for _, withdraw := range withdraws {
		withdrawsResponse = append(withdrawsResponse, TransferWithDrawResponse{
			AccountDestinationID: withdraw.AccountDestinationID,
			Amount:               withdraw.Amount,
			Date:                 withdraw.CreatedAt,
		})
	}

	for _, deposit := range deposits {
		depositsResponse = append(depositsResponse, TransferDepositResponse{
			AccountOriginID: deposit.AccountOriginID,
			Amount:          deposit.Amount,
			Date:            deposit.CreatedAt,
		})
	}

	return TransferListResponse{
		Withdraws: withdrawsResponse,
		Deposits:  depositsResponse,
	}
}
