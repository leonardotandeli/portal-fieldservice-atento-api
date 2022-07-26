package middlewares

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
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
		//Valida token
		if erro := autenticacao.ValidarToken(r); erro != nil {

			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		//valida sessão no banco de dados
		if erro := autenticacao.SessionDB(r); erro != nil {

			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}

		proximaFuncao(w, r)
	}
}

//LoggerOnDb escreve informações da requisições no banco de dados.
func LoggerOnDb(w http.ResponseWriter, r *http.Request, ACTION string) {

	//logger
	var logs modelos.Logs
	logs.IDUSUARIO = autenticacao.ExtrairDadosUsuario(r).IDUSUARIO
	logs.LOGIN_NT = autenticacao.ExtrairDadosUsuario(r).LOGIN_NT
	logs.NOME = autenticacao.ExtrairDadosUsuario(r).NOME
	logs.DATA = time.Now()
	logs.ACTION = ACTION

	banco.DB.Create(&logs)

}
