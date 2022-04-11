package modelos

import (
	"time"
)

// Post representa uma publicação feita por um usuário
type Post struct {
	IDPOST       uint64    `json:"idpost,omitempty"`
	TITULO       string    `json:"titulo,omitempty"`
	CONTEUDO     string    `json:"conteudo,omitempty"`
	ID_CATEGORIA string    `json:"id_categoria,omitempty"`
	ID_USUARIO   string    `json:"id_usuario,omitempty"`
	ID_SITE      string    `json:"id_site,omitempty"`
	ID_CLIENTE   string    `json:"id_cliente,omitempty"`
	DATA_CRIACAO time.Time `json:"data_criacao,omitempty"`
	Usuario      Usuario
	Categoria    Post_Categoria
	Site         Site
}

/*

// Preparar vai chamar os métodos para validar e formatar a publicação recebida
func (post *Post) Preparar() error {
	if erro := post.validar(); erro != nil {
		return erro
	}

	post.formatar()
	return nil
}

func (post *Post) validar() error {
	if post.Titulo == "" {
		return errors.New("O título é obrigatório e não pode estar em branco")
	}

	if post.Conteudo == "" {
		return errors.New("O conteúdo é obrigatório e não pode estar em branco")
	}

	return nil
}

func (post *Post) formatar() {
	post.Titulo = strings.TrimSpace(post.Titulo)
	post.Conteudo = strings.TrimSpace(post.Conteudo)
}
*/
