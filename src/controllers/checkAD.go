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
	//parametros recebe dados através da url
	parametros := mux.Vars(r)
	loginNT := parametros["login"]

	//comando de de verificação de usuário no domínio.
	net := "net"
	user := "user"
	domain := "/domain"

	//executa o comando net user no cmd
	cmd := exec.Command(net, user, loginNT, domain)
	//retorna os dados do cmd
	cmdOut, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	cmdReturn := string(cmdOut)

	//LOGIN NT
	//Define onde estará a informação
	LoginSliceInitial := strings.Index(cmdReturn, "rio")
	LoginSliceEnding := strings.Index(cmdReturn, `Nome completo`)
	//fatia os dados com a informação recebida acima
	logint := string(cmdReturn)[LoginSliceInitial+3 : LoginSliceEnding]
	//Formatação - Remoção dos espaços
	loginNTWithoutSpace := strings.TrimSpace(logint)

	//NOME COMPLETO
	//Define onde estará a informação
	NomeSliceInitial := strings.Index(cmdReturn, "Nome completo")
	NomeSliceEnding := strings.Index(cmdReturn, `Comen`)
	//fatia os dados com a informação recebida acima
	nome := string(cmdReturn)[NomeSliceInitial+13 : NomeSliceEnding]
	//Formatação - Remoção dos espaços
	nomeWithoutSpace := strings.TrimSpace(nome)

	//CONTA ATIVA
	//Define onde estará a informação
	ContaASliceInitial := strings.Index(cmdReturn, "Conta ativa")
	ContaASliceEnding := strings.Index(cmdReturn, `Conta expira`)
	//fatia os dados com a informação recebida acima
	contaAtiva := string(cmdReturn)[ContaASliceInitial+11 : ContaASliceEnding]
	//Formatação - Remoção dos espaços
	contaAtivaWithoutSpace := strings.TrimSpace(contaAtiva)

	//ULTIMA DEFINIÇÃO DE SENHA
	//Define onde estará a informação
	senhaDefinicaoSliceInitial := strings.Index(cmdReturn, "o de senha")
	senhaDefinicaoSliceEnding := strings.Index(cmdReturn, `A senha expira`)
	//fatia os dados com a informação recebida acima
	senhaDefinicao := string(cmdReturn)[senhaDefinicaoSliceInitial+10 : senhaDefinicaoSliceEnding]
	//Formatação - Remoção dos espaços
	senhaDefinicaoWithoutSpace := strings.TrimSpace(senhaDefinicao)

	//EXPIRACAO DE SENHA
	//Define onde estará a informação
	senhaExpiracaoSliceInitial := strings.Index(cmdReturn, "A senha expira")
	senhaExpiracaoSliceEnding := strings.Index(cmdReturn, `Altera`)
	//fatia os dados com a informação recebida acima
	senhaExpiracao := string(cmdReturn)[senhaExpiracaoSliceInitial+14 : senhaExpiracaoSliceEnding]
	//Formatação - Remoção dos espaços
	senhaExpiracaoWithoutSpace := strings.TrimSpace(senhaExpiracao)

	//DATA ULTIMO LOGON
	//Define onde estará a informação
	ultimoLogonSliceInitial := strings.Index(cmdReturn, "ltimo logon")
	ultimoLogonSliceEnding := strings.Index(cmdReturn, `Hor`)
	//fatia os dados com a informação recebida acima
	ultimoLogon := string(cmdReturn)[ultimoLogonSliceInitial+11 : ultimoLogonSliceEnding]
	//Formatação - Remoção dos espaços
	ultimoLogonWithoutSpace := strings.TrimSpace(ultimoLogon)

	//GPO
	//Define onde estará a informação
	gpoSliceInitial := strings.Index(cmdReturn, "es de Grupo Global")
	gpoSliceEnding := strings.Index(cmdReturn, `Comando conclu`)
	//fatia os dados com a informação recebida acima
	gpo := string(cmdReturn)[gpoSliceInitial+18 : gpoSliceEnding]
	//Formatação - Remoção dos espaços
	gpoFormat := strings.ReplaceAll(gpo, "*", "<br/>")

	gpoWithoutSpace := strings.TrimSpace(gpoFormat)

	respostas.JSON(w, http.StatusOK, modelos.DadosAD{LOGIN_NT: loginNTWithoutSpace, NOME: nomeWithoutSpace, CONTA_ATIVA: contaAtivaWithoutSpace, SENHA_ULTIMA_DEFINICAO: senhaDefinicaoWithoutSpace, SENHA_EXPIRACAO: senhaExpiracaoWithoutSpace, DATA_ULTIMO_LOGON: ultimoLogonWithoutSpace, GPO: gpoWithoutSpace})

}

func CheckLAPS(w http.ResponseWriter, r *http.Request) {
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
	if strings.Contains(cmdReturn, "ms-Mcs-AdmPwd") == true {
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
