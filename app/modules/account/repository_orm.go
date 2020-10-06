package account

import (
	"github.com/evandroferreiras/gopher-city-bank/app/common/guid"
	"github.com/evandroferreiras/gopher-city-bank/app/db"
	"github.com/evandroferreiras/gopher-city-bank/app/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type repositoryORM struct {
	db gorm.DB
}

// NewORMRepository is a constructor to ORM Account repository
func NewORMRepository() Repository {
	d := db.DB()
	return &repositoryORM{
		db: *d,
	}
}

// Create a new account.
func (r *repositoryORM) Create(newAccount model.Account) (model.Account, error) {
	newAccount.ID = guid.NewGUID()
	tx := r.db.Create(&newAccount)
	if tx.Error != nil {
		return model.EmptyAccount, tx.Error
	}
	return newAccount, nil
}

// GetAccounts lists all accounts
func (r *repositoryORM) GetAccounts() ([]model.Account, error) {
	var accounts []model.Account
	tx := r.db.Model(&model.Account{}).Find(&accounts)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return accounts, nil
}

// GetAccount return a account given an id
func (r *repositoryORM) GetAccount(id string) (model.Account, error) {
	account := model.Account{}
	tx := r.db.Where(&model.Account{ID: id}).First(&account)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			logrus.Infof("Not found an account with ID:%+v", id)
			return model.EmptyAccount, nil
		}
		return model.EmptyAccount, tx.Error
	}

	return account, nil
}
