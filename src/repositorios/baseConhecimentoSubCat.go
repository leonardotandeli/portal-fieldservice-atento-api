package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type SubCategorias struct {
	db *sql.DB
}

func NovoRepositorioDeSubCategorias(db *sql.DB) *SubCategorias {
	return &SubCategorias{db}
}

func (repositorio SubCategorias) CriarSubCategoria(cat modelos.Post_SubCategoria) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO BDC_SUBCATEGORIAS(NOME, ID_CATEGORIA) VALUES (?,?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(cat.NOME, cat.ID_CATEGORIA)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio SubCategorias) AtualizarSubCategoria(catID uint64, cat modelos.Post_SubCategoria) error {
	statement, erro := repositorio.db.Prepare("UPDATE BDC_SUBCATEGORIAS SET NOME = ?, ID_CATEGORIA = ? WHERE IDSUBCATEGORIA = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(cat.NOME, cat.ID_CATEGORIA, catID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio SubCategorias) DeletarSubCategoria(catID uint64) error {
	statement, erro := repositorio.db.Prepare("UPDATE BDC_SUBCATEGORIAS SET STATUS = ? WHERE IDSUBCATEGORIA = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	STATUS := "INATIVO"

	if _, erro = statement.Exec(STATUS, catID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio SubCategorias) BuscarSubCategoria() ([]modelos.Post_SubCategoria, error) {

	linhas, erro := repositorio.db.Query(
		"SELECT IDSUBCATEGORIA, NOME, STATUS, ID_CATEGORIA FROM BDC_SUBCATEGORIAS WHERE STATUS = ? ORDER BY IDSUBCATEGORIA", "ATIVO",
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var SubCategoria []modelos.Post_SubCategoria

	for linhas.Next() {
		var subcategoria modelos.Post_SubCategoria

		if erro = linhas.Scan(
			&subcategoria.IDSUBCATEGORIA,
			&subcategoria.NOME,
			&subcategoria.STATUS,
			&subcategoria.ID_CATEGORIA,
		); erro != nil {
			return nil, erro
		}

		SubCategoria = append(SubCategoria, subcategoria)
	}

	return SubCategoria, nil

}

func (repositorio SubCategorias) BuscarSubCategoriaPorID(ID uint64) (modelos.Post_SubCategoria, error) {

	linhas, erro := repositorio.db.Query(
		"SELECT IDSUBCATEGORIA, NOME, STATUS, ID_CATEGORIA FROM BDC_SUBCATEGORIAS WHERE IDSUBCATEGORIA = ? AND STATUS = ?", ID, "ATIVO")
	if erro != nil {
		return modelos.Post_SubCategoria{}, erro
	}
	defer linhas.Close()

	var subcategoria modelos.Post_SubCategoria

	if linhas.Next() {
		if erro = linhas.Scan(
			&subcategoria.IDSUBCATEGORIA,
			&subcategoria.NOME,
			&subcategoria.STATUS,
			&subcategoria.ID_CATEGORIA,
		); erro != nil {
			return modelos.Post_SubCategoria{}, erro
		}
	}

	return subcategoria, nil

}

func (repositorio SubCategorias) BuscarSubCategoriaPorCategoria(ID uint64) ([]modelos.Post_SubCategoria, error) {

	linhas, erro := repositorio.db.Query(
		"SELECT IDSUBCATEGORIA, NOME, ID_CATEGORIA, STATUS FROM BDC_SUBCATEGORIAS WHERE ID_CATEGORIA = ? AND STATUS = ?", ID, "ATIVO")

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var SubCategoria []modelos.Post_SubCategoria

	for linhas.Next() {
		var subcategoria modelos.Post_SubCategoria

		if erro = linhas.Scan(
			&subcategoria.IDSUBCATEGORIA,
			&subcategoria.NOME,
			&subcategoria.ID_CATEGORIA,
			&subcategoria.STATUS,
		); erro != nil {
			return nil, erro
		}

		SubCategoria = append(SubCategoria, subcategoria)
	}

	return SubCategoria, nil

}
