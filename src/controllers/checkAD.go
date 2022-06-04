package controllers

import (
	"api/src/modelos"
	"api/src/respostas"
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gorilla/mux"
)

func CheckAD(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID := parametros["login"]

	fmt.Println("retorno ok", ID)

	prg := "net"

	prg1 := "user"

	arg2 := "/domain"

	cmd := exec.Command(prg, prg1, ID, arg2)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user := string(stdout)
	nomeUsuario := strings.Index(user, "rio")
	nomeUsuario2 := strings.Index(user, `Nome completo`)
	username1 := string(stdout)[nomeUsuario+3 : nomeUsuario2]
	usernameCorreto := strings.TrimSpace(username1)

	fmt.Println(usernameCorreto)

	nomeOperador := strings.Index(user, "Nome completo")
	nomeOperador2 := strings.Index(user, `Comen`)
	nome1 := string(stdout)[nomeOperador+13 : nomeOperador2]
	nomeCorreto := strings.TrimSpace(nome1)

	fmt.Println(usernameCorreto)

	gpoUsuario := strings.Index(user, "Grupo Global")
	gpoUsuario2 := strings.Index(user, `Comando concl`)
	gpo1 := string(stdout)[gpoUsuario:gpoUsuario2]
	gpoCorreto := strings.TrimSpace(gpo1)

	//fmt.Print(string(stdout)[122:133])

	//fmt.Println(string(stdout))

	//	fmt.Println(string(stdout))

	i := strings.Contains(user, "Nome")
	fmt.Println(i)
	//	input := string(stdout)

	//	input, _ = string.ReadString('\n')

	respostas.JSON(w, http.StatusOK, modelos.DadosAD{LOGIN_NT: usernameCorreto, NOME: nomeCorreto, GPOS: gpoCorreto})

}
