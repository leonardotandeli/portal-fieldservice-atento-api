package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

//  representa um repositório de Sites
type Sites struct {
	db *sql.DB
}

// NovoRepositorioDeSites inicia o repositório de Sites
func NovoRepositorioDeSites(db *sql.DB) *Sites {
	return &Sites{db}
}

// BuscarSites traz todos os sites do banco de dados
func (repositorio Sites) BuscarSites() ([]modelos.Site, error) {

	linhas, erro := repositorio.db.Query(
		"SELECT IDSITE, NOME, SIGLA, UF, ENDERECO FROM SITES ORDER BY IDSITE",
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var MapasSite []modelos.Site

	for linhas.Next() {
		var mapa_site modelos.Site

		if erro = linhas.Scan(
			&mapa_site.IDSITE,
			&mapa_site.NOME,
			&mapa_site.SIGLA,
			&mapa_site.UF,
			&mapa_site.ENDERECO,
		); erro != nil {
			return nil, erro
		}

		MapasSite = append(MapasSite, mapa_site)
	}

	return MapasSite, nil

}
