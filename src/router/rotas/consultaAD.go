package rotas

import (
	"api/src/controllers"
	"net/http"
)

// rotasClientes define todas as rotas dos clientes
var rotasConsultaAD = []Rota{

	{
		URI:                "/checkad/{login}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.ConsultaAD,
		RequerAutenticacao: true,
	},
	{
		URI:                "/checklaps/{locador}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.ConsultaLAPS,
		RequerAutenticacao: true,
	},
}
