package model

import (
	"time"
)

// Transfer struct to illustrate database object
type Transfer struct {
	ID                   string    `bson:"_id" json:"id" gorm:"type:varchar(36)"`
	AccountOriginID      string    `bson:"account_origin_id" json:"account_origin_id" gorm:"type:varchar(36);index"`
	AccountDestinationID string    `bson:"account_destination_id" json:"account_destination_id" gorm:"type:varchar(36);index"`
	Amount               float64   `bson:"amount" json:"amount"`
	CreatedAt            time.Time `bson:"created_at" json:"created_at"`
	AccountOrigin        Account   `gorm:"foreignkey:account_origin_id;references:id"`
	AccountDestination   Account   `gorm:"foreignkey:account_destination_id;references:id"`
}
