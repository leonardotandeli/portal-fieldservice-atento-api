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

// CriarDadosMapa adiciona uma nova operacao no banco de dados
func CriarDadosMapa(w http.ResponseWriter, r *http.Request) {
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
	middlewares.LoggerOnDb(w, r, "Mapa operacional criado: "+strconv.Itoa(int(mapa.IDMAPA))+" "+mapa.Cliente.NOME)

	respostas.JSON(w, http.StatusCreated, mapa)
}

// BuscarDadosMapa traz as operacoes armazenadas no banco de dados
func BuscarDadosMapa(w http.ResponseWriter, r *http.Request) {

	urlSite := strings.ToLower(r.URL.Query().Get("site"))
	urlCliente := strings.ToLower(r.URL.Query().Get("cliente"))
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeMapasOperacional(db)
	mapa, erro := repositorio.Buscar(urlSite, urlCliente)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, mapa)

}

// BuscarDadosMapaString traz as operacoes armazenadas no banco de dados através de parametros informados na URL
func BuscarDadosMapaString(w http.ResponseWriter, r *http.Request) {

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

// BuscarDadosSite traz os sites armazenados no banco de dados
func BuscarDadosSite(w http.ResponseWriter, r *http.Request) {

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeMapasOperacional(db)
	mapa, erro := repositorio.BuscarSites()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, mapa)

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

// BuscarDadosDacs traz os dacs armazenados no banco de dados
func BuscarDadosDacs(w http.ResponseWriter, r *http.Request) {

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeMapasOperacional(db)
	mapa, erro := repositorio.BuscarDacs()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, mapa)

}

// BuscarDadosDominios traz os dominios armazenados no banco de dados
func BuscarDadosDominios(w http.ResponseWriter, r *http.Request) {

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeMapasOperacional(db)
	mapa, erro := repositorio.BuscarDominios()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, mapa)

}

// BuscarDadoMapa retorna uma única publicação
func BuscarDadoMapa(w http.ResponseWriter, r *http.Request) {
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

// AtualizarDadosMapa altera os dados de uma operação
func AtualizarDadosMapa(w http.ResponseWriter, r *http.Request) {
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
	middlewares.LoggerOnDb(w, r, "Mapa operacional atualizado: "+strconv.Itoa(int(mapa.IDMAPA))+" "+mapa.Cliente.NOME)

	respostas.JSON(w, http.StatusNoContent, nil)
}

// DeletarDadosMapa exclui os dados de uma operação
func DeletarDadosMapa(w http.ResponseWriter, r *http.Request) {

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
	middlewares.LoggerOnDb(w, r, "Mapa operacional excluido: "+strconv.Itoa(int(mapaID)))

	respostas.JSON(w, http.StatusNoContent, nil)
}
