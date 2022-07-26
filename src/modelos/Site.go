package modelos

// Site representa um site cadastrado no sistema
type Site struct {
	IDSITE     uint64 `json:"idsite,omitempty" gorm:"primaryKey;column:IDSITE"`
	NOME       string `json:"nome,omitempty" gorm:"column:NOME"`
	SIGLA      string `json:"sigla,omitempty" gorm:"column:SIGLA"`
	UF         string `json:"uf,omitempty" gorm:"column:UF"`
	ENDERECO   string `json:"endereco,omitempty" gorm:"column:ENDERECO"`
	ID_USUARIO string `json:"id_usuario,omitempty" gorm:"column:ID_USUARIO"`
}
