package transaction

import (
	"fmt"

	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) Store {
	return Store{db}
}

func (s Store) Create(transaction *Transaction) (*Transaction, error) {
	result := s.db.Create(transaction)
	if result.Error != nil {
		return nil, fmt.Errorf("error creating the transaction in the database: %s", result.Error)
	}
	return transaction, nil
}
