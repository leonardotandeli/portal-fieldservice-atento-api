package rotas

import (
	"api/src/controllers"
	"net/http"
)

// rotaSites define todas as rotas dos sites
var rotaSites = []Rota{
	{
		URI:                "/sites",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarSites,
		RequerAutenticacao: true,
	},
}
