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

func (s Store) CriaTransacao(transacao *Transacao) (err error) {
	result := s.db.Raw("SELECT \"criatransacao\"(?,?,?,?)", transacao.ClienteID, transacao.Tipo, transacao.Valor, transacao.Descricao).Row()
	if result.Err() != nil {
		err = fmt.Errorf("erro criando a transacao no banco de dados: %s", result.Err().Error())
	}
	return
}
