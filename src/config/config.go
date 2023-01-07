package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConexaoBanco é a variável onde será iniciado a string de conexão com o MySQL
	StringConexaoBanco = ""

	// Porta inicia a variável onde será indicado a porta HTTP que a API vai rodar
	Porta = 0

	// SecretKey é a chave que vai ser usada para assinar o token
	SecretKey []byte

	// SSOPASS é utilizada para autenticar o usuário via SSO
	SSOPASS = ""
)

// Carregar vai inicializar buscar as variáveis de ambiente informadas no arquivo (.env)
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORTA"),
		os.Getenv("DB_NOME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
	SSOPASS = os.Getenv("SSO_PASS")
}
