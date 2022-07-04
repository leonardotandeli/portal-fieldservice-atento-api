package controllers

import (
	"api/src/banco"
	"api/src/middlewares"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CriarCategoriaBase adiciona uma nova categoria no banco de dados
func CriarCategoriaBase(w http.ResponseWriter, r *http.Request) {

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var cat modelos.Post_Categoria
	if erro = json.Unmarshal(corpoRequisicao, &cat); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeCategorias(db)
	cat.IDCATEGORIA, erro = repositorio.CriarCategoria(cat)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	//logger db
	middlewares.LoggerOnDb(w, r, "Categoria criada na base de conhecimento: ["+"IDPost: "+strconv.Itoa(int(cat.IDCATEGORIA))+" | "+"Nome: "+cat.NOME+"]")

	respostas.JSON(w, http.StatusCreated, cat)
}

// AtualizarCategoria altera os dados de uma categoria
func AtualizarCategoria(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	catID, erro := strconv.ParseUint(parametros["catId"], 10, 64)
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

	repositorio := repositorios.NovoRepositorioDeCategorias(db)
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var cat modelos.Post_Categoria
	if erro = json.Unmarshal(corpoRequisicao, &cat); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorio.AtualizarCategoria(catID, cat); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	//logger db
	middlewares.LoggerOnDb(w, r, "Categoria atualizada na base de conhecimento: ["+"IDPost: "+strconv.Itoa(int(catID))+" | "+"Nome: "+cat.NOME+"]")

	respostas.JSON(w, http.StatusNoContent, nil)
}

// DeletarCategoria exclui uma categoria
func DeletarCategoria(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	catID, erro := strconv.ParseUint(parametros["catId"], 10, 64)
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

	repositorio := repositorios.NovoRepositorioDeCategorias(db)
	if erro = repositorio.DeletarCategoria(catID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	//logger db
	middlewares.LoggerOnDb(w, r, "Categoria deletada na base de conhecimento: ["+"IDPost: "+strconv.Itoa(int(catID))+"]")

	respostas.JSON(w, http.StatusNoContent, nil)
}

// BuscarTodasCategorias traz todas as categorias armazenadas no banco de dados
func BuscarTodasCategorias(w http.ResponseWriter, r *http.Request) {

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeCategorias(db)
	categoria, erro := repositorio.BuscarCategoria()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, categoria)

}

// BuscarCategoria traz uma categoria armazenadas no banco de dados
func BuscarCategoria(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["catId"], 10, 64)
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

	repositorio := repositorios.NovoRepositorioDeCategorias(db)
	categoria, erro := repositorio.BuscarCategoriaPorID(ID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, categoria)

}
