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

// CriarOperacaoMapa adiciona uma nova operacao no banco de dados
func CriarOperacaoMapa(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var mapa modelos.MapaOperacional
	if erro = json.Unmarshal(corpoRequest, &mapa); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeMapasOperacional(db)
	mapa.IDMAPA, erro = repositorio.CriarDadosMapa(mapa)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	//logger db
	middlewares.LoggerOnDb(w, r, "Operação criada no mapa operacional: ["+"IDOperacao: "+strconv.Itoa(int(mapa.IDMAPA))+" | "+"IDCliente: "+mapa.ID_CLIENTE+" | "+"Operação: "+mapa.OPERACAO+"]")

	respostas.JSON(w, http.StatusCreated, mapa)
}

// AtualizarOperacaoMapa altera os dados de uma operação
func AtualizarOperacaoMapa(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	mapaID, erro := strconv.ParseUint(parametros["mapaId"], 10, 64)
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

	repositorio := repositorios.NovoRepositorioDeMapasOperacional(db)

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var mapa modelos.MapaOperacional
	if erro = json.Unmarshal(corpoRequisicao, &mapa); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorio.Atualizar(mapaID, mapa); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	//logger db
	middlewares.LoggerOnDb(w, r, "Operação atualizada no mapa operacional: ["+"IDOperacao: "+strconv.Itoa(int(mapaID))+" | "+"IDCliente: "+mapa.ID_CLIENTE+" | "+"Operação: "+mapa.OPERACAO+"]")

	respostas.JSON(w, http.StatusNoContent, nil)
}

// DeletarOperacaoMapa exclui os dados de uma operação
func DeletarOperacaoMapa(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	mapaID, erro := strconv.ParseUint(parametros["mapaId"], 10, 64)
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

	repositorio := repositorios.NovoRepositorioDeMapasOperacional(db)
	if erro = repositorio.Deletar(mapaID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	//logger db
	middlewares.LoggerOnDb(w, r, "Operação deletada no mapa operacional: ["+"IDOperacao: "+strconv.Itoa(int(mapaID))+"]")

	respostas.JSON(w, http.StatusNoContent, nil)
}

// BuscarOperacoesMapa traz as operacoes armazenadas no banco de dados
func BuscarOperacoesMapa(w http.ResponseWriter, r *http.Request) {

	urlSite := strings.ToLower(r.URL.Query().Get("site"))
	urlCliente := strings.ToLower(r.URL.Query().Get("cliente"))

	urlPage := strings.ToLower(r.URL.Query().Get("pagina"))

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeMapasOperacional(db)
	mapa, erro := repositorio.Buscar(urlSite, urlCliente, urlPage)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, mapa)

}

// SearchOperacoesMapa traz as operacoes armazenadas no banco de dados através de parametros informados na URL
func SearchOperacoesMapa(w http.ResponseWriter, r *http.Request) {

	urlSite := strings.ToLower(r.URL.Query().Get("site"))
	urlCliente := strings.ToLower(r.URL.Query().Get("cliente"))
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeMapasOperacional(db)
	mapa, erro := repositorio.BuscarString(urlSite, urlCliente)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, mapa)

}

// BuscarOperacaoMapa retorna uma única publicação
func BuscarOperacaoMapa(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["mapaId"], 10, 64)
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

	repositorio := repositorios.NovoRepositorioDeMapasOperacional(db)
	post, erro := repositorio.BuscarPorID(ID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, post)
}
