package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/respostas"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// DeletarSessionBanco deleta a sessão armazenada no banco
func DeletarSessionBanco(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	userID, erro := strconv.ParseUint(parametros["userId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	var session modelos.Session
	banco.DB.Where("ID_USUARIO = ?", userID).Delete(&session)

}
