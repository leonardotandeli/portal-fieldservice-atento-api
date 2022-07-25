package main

import (
	"api/src/banco"
	"api/src/config"
	"api/src/cron"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

// Função init gera a secret key do token
/*
func init() {
	secret_key := make([]byte, 64)
	if _, err := rand.Read(secret_key); err != nil {
		log.Fatal(err)
	}
	string64 := base64.StdEncoding.EncodeToString(secret_key)
	fmt.Println(string64)
}
*/

func main() {

	//carrega o banco através do gorm
	banco.ConectaComOBanco()
	// carrega variaveis de ambiente
	config.Carregar()

	// gera as rotas http
	r := router.Gerar()

	// define autorização de acesso a partir da url do front-end
	c := cors.New(cors.Options{AllowedOrigins: []string{os.Getenv("FRONTEND_URL")}})

	// handler de rotas
	handler := c.Handler(r)

	// imprime informações de ambiente no console
	fmt.Printf("API Executando na porta %d\n", config.Porta)
	fmt.Printf("URL definida como Front-end: %s\n", os.Getenv("FRONTEND_URL"))
	fmt.Printf("Nome do Banco de dados: %s\n", os.Getenv("DB_NOME"))

	//inicia as tarefas
	go cron.Tarefas()

	// inicia o servidor http
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), handler))
}
