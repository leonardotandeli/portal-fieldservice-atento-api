package main

import (
	"api/src/banco"
	"api/src/config"
	"api/src/controllers"
	"api/src/middlewares"
	"api/src/modelos"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func SetupRotasTeste() *mux.Router {
	r := mux.NewRouter()

	return r

}

func TestVerificaStatusCodeLogin(t *testing.T) {
	//inicializa arquivo de variaveis de ambiente
	config.Carregar()
	// inicializa conexão com o banco de dados
	banco.Conectar()
	// inicializa setup de rotas de testes
	r := SetupRotasTeste()

	// sobe a rota que será utilizada para teste
	r.HandleFunc("/login", middlewares.Logger(controllers.Login)).Methods(http.MethodOptions, http.MethodPost)

	// construção do body que será utilizado na requisição

	var m modelos.DadosAutenticacao

	m.LOGIN_NT = "adm"
	m.SENHA = "adm"

	login, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}

	//requisição para a rota definida anteriormente passando os dados de autenticação no body
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(login))

	// guarda a resposta
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	// realiza o teste e verifica se o status code está correto.
	if response.Code != http.StatusOK {
		t.Fatalf("Status error: valor recebido foi %d e o esperado é %d", response.Code, http.StatusOK)
	}
}
