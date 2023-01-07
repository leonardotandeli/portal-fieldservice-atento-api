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

func main() {
	// carrega variaveis de ambiente
	config.Carregar()

	// inicia conexão com o banco via ORM
	banco.ConectarComORM()

	// gera as rotas http
	r := router.Gerar()

	// define autorização de acesso a partir da url do front-end
	c := cors.New(cors.Options{AllowedOrigins: []string{os.Getenv("FRONTEND_URL")}})

	// handler de rotas
	handler := c.Handler(r)

	// imprime informações no console
	fmt.Printf("API Executando na porta %d\n", config.Porta)

	//inicia as tarefas
	go cron.Tarefas()

	// inicia o servidor http
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), handler))
}
