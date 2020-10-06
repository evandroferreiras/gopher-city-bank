package login

import (
	"github.com/evandroferreiras/gopher-city-bank/app/db"
	"github.com/evandroferreiras/gopher-city-bank/app/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type repositoryORM struct {
	db gorm.DB
}

// NewORMRepository is a constructor to ORM Login repository
func NewORMRepository() Repository {
	d := db.DB()
	return &repositoryORM{
		db: *d,
	}
}

func (r repositoryORM) GetAccountByCpf(cpf string) (model.Account, error) {
	account := model.Account{}
	tx := r.db.Where(&model.Account{Cpf: cpf}).First(&account)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			logrus.Infof("Not found an account with CPF:'%+v'", cpf)
			return model.EmptyAccount, nil
		}
		return model.EmptyAccount, tx.Error
	}
	return account, nil
}
