package transaction

const (
	TYPE_CREDIT = "c"
	TYPE_DEBIT  = "d"
)

type Transaction struct {
	ID          uint   `gorm:"primaryKey"`
	Description string `gorm:"descricao" json:"descricao"`
	Type        string `gorm:"tipo" json:"tipo"`
	Value       uint   `gorm:"valor" json:"valor"`
}
