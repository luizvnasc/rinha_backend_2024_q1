package transacao

import "time"

const (
	TIPO_CREDITO = "c"
	TIPO_DEBITO  = "d"
)

type Transacao struct {
	ID        uint      `gorm:"primaryKey"`
	ClienteID uint      `gorm:"cliente_id"`
	Descricao string    `gorm:"descricao" json:"descricao"`
	Tipo      string    `gorm:"tipo" json:"tipo"`
	Valor     uint64    `gorm:"valor" json:"valor"`
	CriadoEm  time.Time `gorm:"criado_em" json:"criado_em"`
}
