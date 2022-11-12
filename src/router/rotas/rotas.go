package rotas

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar inicia todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)
	rotas = append(rotas, rotasBaseConhecimento...)
	rotas = append(rotas, rotasMapaOperacional...)
	rotas = append(rotas, rotasClientes...)
	rotas = append(rotas, rotasDACS...)
	rotas = append(rotas, rotasDominios...)
	rotas = append(rotas, rotaSites...)
	rotas = append(rotas, rotasCategoriasBaseConhecimento...)
	rotas = append(rotas, rotasSubCategoriasBaseConhecimento...)

	for _, rota := range rotas {

		if rota.RequerAutenticacao {
			r.HandleFunc(rota.URI,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao)),
			).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}

	}

	return r
}
