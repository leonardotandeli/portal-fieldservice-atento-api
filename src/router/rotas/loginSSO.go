package rotas

import (
	"api/src/controllers"
	"net/http"
)

// rotaLoginSSO define a rota de login via single sign-on
var rotaLoginSSO = Rota{
	URI:                "/sso",
	Metodo:             http.MethodPost,
	Funcao:             controllers.LoginSSO,
	RequerAutenticacao: false,
}
