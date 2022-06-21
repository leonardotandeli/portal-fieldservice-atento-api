package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar retorna um router e inicia as rotas de acesso
func Gerar() *mux.Router {
	r := mux.NewRouter()

	return rotas.Configurar(r)

}
