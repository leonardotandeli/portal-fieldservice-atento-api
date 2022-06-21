package rotas

import (
	"api/src/controllers"
	"net/http"
)

// rotasMapaOperacional define todas as rotas do mapa operacional
var rotasMapaOperacional = []Rota{
	{
		URI:                "/mapa/operacoes",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarDadosMapa,
		RequerAutenticacao: true,
	},
	{
		URI:                "/mapa/operacoes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarDadosMapa,
		RequerAutenticacao: true,
	},
	{
		URI:                "/mapa/busca",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarDadosMapaString,
		RequerAutenticacao: true,
	},
	{
		URI:                "/mapa/operacoes/{mapaId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarDadoMapa,
		RequerAutenticacao: true,
	},
	{
		URI:                "/mapa/operacoes/{mapaId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarDadosMapa,
		RequerAutenticacao: true,
	},
	{
		URI:                "/mapa/operacoes/{mapaId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarDadosMapa,
		RequerAutenticacao: true,
	},
	{
		URI:                "/sites",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarDadosSite,
		RequerAutenticacao: true,
	},
	{
		URI:                "/clientes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarDadosCliente,
		RequerAutenticacao: true,
	},
	{
		URI:                "/dacs",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarDadosDacs,
		RequerAutenticacao: true,
	},
	{
		URI:                "/dominios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarDadosDominios,
		RequerAutenticacao: true,
	},
}
