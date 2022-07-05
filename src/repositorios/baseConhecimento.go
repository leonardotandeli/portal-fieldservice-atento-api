package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

// Posts representa um repositório de publicações
type Posts struct {
	db *sql.DB
}

// NovoRepositorioDePosts cria um repositório de publicações
func NovoRepositorioDePosts(db *sql.DB) *Posts {
	return &Posts{db}
}

// Criar insere uma publicação no banco de dados
func (repositorio Posts) Criar(post modelos.Post) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO BDC_POSTS(TITULO, CONTEUDO, ID_CATEGORIA, ID_USUARIO, ID_SITE, ID_CLIENTE) VALUES (?, ?, ?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(post.TITULO, post.CONTEUDO, post.ID_CATEGORIA, post.ID_USUARIO, post.ID_SITE, post.ID_CLIENTE)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// Atualizar altera os dados de uma publicação no banco de dados
func (repositorio Posts) Atualizar(postID uint64, post modelos.Post) error {
	statement, erro := repositorio.db.Prepare("UPDATE BDC_POSTS SET TITULO = ?, CONTEUDO = ?, ID_CATEGORIA = ?, ID_USUARIO = ?, ID_SITE = ?, ID_CLIENTE = ? WHERE IDPOST = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(post.TITULO, post.CONTEUDO, post.ID_CATEGORIA, post.ID_USUARIO, post.ID_SITE, post.ID_CLIENTE, postID); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui uma publicação do banco de dados
func (repositorio Posts) Deletar(postID uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE FROM BDC_POSTS WHERE IDPOST = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(postID); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorID busca um chamado do banco de dados pelo id
func (repositorio Posts) BuscarPorID(ID uint64) (modelos.Post, error) {

	linhas, erro := repositorio.db.Query(
		"SELECT P.IDPOST, P.TITULO, P.CONTEUDO, P.ID_CATEGORIA, P.ID_USUARIO, P.ID_SITE, P.ID_CLIENTE, C.IDCATEGORIA, C.NOME, U.NOME, U.RE, S.IDSITE, S.NOME, B.IDCLIENTE, B.NOME, P.DATA_CRIACAO FROM BDC_POSTS P INNER JOIN BDC_CATEGORIAS C ON P.ID_CATEGORIA = C.IDCATEGORIA INNER JOIN USUARIOS U ON P.ID_USUARIO = U.IDUSUARIO INNER JOIN SITES S ON P.ID_SITE = S.IDSITE INNER JOIN CLIENTES B ON P.ID_CLIENTE = B.IDCLIENTE WHERE P.IDPOST = ?", ID)
	if erro != nil {
		return modelos.Post{}, erro
	}
	defer linhas.Close()

	var post modelos.Post

	if linhas.Next() {
		if erro = linhas.Scan(
			&post.IDPOST,
			&post.TITULO,
			&post.CONTEUDO,
			&post.ID_CATEGORIA,
			&post.ID_USUARIO,
			&post.ID_SITE,
			&post.ID_CLIENTE,
			&post.Categoria.IDCATEGORIA,
			&post.Categoria.NOME,
			&post.Usuario.NOME,
			&post.Usuario.RE,
			&post.Site.IDSITE,
			&post.Site.NOME,
			&post.Cliente.IDCLIENTE,
			&post.Cliente.NOME,
			&post.DATA_CRIACAO,
		); erro != nil {
			return modelos.Post{}, erro
		}
	}

	return post, nil

}

func (repositorio Posts) BuscaPorNome(nomeDoc string) ([]modelos.Post, error) {

	//filtro nome ou login
	nomeDoc = fmt.Sprintf("%%%s%%", nomeDoc)

	linhas, erro := repositorio.db.Query(

		"SELECT P.IDPOST, P.TITULO, P.CONTEUDO, P.ID_CATEGORIA, P.ID_USUARIO, P.ID_SITE, P.ID_CLIENTE, C.NOME, U.NOME, S.NOME, P.DATA_CRIACAO FROM BDC_POSTS P INNER JOIN BDC_CATEGORIAS C ON P.ID_CATEGORIA = C.IDCATEGORIA INNER JOIN USUARIOS U ON P.ID_USUARIO = U.IDUSUARIO INNER JOIN SITES S ON P.ID_SITE = S.IDSITE INNER JOIN CLIENTES B ON P.ID_CLIENTE = B.IDCLIENTE WHERE P.TITULO LIKE ?",
		nomeDoc,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var Posts []modelos.Post
	for linhas.Next() {
		var post modelos.Post

		if erro = linhas.Scan(
			&post.IDPOST,
			&post.TITULO,
			&post.CONTEUDO,
			&post.ID_CATEGORIA,
			&post.ID_USUARIO,
			&post.ID_SITE,
			&post.ID_CLIENTE,
			&post.Categoria.NOME,
			&post.Usuario.NOME,
			&post.Site.NOME,
			&post.DATA_CRIACAO,
		); erro != nil {
			return nil, erro
		}

		Posts = append(Posts, post)
	}

	return Posts, nil
}

// Buscar traz todas as publicações
func (repositorio Posts) BuscarTodos(urlCategoria string, urlCliente string) ([]modelos.Post, error) {

	urlCategoria = fmt.Sprintf("%s", urlCategoria) // filtra por parametro na url ?categoria=[id]
	fmt.Println(urlCategoria)
	urlCliente = fmt.Sprintf("%s", urlCliente) // filtra por parametro na url ?cliente=[id]
	fmt.Println(urlCliente)

	linhas, erro := repositorio.db.Query(
		"SELECT P.IDPOST, P.TITULO, P.CONTEUDO, P.ID_CATEGORIA, P.ID_USUARIO, P.ID_SITE, P.ID_CLIENTE, C.NOME, U.NOME, S.NOME, B.IDCLIENTE, B.NOME, P.DATA_CRIACAO FROM BDC_POSTS P INNER JOIN BDC_CATEGORIAS C ON P.ID_CATEGORIA = C.IDCATEGORIA INNER JOIN USUARIOS U ON P.ID_USUARIO = U.IDUSUARIO INNER JOIN SITES S ON P.ID_SITE = S.IDSITE INNER JOIN CLIENTES B ON P.ID_CLIENTE = B.IDCLIENTE INNER JOIN BDC_CATEGORIAS N ON P.ID_CATEGORIA = N.IDCATEGORIA ORDER BY P.IDPOST DESC",
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var Posts []modelos.Post

	for linhas.Next() {
		var post modelos.Post

		if erro = linhas.Scan(
			&post.IDPOST,
			&post.TITULO,
			&post.CONTEUDO,
			&post.ID_CATEGORIA,
			&post.ID_USUARIO,
			&post.ID_SITE,
			&post.ID_CLIENTE,
			&post.Categoria.NOME,
			&post.Usuario.NOME,
			&post.Site.NOME,
			&post.Cliente.IDCLIENTE,
			&post.Cliente.NOME,
			&post.DATA_CRIACAO,
		); erro != nil {
			return nil, erro
		}

		Posts = append(Posts, post)
	}

	return Posts, nil

}

func (repositorio Posts) Busca(urlBusca string) ([]modelos.Post, error) {

	urlBusca = fmt.Sprintf("%%%s%%", urlBusca) // filtra por parametro na url ?categoria=[id]
	fmt.Println(urlBusca)

	linhas, erro := repositorio.db.Query(

		"SELECT P.IDPOST, P.TITULO, P.CONTEUDO, P.ID_CATEGORIA, P.ID_USUARIO, P.ID_SITE, P.ID_CLIENTE, C.NOME, U.NOME, S.NOME, P.DATA_CRIACAO FROM BDC_POSTS P INNER JOIN BDC_CATEGORIAS C ON P.ID_CATEGORIA = C.IDCATEGORIA INNER JOIN USUARIOS U ON P.ID_USUARIO = U.IDUSUARIO INNER JOIN SITES S ON P.ID_SITE = S.IDSITE INNER JOIN CLIENTES B ON P.ID_CLIENTE = B.IDCLIENTE WHERE P.TITULO LIKE ?", urlBusca)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var Posts []modelos.Post
	for linhas.Next() {
		var post modelos.Post

		if erro = linhas.Scan(
			&post.IDPOST,
			&post.TITULO,
			&post.CONTEUDO,
			&post.ID_CATEGORIA,
			&post.ID_USUARIO,
			&post.ID_SITE,
			&post.ID_CLIENTE,
			&post.Categoria.NOME,
			&post.Usuario.NOME,
			&post.Site.NOME,
			&post.DATA_CRIACAO,
		); erro != nil {
			return nil, erro
		}

		Posts = append(Posts, post)
	}

	return Posts, nil
}

// Buscar traz todas as publicações
func (repositorio Posts) BuscarPorString(urlCategoria string, urlCliente string) ([]modelos.Post, error) {

	urlCategoria = fmt.Sprintf("%s", urlCategoria) // filtra por parametro na url ?categoria=[id]
	fmt.Println(urlCategoria)
	urlCliente = fmt.Sprintf("%s", urlCliente) // filtra por parametro na url ?cliente=[id]
	fmt.Println(urlCliente)

	linhas, erro := repositorio.db.Query(
		"SELECT P.IDPOST, P.TITULO, P.CONTEUDO, P.ID_CATEGORIA, P.ID_USUARIO, P.ID_SITE, P.ID_CLIENTE, C.NOME, U.NOME, S.NOME, P.DATA_CRIACAO FROM BDC_POSTS P INNER JOIN BDC_CATEGORIAS C ON P.ID_CATEGORIA = C.IDCATEGORIA INNER JOIN USUARIOS U ON P.ID_USUARIO = U.IDUSUARIO INNER JOIN SITES S ON P.ID_SITE = S.IDSITE INNER JOIN CLIENTES B ON P.ID_CLIENTE = B.IDCLIENTE INNER JOIN BDC_CATEGORIAS N ON P.ID_CATEGORIA = N.IDCATEGORIA WHERE N.IDCATEGORIA = ? OR B.IDCLIENTE = ? ORDER BY P.IDPOST DESC",
		urlCategoria, urlCliente,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var Posts []modelos.Post

	for linhas.Next() {
		var post modelos.Post

		if erro = linhas.Scan(
			&post.IDPOST,
			&post.TITULO,
			&post.CONTEUDO,
			&post.ID_CATEGORIA,
			&post.ID_USUARIO,
			&post.ID_SITE,
			&post.ID_CLIENTE,
			&post.Categoria.NOME,
			&post.Usuario.NOME,
			&post.Site.NOME,
			&post.DATA_CRIACAO,
		); erro != nil {
			return nil, erro
		}

		Posts = append(Posts, post)
	}

	return Posts, nil

}

// Buscar traz todas as publicações
func (repositorio Posts) BuscarPorStringCat(urlCategoria string, urlCliente string) ([]modelos.Post, error) {

	urlCategoria = fmt.Sprintf("%s", urlCategoria) // filtra por parametro na url ?categoria=[id]
	fmt.Println(urlCategoria)
	urlCliente = fmt.Sprintf("%s", urlCliente) // filtra por parametro na url ?cliente=[id]
	fmt.Println(urlCliente)

	linhas, erro := repositorio.db.Query(
		"SELECT P.IDPOST, P.TITULO, P.CONTEUDO, P.ID_CATEGORIA, P.ID_USUARIO, P.ID_SITE, P.ID_CLIENTE, C.NOME, U.NOME, S.NOME, P.DATA_CRIACAO FROM BDC_POSTS P INNER JOIN BDC_CATEGORIAS C ON P.ID_CATEGORIA = C.IDCATEGORIA INNER JOIN USUARIOS U ON P.ID_USUARIO = U.IDUSUARIO INNER JOIN SITES S ON P.ID_SITE = S.IDSITE INNER JOIN CLIENTES B ON P.ID_CLIENTE = B.IDCLIENTE INNER JOIN BDC_CATEGORIAS N ON P.ID_CATEGORIA = N.IDCATEGORIA WHERE N.IDCATEGORIA = ? AND B.IDCLIENTE = ? ORDER BY P.IDPOST DESC",
		urlCategoria, urlCliente,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var Posts []modelos.Post

	for linhas.Next() {
		var post modelos.Post

		if erro = linhas.Scan(
			&post.IDPOST,
			&post.TITULO,
			&post.CONTEUDO,
			&post.ID_CATEGORIA,
			&post.ID_USUARIO,
			&post.ID_SITE,
			&post.ID_CLIENTE,
			&post.Categoria.NOME,
			&post.Usuario.NOME,
			&post.Site.NOME,
			&post.DATA_CRIACAO,
		); erro != nil {
			return nil, erro
		}

		Posts = append(Posts, post)
	}

	return Posts, nil

}
