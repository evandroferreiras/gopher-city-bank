package healthcheck

import (
	"github.com/evandroferreiras/gopher-city-bank/app/db"
	"gorm.io/gorm"
)

type repositoryORM struct {
	db gorm.DB
}

// NewORMRepository is a constructor to ORM Healthcheck repository
func NewORMRepository() Repository {
	d := db.DB()
	return &repositoryORM{
		db: *d,
	}
}

func (r *repositoryORM) Ping() (bool, error) {
	sqlDB, err := r.db.DB()
	if err != nil {
		return false, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return false, err
	}
	return true, nil
}
