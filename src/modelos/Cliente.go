package modelos

// Cliente representa um cliente cadastrado em sistema
type Cliente struct {
	IDCLIENTE uint64 `json:"idcliente,omitempty" gorm:"primaryKey;column:IDCLIENTE"`
	NOME      string `json:"nome,omitempty"  gorm:"column:NOME"`
	LOGO_URL  string `json:"logo_url,omitempty"  gorm:"column:LOGO_URL"`
}
