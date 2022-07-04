package controllers

import (
	"api/src/banco"
	"api/src/repositorios"
	"api/src/respostas"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// DeletarSessionDB deleta a sessão armazenada no banco
func DeletarSessionDB(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	userID, erro := strconv.ParseUint(parametros["userId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorioD := repositorios.NovoRepositorioDeSessions(db)
	sessionArmazenadaDB, erro := repositorioD.BuscarPorID(userID)
	if erro != nil {
		return
	}

	repositorio := repositorios.NovoRepositorioDeSessions(db)
	if erro = repositorio.DeletarSession(sessionArmazenadaDB.ID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
}
