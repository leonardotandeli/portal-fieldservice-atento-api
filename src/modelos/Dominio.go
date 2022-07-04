package modelos

// Dominio representa um domínio cadastrado no sistema
type Dominio struct {
	IDDOMINIO uint64 `json:"iddominio,omitempty"`
	NOME      string `json:"nome,omitempty"`
}
