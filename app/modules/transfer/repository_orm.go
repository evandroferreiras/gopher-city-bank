package transfer

import (
	"context"

	"github.com/evandroferreiras/gopher-city-bank/app/common/constant"
	"github.com/evandroferreiras/gopher-city-bank/app/common/guid"
	"github.com/evandroferreiras/gopher-city-bank/app/db"
	"github.com/sirupsen/logrus"

	"github.com/evandroferreiras/gopher-city-bank/app/model"
	"gorm.io/gorm"
)

type repositoryORM struct {
	db gorm.DB
}

// NewORMRepository is a constructor to ORM Transfer repository
func NewORMRepository() Repository {
	d := db.DB()
	return &repositoryORM{
		db: *d,
	}
}

// GetAccount return a account given an id
func (r repositoryORM) GetAccount(id string) (model.Account, error) {
	account := model.Account{}
	tx := r.db.Where(&model.Account{ID: id}).First(&account)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			logrus.Infof("Not found an account with ID:%+v", id)
			return account, nil
		}
		return account, tx.Error
	}

	return account, nil
}

// UpdateAccountBalance subtracts the amount of money from accountID
func (r repositoryORM) UpdateAccountBalance(ctx context.Context, id string, newBalance float64) error {
	tx := getTransactionFromCtx(ctx)
	rTx := tx.Where(&model.Account{ID: id}).Updates(model.Account{Balance: newBalance})
	if rTx.Error != nil {
		return rTx.Error
	}
	return nil
}

func (r repositoryORM) LogTransfer(ctx context.Context, transfer model.Transfer) error {
	tx := getTransactionFromCtx(ctx)
	transfer.ID = guid.NewGUID()
	rTx := tx.Create(&transfer)
	if rTx.Error != nil {
		return rTx.Error
	}
	return nil
}

func (r repositoryORM) StartTransaction() (interface{}, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}

func (r repositoryORM) CommitTransaction(ctx context.Context) {
	tx := getTransactionFromCtx(ctx)
	tx.Commit()
}

func (r repositoryORM) RollbackTransaction(ctx context.Context) {
	tx := getTransactionFromCtx(ctx)
	tx.Rollback()
}

func (r repositoryORM) GetAllWithdrawsOf(accountOriginID string) ([]model.Transfer, error) {
	var withdraws []model.Transfer
	tx := r.db.Where(&model.Transfer{AccountOriginID: accountOriginID}).Find(&withdraws)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return withdraws, nil
}

func (r repositoryORM) GetAllDepositsTo(accountOriginID string) ([]model.Transfer, error) {
	var deposits []model.Transfer
	tx := r.db.Where(&model.Transfer{AccountDestinationID: accountOriginID}).Find(&deposits)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return deposits, nil
}

func getTransactionFromCtx(ctx context.Context) *gorm.DB {
	return ctx.Value(constant.TransactionCtxKey).(*gorm.DB)
}
