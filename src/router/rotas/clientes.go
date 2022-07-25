package rotas

import (
	"api/src/controllers"
	"net/http"
)

// rotasClientes define todas as rotas dos clientes
var rotasClientes = []Rota{

	{
		URI:                "/clientes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarTodosClientes,
		RequerAutenticacao: true,
	},
	{
		URI:                "/clientes/{clienteId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUmCliente,
		RequerAutenticacao: true,
	},
}
