package rotas

import (
	"api/src/controllers"
	"net/http"
)

// rotasDominios define todas as rotas dos domínios
var rotasDominios = []Rota{

	{
		URI:                "/dominios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarDadosDominios,
		RequerAutenticacao: true,
	},
}
