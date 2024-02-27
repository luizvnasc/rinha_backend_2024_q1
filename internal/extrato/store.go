package extrato

import (
	"errors"
	"rinha_backend2024_q1/internal/cliente"
	"time"

	"gorm.io/gorm"
)

var store extratoStore

type extratoStore struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) {
	store = extratoStore{db}
}

func (s extratoStore) getExtrato(clienteID uint) (*Extrato, error) {
	cliente := &cliente.Cliente{}
	if err := s.db.Preload("Transacoes", func(db *gorm.DB) *gorm.DB {
		return db.Order("transacao.criado_em DESC").Limit(10)
	}).First(&cliente, clienteID).Error; err != nil {
		return nil, errors.New("erro ao obter extrato: " + err.Error())
	}

	if cliente == nil {
		return nil, errors.New("cliente não encontrado")
	}

	for i, t := range cliente.Transacoes {
		cliente.Transacoes[i] = t.Sanitize()
	}
	return &Extrato{
		Saldo: Saldo{
			DataExtrato: time.Now(),
			Total:       cliente.Saldo,
			Limite:      cliente.Limite,
		},
		UltimasTransacoes: cliente.Transacoes,
	}, nil
}
