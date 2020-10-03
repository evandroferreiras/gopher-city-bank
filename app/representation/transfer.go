package representation

// TransferBody struct to illustrate post transfer
type TransferBody struct {
	AccountDestinationID string  `bson:"account_destination_id" json:"account_destination_id" validate:"required"`
	Amount               float64 `bson:"amount" json:"amount" validate:"required"`
}
