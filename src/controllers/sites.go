package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/respostas"
	"net/http"
)

// BuscarSites traz os sites armazenados no banco de dados
func BuscarTodosSites(w http.ResponseWriter, r *http.Request) {

	var sites []modelos.Site
	banco.DB.Find(&sites)

	respostas.JSON(w, http.StatusOK, sites)

}
