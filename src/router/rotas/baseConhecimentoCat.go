package rotas

import (
	"api/src/controllers"
	"net/http"
)

// rotasCategoriasBaseConhecimento define todas as rotas de categorias da base de conhecimento
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
	{
		URI:                "/categorias/cliente/{clienteId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarCategoriaPorCliente,
		RequerAutenticacao: true,
	},
	{
		URI:                "/categorias/subcategoria/{catId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarSubCategoriaPorCategoria,
		RequerAutenticacao: true,
	},
}
