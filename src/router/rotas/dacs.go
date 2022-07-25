package rotas

import (
	"api/src/controllers"
	"net/http"
)

// rotasDACS define todas as rotas de DACS
var rotasDACS = []Rota{

	{
		URI:                "/dacs",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarTodosDacs,
		RequerAutenticacao: true,
	},
}
