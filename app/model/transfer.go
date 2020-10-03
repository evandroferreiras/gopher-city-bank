package model

import (
	"time"
)

// Transfer struct to illustrate database object
type Transfer struct {
	ID                   string    `bson:"_id" json:"id"`
	AccountOriginID      string    `bson:"account_origin_id" json:"account_origin_id"`
	AccountDestinationID string    `bson:"account_destination_id" json:"account_destination_id"`
	Amount               float64   `bson:"amount" json:"amount"`
	CreatedAt            time.Time `bson:"created_at" json:"created_at"`
}
