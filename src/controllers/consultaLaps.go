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

// ConsultaAD faz a busca de informações do Active Directory no Powershell e retorna as informações para formatação e retorno da senha.
func ConsultaLAPS(w http.ResponseWriter, r *http.Request) {
	//parametros recebe dados através da url
	parametros := mux.Vars(r)
	locador := parametros["locador"]

	//fmt.Println(locador)

	//comando de de verificação de usuário no domínio.
	ps := "powershell.exe"
	cm := "-command"
	ad := "Get-ADComputer -Identity"
	ad2 := "-Properties *"

	//executa o comando net user no cmd
	cmd := exec.Command(ps, cm, ad, locador, ad2)
	//retorna os dados do cmd
	cmdOut, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	cmdReturn := string(cmdOut)

	//	fmt.Println(cmdReturn)

	//Define onde estará a informação
	var senhaWithoutSpace = ""
	if strings.Contains(cmdReturn, "ms-Mcs-AdmPwd") {
		SenhaSliceInitial := strings.Index(cmdReturn, "ms-Mcs-AdmPwd")
		SenhaSliceEnding := strings.Index(cmdReturn, `ms-Mcs-AdmPwdExpirationTime`)
		//fatia os dados com a informação recebida acima
		senha := string(cmdReturn)[SenhaSliceInitial+38 : SenhaSliceEnding]
		//Formatação - Remoção dos espaços
		senhaWithoutSpace = strings.TrimSpace(senha)
		//	fmt.Println(senhaWithoutSpace)
	} else {
		senhaWithoutSpace = "Nenhuma senha encontrada"
	}

	respostas.JSON(w, http.StatusOK, modelos.DadosLAPS{SENHA: senhaWithoutSpace})

}
