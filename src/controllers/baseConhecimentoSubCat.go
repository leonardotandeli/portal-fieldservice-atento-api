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
func CriarSubCategoriaBase(w http.ResponseWriter, r *http.Request) {

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var cat modelos.Post_SubCategoria
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

	repositorio := repositorios.NovoRepositorioDeSubCategorias(db)
	cat.IDSUBCATEGORIA, erro = repositorio.CriarSubCategoria(cat)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	//logger db
	middlewares.LoggerOnDb(w, r, "Categoria criada na base de conhecimento: ["+"IDPost: "+strconv.Itoa(int(cat.IDSUBCATEGORIA))+" | "+"Nome: "+cat.NOME+"]")

	respostas.JSON(w, http.StatusCreated, cat)
}

// AtualizarCategoria altera os dados de uma categoria
func AtualizarSubCategoria(w http.ResponseWriter, r *http.Request) {
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

	repositorio := repositorios.NovoRepositorioDeSubCategorias(db)
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var cat modelos.Post_SubCategoria
	if erro = json.Unmarshal(corpoRequisicao, &cat); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorio.AtualizarSubCategoria(catID, cat); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	//logger db
	middlewares.LoggerOnDb(w, r, "Categoria atualizada na base de conhecimento: ["+"IDPost: "+strconv.Itoa(int(catID))+" | "+"Nome: "+cat.NOME+"]")

	respostas.JSON(w, http.StatusNoContent, nil)
}

// DeletarCategoria exclui uma categoria
func DeletarSubCategoria(w http.ResponseWriter, r *http.Request) {

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

	repositorio := repositorios.NovoRepositorioDeSubCategorias(db)
	if erro = repositorio.DeletarSubCategoria(catID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	//logger db
	middlewares.LoggerOnDb(w, r, "Categoria deletada na base de conhecimento: ["+"IDPost: "+strconv.Itoa(int(catID))+"]")

	respostas.JSON(w, http.StatusNoContent, nil)
}

// BuscarTodasCategorias traz todas as categorias armazenadas no banco de dados
func BuscarTodasSubCategorias(w http.ResponseWriter, r *http.Request) {

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeSubCategorias(db)
	categoria, erro := repositorio.BuscarSubCategoria()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, categoria)

}

// BuscarCategoria traz uma categoria armazenadas no banco de dados
func BuscarSubCategoria(w http.ResponseWriter, r *http.Request) {
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

	repositorio := repositorios.NovoRepositorioDeSubCategorias(db)
	categoria, erro := repositorio.BuscarSubCategoriaPorID(ID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, categoria)

}

// BuscarCategoria traz uma categoria armazenadas no banco de dados
func BuscarSubCategoriaPorCategoria(w http.ResponseWriter, r *http.Request) {
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

	repositorio := repositorios.NovoRepositorioDeSubCategorias(db)
	subcategoria, erro := repositorio.BuscarSubCategoriaPorCategoria(ID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, subcategoria)

}
