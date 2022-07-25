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
	"strings"

	"github.com/gorilla/mux"
)

// CriarPost adiciona uma novo post no banco de dados
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

	//logger db
	middlewares.LoggerOnDb(w, r, "Post criado: "+strconv.Itoa(int(post.IDPOST))+" "+post.TITULO)

	respostas.JSON(w, http.StatusCreated, post)
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

	//logger db
	middlewares.LoggerOnDb(w, r, "Post atualizado na base de conhecimento: ["+"IDPost: "+strconv.Itoa(int(postID))+" | "+"Titulo: "+post.TITULO+" | "+"IDCategoria: "+post.ID_CATEGORIA+" | "+"IDCliente: "+post.ID_CLIENTE+"]")

	respostas.JSON(w, http.StatusNoContent, nil)
}

// DeletarPost exclui um post no banco de dados
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

	//logger db
	middlewares.LoggerOnDb(w, r, "Post deletado na base de conhecimento: ["+"IDPost: "+strconv.Itoa(int(postID))+"]")

	respostas.JSON(w, http.StatusNoContent, nil)
}

// BuscarTodosPosts traz todas os posts existentes no banco de dados.
func BuscarTodosPosts(w http.ResponseWriter, r *http.Request) {

	//filtro de url para categorias e clientes
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

// BuscarPost retorna um único post
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

// Searchbox traz um post filtrando por uma palavra chave
func SearchBox(w http.ResponseWriter, r *http.Request) {

	//filtro por palavras chaves
	urlBusca := strings.ToLower(r.URL.Query().Get("busca"))
	println(urlBusca)
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePosts(db)
	posts, erro := repositorio.Busca(urlBusca)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, posts)

}

// BuscarPostsPorString traz os posts filtrando pelos parametros informados na URL (Necessário refatorar)
func SearchPostsPorString(w http.ResponseWriter, r *http.Request) {

	//filtro de url para categorias e clientes
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

// BuscarPostsPorStringCat traz os posts filtrando pelos parametros informados na URL (Necessário refatorar)
func SearchPostsPorCategoria(w http.ResponseWriter, r *http.Request) {
	//filtro de url para categorias e clientes
	urlCategoria := strings.ToLower(r.URL.Query().Get("categoria"))
	urlCliente := strings.ToLower(r.URL.Query().Get("cliente"))

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePosts(db)
	posts, erro := repositorio.BuscarPorStringCat(urlCategoria, urlCliente)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, posts)

}

// BuscarPostsPorStringCat traz os posts filtrando pelos parametros informados na URL (Necessário refatorar)
func SearchPostsPorSubCategoria(w http.ResponseWriter, r *http.Request) {
	//filtro de url para categorias e clientes
	urlCategoria := strings.ToLower(r.URL.Query().Get("subcategoria"))
	urlCliente := strings.ToLower(r.URL.Query().Get("cliente"))

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePosts(db)
	posts, erro := repositorio.BuscarPorStringSubCat(urlCategoria, urlCliente)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, posts)

}
