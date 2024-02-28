package extrato

import (
	"rinha_backend2024_q1/internal/transacao"
	"time"
)

type Saldo struct {
	Total       int       `json:"total"`
	DataExtrato time.Time `json:"data_extrato"`
	Limite      uint      `json:"limite"`
}

type Extrato struct {
	Saldo             Saldo                 `json:"saldo"`
	UltimasTransacoes []transacao.Transacao `json:"ultimas_transacoes"`
}
