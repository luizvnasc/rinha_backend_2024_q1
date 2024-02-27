package transacao

import "time"

const (
	TIPO_CREDITO = "c"
	TIPO_DEBITO  = "d"
)

type Transacao struct {
	ID        uint      `gorm:"primaryKey" json:"id,omitempty"`
	ClienteID uint      `gorm:"cliente_id" json:"cliente_id,omitempty"`
	Descricao string    `gorm:"descricao" json:"descricao"`
	Tipo      string    `gorm:"tipo" json:"tipo"`
	Valor     uint64    `gorm:"valor" json:"valor"`
	CriadoEm  time.Time `gorm:"criado_em" json:"criado_em"`
}

func (t Transacao) Sanitize() Transacao {
	t.ClienteID = 0
	t.ID = 0
	return t
}
