package rotas

import (
	"api/src/controllers"
	"net/http"
)

// rotasBaseConhecimento define todas as rotas da base de conhecimento
var rotasCategoriasBaseConhecimento = []Rota{
	{
		URI:                "/categorias",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarCategoriaBase,
		RequerAutenticacao: true,
	},
	{
		URI:                "/categorias/{catId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarCategoria,
		RequerAutenticacao: true,
	},
	{
		URI:                "/categorias/{catId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarCategoria,
		RequerAutenticacao: true,
	},
	{
		URI:                "/categorias",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarTodasCategorias,
		RequerAutenticacao: true,
	},
	{
		URI:                "/categorias/{catId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarCategoria,
		RequerAutenticacao: true,
	},
}
