package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Categorias representa um repositório de categorias
type Categorias struct {
	db *sql.DB
}

// NovoRepositorioDeCategorias cria um repositório de categorias
func NovoRepositorioDeCategorias(db *sql.DB) *Categorias {
	return &Categorias{db}
}

// CriarCategoria insere uma categoria no banco de dados
func (repositorio Categorias) CriarCategoria(cat modelos.Post_Categoria) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO BDC_CATEGORIAS(NOME) VALUES (?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(cat.NOME)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// Atualizar altera os dados de uma categoria no banco de dados
func (repositorio Categorias) AtualizarCategoria(catID uint64, cat modelos.Post_Categoria) error {
	statement, erro := repositorio.db.Prepare("UPDATE BDC_CATEGORIAS SET NOME = ? WHERE IDCATEGORIA = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(cat.NOME, catID); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui uma Categoria do banco de dados
func (repositorio Categorias) DeletarCategoria(catID uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE FROM BDC_CATEGORIAS WHERE IDCATEGORIA = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(catID); erro != nil {
		return erro
	}

	return nil
}

// BuscarCategoria traz todas as categorias do banco de dados
func (repositorio Categorias) BuscarCategoria() ([]modelos.Post_Categoria, error) {

	linhas, erro := repositorio.db.Query(
		"SELECT IDCATEGORIA, NOME FROM BDC_CATEGORIAS ORDER BY IDCATEGORIA",
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var Categoria []modelos.Post_Categoria

	for linhas.Next() {
		var categoria modelos.Post_Categoria

		if erro = linhas.Scan(
			&categoria.IDCATEGORIA,
			&categoria.NOME,
		); erro != nil {
			return nil, erro
		}

		Categoria = append(Categoria, categoria)
	}

	return Categoria, nil

}

// BuscarPorID busca um chamado do banco de dados pelo id
func (repositorio Categorias) BuscarCategoriaPorID(ID uint64) (modelos.Post_Categoria, error) {

	linhas, erro := repositorio.db.Query(
		"SELECT IDCATEGORIA, NOME FROM BDC_CATEGORIAS WHERE IDCATEGORIA = ?", ID)
	if erro != nil {
		return modelos.Post_Categoria{}, erro
	}
	defer linhas.Close()

	var categoria modelos.Post_Categoria

	if linhas.Next() {
		if erro = linhas.Scan(
			&categoria.IDCATEGORIA,
			&categoria.NOME,
		); erro != nil {
			return modelos.Post_Categoria{}, erro
		}
	}

	return categoria, nil

}
