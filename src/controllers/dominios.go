package controllers

import (
	"api/src/banco"
	"api/src/repositorios"
	"api/src/respostas"
	"net/http"
)

// BuscarDadosDominios traz os dominios armazenados no banco de dados
func BuscarDadosDominios(w http.ResponseWriter, r *http.Request) {

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeDominios(db)
	mapa, erro := repositorio.BuscarDominios()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, mapa)

}
