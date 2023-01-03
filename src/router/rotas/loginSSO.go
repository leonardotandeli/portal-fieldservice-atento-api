package rotas

import (
	"api/src/controllers"
	"net/http"
)

// rotaLogin define a rota de login
var rotaLoginSSO = Rota{
	URI:                "/sso",
	Metodo:             http.MethodPost,
	Funcao:             controllers.LoginSSO,
	RequerAutenticacao: false,
}
