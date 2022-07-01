package cron

import (
	"api/src/banco"
	"api/src/repositorios"
	"net/http"

	"github.com/jasonlvhit/gocron"
)

//SessionDelete deleta as sessões com mais de 12 horas no banco de dados...
func SessionDelete(r *http.Request) error {

	db, erro := banco.Conectar()
	if erro != nil {
		return erro
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeSessions(db)
	if erro = repositorio.CronDeletarSessionApos12Horas(); erro != nil {
		return erro
	}

	return erro
}

//task chama a função que deleta as sessões do banco de dados.
func task() {
	//fmt.Println("Tarefa sendo executada")
	SessionDelete(&http.Request{})
}

//Tarefas executa as tarefas a cada 1 minuto
func Tarefas() {
	// executa as tarefas a cada 1 minuto
	gocron.Every(1).Minute().Do(task)

	// inicia todos as tarefas pendentes
	<-gocron.Start()
}
