package rotas

import (
	"api/src/controllers"
	"net/http"
)

// rotaLogin define a rota de login
var rotaLogin = Rota{
	URI:                "/login",
	Metodo:             http.MethodPost,
	Funcao:             controllers.Login,
	RequerAutenticacao: false,
}
