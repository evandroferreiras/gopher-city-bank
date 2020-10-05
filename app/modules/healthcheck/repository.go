//go:generate mockery --name Repository --filename=repository.go

package healthcheck

import (
	"github.com/evandroferreiras/gopher-city-bank/app/common/envvar"
	"github.com/sirupsen/logrus"
)

// Repository is an interface to Healthcheck repository
type Repository interface {
	Ping() (bool, error)
}

// BuildRepository is a factory constructor for Healthcheck Repository
func BuildRepository() Repository {
	if envvar.UsingMemoryDB() {
		return NewInMemoryDBRepository()
	}
	logrus.Fatal("Repository not implemented for healthcheck")
	return nil
}
