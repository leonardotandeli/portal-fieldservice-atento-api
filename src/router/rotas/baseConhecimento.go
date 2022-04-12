package rotas

import (
	"api/src/controllers"
	"net/http"
)

// Define todas as rotas da base de conhecimento
var rotasBaseConhecimento = []Rota{
	{
		URI:                "/base",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarPost,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPosts,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base/busca",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPostsPorString,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base/{postId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPost,
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
		URI:                "/categorias",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarDadosCategorias,
		RequerAutenticacao: true,
	},
}
