package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Dominios representa um repositório de Dominios
type Dominios struct {
	db *sql.DB
}

// NovoRepositorioDeDominios inicia o repositório de Dominios
func NovoRepositorioDeDominios(db *sql.DB) *Dominios {
	return &Dominios{db}
}

// BuscarDominios traz todos os dominios do banco de dados
func (repositorio Dominios) BuscarDominios() ([]modelos.Dominio, error) {

	linhas, erro := repositorio.db.Query(
		"SELECT IDDOMINIO, NOME FROM DOMINIOS ORDER BY IDDOMINIO",
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var MapasDominio []modelos.Dominio

	for linhas.Next() {
		var mapa_dominio modelos.Dominio

		if erro = linhas.Scan(
			&mapa_dominio.IDDOMINIO,
			&mapa_dominio.NOME,
		); erro != nil {
			return nil, erro
		}

		MapasDominio = append(MapasDominio, mapa_dominio)
	}

	return MapasDominio, nil

}
