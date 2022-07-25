package modelos

// Dac representa um dac cadastrado em sistema
type Dac struct {
	IDDAC uint64 `json:"iddac,omitempty" gorm:"primaryKey;column:IDCLIENTE"`
	NOME  string `json:"nome,omitempty" gorm:"column:NOME"`
}
