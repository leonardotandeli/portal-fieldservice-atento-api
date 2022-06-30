package middlewares

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"log"
	"net/http"
	"time"
)

// Logger escreve informações da requisição no terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

// Autenticar verifica se o usuário fazendo a requisição está autenticado
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if erro := autenticacao.ValidarToken(r); erro != nil {

			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}

		if erro := autenticacao.SessionDB(r); erro != nil {

			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}

		proximaFuncao(w, r)
	}
}

//LoggerOnDb escreve informações da requisições POST, PUT e DELETE no banco de dados.
func LoggerOnDb(w http.ResponseWriter, r *http.Request, ACTION string) {

	//logger
	var logs modelos.Logs
	logs.Usuario.IDUSUARIO = autenticacao.ExtrairDadosUsuario(r).Usuario.IDUSUARIO
	logs.Usuario.LOGIN_NT = autenticacao.ExtrairDadosUsuario(r).Usuario.LOGIN_NT
	logs.Usuario.NOME = autenticacao.ExtrairDadosUsuario(r).Usuario.NOME
	logs.DATA = time.Now()
	logs.ACTION = ACTION
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repositorioLogs := repositorios.NovoRepositorioDeLogs(db)
	logs.Usuario.IDUSUARIO, erro = repositorioLogs.LoggerDB(logs)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

}
