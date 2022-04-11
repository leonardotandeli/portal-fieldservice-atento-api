package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CriarPost adiciona uma nova publicação no banco de dados
func CriarPost(w http.ResponseWriter, r *http.Request) {

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var post modelos.Post
	if erro = json.Unmarshal(corpoRequisicao, &post); erro != nil {
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
	post.IDPOST, erro = repositorio.Criar(post)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, post)
}

// BuscarPosts traz as publicações
func BuscarPosts(w http.ResponseWriter, r *http.Request) {

	urlCategoria := strings.ToLower(r.URL.Query().Get("categoria"))
	urlCliente := strings.ToLower(r.URL.Query().Get("cliente"))

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePosts(db)
	posts, erro := repositorio.BuscarTodos(urlCategoria, urlCliente)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, posts)

}

// BuscarPosts traz as publicações
func BuscarPostsPorString(w http.ResponseWriter, r *http.Request) {

	urlCategoria := strings.ToLower(r.URL.Query().Get("categoria"))
	urlCliente := strings.ToLower(r.URL.Query().Get("cliente"))

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePosts(db)
	posts, erro := repositorio.BuscarPorString(urlCategoria, urlCliente)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, posts)

}

// BuscarPost traz uma única publicação
func BuscarPost(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["postId"], 10, 64)
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
	post, erro := repositorio.BuscarPorID(ID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, post)
}

// AtualizarPost altera os dados de uma publicação
func AtualizarPost(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	postID, erro := strconv.ParseUint(parametros["postId"], 10, 64)
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
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var post modelos.Post
	if erro = json.Unmarshal(corpoRequisicao, &post); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorio.Atualizar(postID, post); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// DeletarPost exclui os dados de uma publicação
func DeletarPost(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	postID, erro := strconv.ParseUint(parametros["postId"], 10, 64)
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
	if erro = repositorio.Deletar(postID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// BuscarDadosCategorias traz as categorias armazenadas no banco de dados
func BuscarDadosCategorias(w http.ResponseWriter, r *http.Request) {

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePosts(db)
	categoria, erro := repositorio.BuscarCategoria()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, categoria)

}
