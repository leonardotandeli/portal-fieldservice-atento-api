package rotas

import (
	"api/src/controllers"
	"net/http"
)

// rotasBaseConhecimento define todas as rotas da base de conhecimento
var rotasBaseConhecimento = []Rota{
	{
		URI:                "/base",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarPost,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base/{postId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarPost,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base/{postId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarPost,
		RequerAutenticacao: true,
	},

	{
		URI:                "/base",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarTodosPosts,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base/busca",
		Metodo:             http.MethodGet,
		Funcao:             controllers.SearchPostsPorString,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base/busca-cat-cliente",
		Metodo:             http.MethodGet,
		Funcao:             controllers.SearchPostsPorCategoria,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base/busca-subcat-cliente",
		Metodo:             http.MethodGet,
		Funcao:             controllers.SearchPostsPorSubCategoria,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base/search",
		Metodo:             http.MethodGet,
		Funcao:             controllers.SearchBox,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base/{postId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPost,
		RequerAutenticacao: true,
	},
}
