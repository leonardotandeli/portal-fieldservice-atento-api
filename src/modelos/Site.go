package modelos

// Site representa um site cadastrado no sistema
type Site struct {
	IDSITE     uint64 `json:"idsite,omitempty"`
	NOME       string `json:"nome,omitempty"`
	SIGLA      string `json:"sigla,omitempty"`
	UF         string `json:"uf,omitempty"`
	ENDERECO   string `json:"endereco,omitempty"`
	ID_USUARIO string `json:"id_usuario,omitempty"`
}
