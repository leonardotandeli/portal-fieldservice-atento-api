package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/respostas"
	"net/http"
)

// BuscarTodosDominios traz os dominios armazenados no banco de dados
func BuscarTodosDominios(w http.ResponseWriter, r *http.Request) {

	var dominios []modelos.Dominio
	banco.DB.Find(&dominios)

	respostas.JSON(w, http.StatusOK, dominios)
}
