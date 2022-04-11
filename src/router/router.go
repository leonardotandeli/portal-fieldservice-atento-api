package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Função Gerar retorna um router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()

	return rotas.Configurar(r)

}
