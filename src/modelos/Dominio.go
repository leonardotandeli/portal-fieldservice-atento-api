package modelos

// Dominio representa um domínio cadastrado no sistema
type Dominio struct {
	IDDOMINIO uint64 `json:"iddominio,omitempty" gorm:"primaryKey;column:IDDOMINIO"`
	NOME      string `json:"nome,omitempty"  gorm:"column:NOME"`
}
