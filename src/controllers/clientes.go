package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/respostas"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// BuscarUmCliente retorna os dados de um cliente através do ID
func BuscarUmCliente(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["clienteId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	var cliente modelos.Cliente
	banco.DB.First(&cliente, ID)

	respostas.JSON(w, http.StatusOK, cliente)

}

// BuscarTodosClientes traz os clientes armazenados no banco de dados
func BuscarTodosClientes(w http.ResponseWriter, r *http.Request) {

	var clientes []modelos.Cliente
	banco.DB.Find(&clientes)

	respostas.JSON(w, http.StatusOK, clientes)

}
