package modelos

// Cliente representa um cliente cadastrado em sistema
type Cliente struct {
	IDCLIENTE uint64 `json:"idcliente,omitempty"`
	NOME      string `json:"nome,omitempty"`
	LOGO_URL  string `json:"logo_url,omitempty"`
}
