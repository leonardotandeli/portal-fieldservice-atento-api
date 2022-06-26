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
		Funcao:             controllers.CriarOperacaoMapa,
		RequerAutenticacao: true,
	},
	{
		URI:                "/mapa/operacoes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarOperacoesMapa,
		RequerAutenticacao: true,
	},
	{
		URI:                "/mapa/operacoes/{mapaId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarOperacaoMapa,
		RequerAutenticacao: true,
	},
	{
		URI:                "/mapa/operacoes/{mapaId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarOperacaoMapa,
		RequerAutenticacao: true,
	},
	{
		URI:                "/mapa/operacoes/{mapaId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarOperacaoMapa,
		RequerAutenticacao: true,
	},
	{
		URI:                "/mapa/busca",
		Metodo:             http.MethodGet,
		Funcao:             controllers.SearchOperacoesMapa,
		RequerAutenticacao: true,
	},
}
