package inmemorydb

import (
	"fmt"
	"sync"
	"time"

	"github.com/evandroferreiras/gopher-city-bank/app/model"
)

// MemoryDatabase struct
type MemoryDatabase struct {
	sync.RWMutex
	accounts  map[int]*model.Account
	transfers map[int]*model.Transfer
}

var memoryDatabase *MemoryDatabase

var once sync.Once

// GetMemoryDB returns a map of Accounts.
func GetMemoryDB() *MemoryDatabase {
	once.Do(func() {
		memoryDatabase = &MemoryDatabase{
			accounts:  make(map[int]*model.Account),
			transfers: make(map[int]*model.Transfer),
		}
	})
	return memoryDatabase
}

// AddAccount adds an account to the inMemoryDatabase
func AddAccount(newAccount model.Account) model.Account {
	db := GetMemoryDB()
	db.Lock()
	defer db.Unlock()
	idx := len(db.accounts)

	account := model.Account{
		ID:        fmt.Sprintf("%d", idx+1),
		Name:      newAccount.Name,
		Cpf:       newAccount.Cpf,
		Secret:    newAccount.Secret,
		Balance:   newAccount.Balance,
		CreatedAt: time.Now(),
	}
	db.accounts[idx] = &account
	return account
}

// GetAccounts returns all accounts from the inMemoryDatabase
func GetAccounts(page int, size int) []model.Account {
	db := GetMemoryDB()
	db.RLock()
	defer db.RUnlock()

	var accounts = make([]model.Account, 0)

	total := page * size
	initial := total - (size)
	for i, account := range db.accounts {
		if i >= initial && i < total {
			accounts = append(accounts, *account)
		}
	}
	return accounts
}

// GetAccount returns an account given an id
func GetAccount(id string) model.Account {
	db := GetMemoryDB()
	db.RLock()
	defer db.RUnlock()

	for _, account := range db.accounts {
		if account.ID == id {

			return *account
		}
	}
	return model.EmptyAccount
}

// GetAccountByCpf returns an account given an cpf
func GetAccountByCpf(cpf string) model.Account {
	db := GetMemoryDB()
	db.RLock()
	defer db.RUnlock()

	for _, account := range db.accounts {
		if account.Cpf == cpf {
			return *account
		}
	}
	return model.EmptyAccount
}

// UpdateAccountBalance updates and account balance
func UpdateAccountBalance(id string, newBalance float64) {
	db := GetMemoryDB()
	db.Lock()
	defer db.Unlock()

	for _, account := range db.accounts {
		if account.ID == id {
			account.Balance = newBalance
			return
		}
	}
}

// LogTransfer register transfer
func LogTransfer(newTransfer model.Transfer) {
	db := GetMemoryDB()
	db.Lock()
	defer db.Unlock()

	idx := len(db.transfers)
	newTransfer.ID = string(idx)
	db.transfers[idx] = &newTransfer
}

// GetAllWithdrawsOf account origin
func GetAllWithdrawsOf(accountOriginID string) []model.Transfer {
	db := GetMemoryDB()
	db.RLock()
	defer db.RUnlock()

	var transfers = make([]model.Transfer, 0)

	for _, transfer := range db.transfers {
		if transfer.AccountOriginID == accountOriginID {
			transfers = append(transfers, *transfer)
		}
	}
	return transfers
}

// GetAllDepositsTo account origin
func GetAllDepositsTo(accountOriginID string) []model.Transfer {
	db := GetMemoryDB()
	db.RLock()
	defer db.RUnlock()

	var transfers = make([]model.Transfer, 0)

	for _, transfer := range db.transfers {
		if transfer.AccountDestinationID == accountOriginID {
			transfers = append(transfers, *transfer)
		}
	}
	return transfers
}
