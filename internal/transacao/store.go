package transacao

import (
	"fmt"

	"gorm.io/gorm"
)

var store Store

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) {
	store = Store{db}
}

func GetStore() Store {
	return store
}

func (s Store) CriaTransacao(transacao *Transacao) (*Transacao, error) {
	result := s.db.Create(transacao)
	if result.Error != nil {
		return nil, fmt.Errorf("ewrro criando a transacao no banco de dados: %s", result.Error)
	}
	return transacao, nil
}
