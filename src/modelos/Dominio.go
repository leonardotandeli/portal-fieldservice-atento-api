package modelos

// Dominio representa os domínios cadastrados no sistema
type Dominio struct {
	IDDOMINIO uint64 `json:"iddominio,omitempty"`
	NOME      string `json:"nome,omitempty"`
}
