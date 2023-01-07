package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Dacs struct {
	db *sql.DB
}

func NovoRepositorioDeDacs(db *sql.DB) *Dacs {
	return &Dacs{db}
}

func (repositorio Dacs) BuscarDacs() ([]modelos.Dac, error) {

	linhas, erro := repositorio.db.Query(
		"SELECT IDDAC, NOME FROM DACS ORDER BY IDDAC",
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var MapasDac []modelos.Dac

	for linhas.Next() {
		var mapa_dac modelos.Dac

		if erro = linhas.Scan(
			&mapa_dac.IDDAC,
			&mapa_dac.NOME,
		); erro != nil {
			return nil, erro
		}

		MapasDac = append(MapasDac, mapa_dac)
	}

	return MapasDac, nil

}
