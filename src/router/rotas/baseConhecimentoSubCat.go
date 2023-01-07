package rotas

import (
	"api/src/controllers"
	"net/http"
)

// rotasSubCategoriasBaseConhecimento define todas as rotas de sub-categorias da base de conhecimento
var rotasSubCategoriasBaseConhecimento = []Rota{
	{
		URI:                "/subcategorias",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarSubCategoriaBase,
		RequerAutenticacao: true,
	},
	{
		URI:                "/subcategorias/{catId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarSubCategoria,
		RequerAutenticacao: true,
	},
	{
		URI:                "/subcategorias/{catId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarSubCategoria,
		RequerAutenticacao: true,
	},
	{
		URI:                "/subcategorias",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarTodasSubCategorias,
		RequerAutenticacao: true,
	},
	{
		URI:                "/subcategorias/{catId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarSubCategoria,
		RequerAutenticacao: true,
	},
}
