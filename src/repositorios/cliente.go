package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Clientes representa um repositório de Clientes
type Clientes struct {
	db *sql.DB
}

// NovoRepositorioDeClientes inicia o repositório de Clientes
func NovoRepositorioDeClientes(db *sql.DB) *Clientes {
	return &Clientes{db}
}

// BuscarClientes traz todos os clientes do banco de dados
func (repositorio Clientes) BuscarClientes() ([]modelos.Cliente, error) {

	linhas, erro := repositorio.db.Query(
		"SELECT IDCLIENTE, NOME, LOGO_URL FROM CLIENTES ORDER BY IDCLIENTE",
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var MapasCliente []modelos.Cliente

	for linhas.Next() {
		var mapa_cliente modelos.Cliente

		if erro = linhas.Scan(
			&mapa_cliente.IDCLIENTE,
			&mapa_cliente.NOME,
			&mapa_cliente.LOGO_URL,
		); erro != nil {
			return nil, erro
		}

		MapasCliente = append(MapasCliente, mapa_cliente)
	}

	return MapasCliente, nil

}

// BuscarClientePorID busca um chamado do banco de dados pelo id
func (repositorio Clientes) BuscarClientePorID(ID uint64) (modelos.Cliente, error) {

	linhas, erro := repositorio.db.Query(
		"SELECT IDCLIENTE, NOME FROM CLIENTES WHERE IDCLIENTE = ?", ID)
	if erro != nil {
		return modelos.Cliente{}, erro
	}
	defer linhas.Close()

	var cliente modelos.Cliente

	if linhas.Next() {
		if erro = linhas.Scan(
			&cliente.IDCLIENTE,
			&cliente.NOME,
		); erro != nil {
			return modelos.Cliente{}, erro
		}
	}

	return cliente, nil

}
