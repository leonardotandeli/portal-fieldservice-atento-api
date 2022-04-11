package modelos

// Post_Categoria representa uma categoria de posts cadastrado em sistema
type Post_Categoria struct {
	IDCATEGORIA uint64 `json:"idcategoria,omitempty"`
	NOME        string `json:"NOME,omitempty"`
}
