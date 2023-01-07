package modelos

import (
	"time"
)

// Post_Categoria representa uma categoria de posts cadastrado em sistema
type Post_Categoria struct {
	IDCATEGORIA uint64 `json:"idcategoria,omitempty"`
	NOME        string `json:"NOME,omitempty"`
	ID_CLIENTE  string `json:"id_cliente,omitempty"`
	STATUS      string `json:"status,omitempty"`
}

// Post_SubCategoria representa uma sub-categoria de posts cadastrado em sistema
type Post_SubCategoria struct {
	IDSUBCATEGORIA uint64 `json:"idsubcategoria,omitempty"`
	NOME           string `json:"NOME,omitempty"`
	ID_CATEGORIA   string `json:"id_categoria,omitempty"`
	STATUS         string `json:"status,omitempty"`
}

// Post representa uma publicação feita por um usuário
type Post struct {
	IDPOST          uint64    `json:"idpost,omitempty"`
	TITULO          string    `json:"titulo,omitempty"`
	CONTEUDO        string    `json:"conteudo,omitempty"`
	ID_CATEGORIA    string    `json:"id_categoria,omitempty"`
	ID_SUBCATEGORIA string    `json:"id_subcategoria,omitempty"`
	ID_USUARIO      string    `json:"id_usuario,omitempty"`
	ID_SITE         string    `json:"id_site,omitempty"`
	ID_CLIENTE      string    `json:"id_cliente,omitempty"`
	DATA_CRIACAO    time.Time `json:"data_criacao,omitempty"`
	Usuario         Usuario
	Categoria       Post_Categoria
	SubCategoria    Post_SubCategoria
	Site            Site
	Cliente         Cliente
}
