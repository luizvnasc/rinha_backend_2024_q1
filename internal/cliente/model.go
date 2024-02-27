package cliente

import "rinha_backend2024_q1/internal/transacao"

type Cliente struct {
	ID         uint                  `gorm:"id"`
	Limite     uint                  `gorm:"limite"`
	Saldo      int                   `gorm:"saldo"`
	Transacoes []transacao.Transacao `gorm:"foreignKey:ClienteID;references:ID"`
}
