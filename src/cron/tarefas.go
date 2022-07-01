package cron

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"fmt"
	"net/http"

	"github.com/jasonlvhit/gocron"
)

//SessionDB escreve informações no banco de dados.
func SessionGetAllDB(r *http.Request) error {

	db, erro := banco.Conectar()
	if erro != nil {
		return erro
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeSessions(db)
	if erro = repositorio.CronDeletarSessionApos12Horas(); erro != nil {
		return erro
	}

	var session modelos.Session
	sessionArmazenadaDB = append(sessionArmazenadaDB, session)

	//	d := sessionArmazenadaDB.DATA.Unix()      // 85
	//	d2 := sessionArmazenadaDB.DATA.Unix() + 5 // 90

	fmt.Println(sessionArmazenadaDB)

	//	if d >= d2 {
	//		fmt.Println("delete")
	//		fmt.Println(sessionArmazenadaDB.ID)
	//		repositorio := repositorios.NovoRepositorioDeLogs(db)
	//		repositorio.CronDeleteSession(sessionArmazenadaDB.ID)
	//		if erro != nil {
	//			return erro
	//		}
	//	}
	////////////////////////
	/*


		//if sessionArmazenadaDB.DATA > sessionArmazenadaDB.DATA {
		//		return errors.New("Token não está válido!")
		//	}
		return errors.New("Token não está válido!")
	*/

	return erro
}

func taskRemoveSession() {
	fmt.Println("Task is being performed.")
	SessionGetAllDB(&http.Request{})
}

func RemoveSession() {

	// Do jobs without params
	gocron.Every(2).Second().Do(taskRemoveSession)

	// Start all the pending jobs
	<-gocron.Start()
}
