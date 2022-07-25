package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/respostas"
	"net/http"
)

// BuscarTodosDacs retorna todos os dacs armazenados no banco de dados
func BuscarTodosDacs(w http.ResponseWriter, r *http.Request) {

	var dacs []modelos.Dac
	banco.DB.Find(&dacs)

	respostas.JSON(w, http.StatusOK, dacs)

}
