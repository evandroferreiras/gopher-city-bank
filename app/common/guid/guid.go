package guid

import (
	"github.com/google/uuid"
)

// NewGUID returns a guid as a string
func NewGUID() string {
	return uuid.New().String()
}
