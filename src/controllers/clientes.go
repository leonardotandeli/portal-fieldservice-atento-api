package controllers

import (
	"api/src/banco"
	"api/src/repositorios"
	"api/src/respostas"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// BuscarCliente traz um cliente armazenadas no banco de dados através do ID.
func BuscarCliente(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["clienteId"], 10, 64)
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

	repositorio := repositorios.NovoRepositorioDePosts(db)
	cliente, erro := repositorio.BuscarClientePorID(ID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, cliente)

}

// BuscarDadosCliente traz os clientes armazenados no banco de dados
func BuscarDadosCliente(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeMapasOperacional(db)
	mapa, erro := repositorio.BuscarClientes()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, mapa)

}
