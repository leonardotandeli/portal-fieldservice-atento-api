package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/config"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/securecookie"
)

// LoginSSO é responsável por autenticar um usuário na API utilizando as informações da AZURE
func LoginSSO(w http.ResponseWriter, r *http.Request) {

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioIDSalvoNoBanco, erro := repositorio.BuscarPorLogin(usuario.LOGIN_NT)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if usuario.SENHA != config.SSOPASS {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	var (
		nomeUser      = usuarioIDSalvoNoBanco.NOME
		cargoUser     = usuarioIDSalvoNoBanco.CARGO
		emailUser     = usuarioIDSalvoNoBanco.EMAIL
		reFinal       = usuarioIDSalvoNoBanco.RE
		idSite        = usuario.ID_SITE
		siteTemNumero = regexp.MustCompile(`\d`).MatchString(idSite)
		reUser        = usuario.LOGIN_NT
		reTemNumero   = regexp.MustCompile(`\d`).MatchString(reUser)
		reCheck       = strings.Contains(usuarioIDSalvoNoBanco.RE, "nao_informado_")
	)

	if usuarioIDSalvoNoBanco.NOME != usuario.NOME {
		//fmt.Println("nome armazenado no banco é diferente, realizando atualização...")
		nomeUser = usuario.NOME
	}

	if usuarioIDSalvoNoBanco.CARGO != usuario.CARGO {
		//fmt.Println("cargo armazenado no banco é diferente, realizando atualização...")
		cargoUser = usuario.CARGO
	}

	if usuarioIDSalvoNoBanco.EMAIL != usuario.EMAIL {
		//fmt.Println("email armazenado no banco é diferente, realizando atualização...")
		emailUser = usuario.EMAIL
	}

	if reTemNumero == true { //checa se o login é com números (ab+re) ou se é login fora do padrão ab
		// se for login padrão ab+re remove o ab e o @br.atento.com para que sobre apenas o RE
		reUser = strings.ReplaceAll(reUser, "AB", "")
		reUser = strings.ReplaceAll(reUser, "@br.atento.com", "")
	} else { // se não for login padrão ab+re armazena informação aleatoria no banco.
		reUser = "nao_informado_" + hex.EncodeToString(securecookie.GenerateRandomKey(4))
	}

	if usuarioIDSalvoNoBanco.RE != reUser && reCheck == true {
		reFinal = reUser
	}

	if siteTemNumero == true {
		idSite = strings.TrimSpace(idSite[2:5])
		switch idSite {
		case "71": // zona leste
			usuario.ID_SITE = "33"
		case "60": // zona sul
			usuario.ID_SITE = "31"
		case "62": // uruguai
			usuario.ID_SITE = "30"
		case "04": // teleporto
			usuario.ID_SITE = "29"
		case "05": // sjc
			usuario.ID_SITE = "28"
		case "67": // scs
			usuario.ID_SITE = "27"
		case "06": // sbc
			usuario.ID_SITE = "26"
		case "19": // sao bento 1
			usuario.ID_SITE = "24"
		case "47": // sao bento 2
			usuario.ID_SITE = "25"
		case "37": // santo antonio
			usuario.ID_SITE = "23"
		case "45": // santo andré
			usuario.ID_SITE = "22"
		case "40": // santana
			usuario.ID_SITE = "21"
		case "72": // nova sp 2
			usuario.ID_SITE = "20"
		case "01": // nova sp
			usuario.ID_SITE = "15"
		case "13": // republica
			usuario.ID_SITE = "19"
		case "03": // porto alegre
			usuario.ID_SITE = "18"
		case "57": // penha
			usuario.ID_SITE = "17"
		case "21": // oliveira coutinho
			usuario.ID_SITE = "16"
		case "50": // madureira
			usuario.ID_SITE = "14"
		case "32": // liberdade
			usuario.ID_SITE = "13"
		case "68": // gru
			usuario.ID_SITE = "12"
		case "15": // goiania
			usuario.ID_SITE = "11"
		case "64": // feira de santana
			usuario.ID_SITE = "10"
		case "53": // del castilho
			usuario.ID_SITE = "9"
		case "56": // casa
			usuario.ID_SITE = "7"
		case "52": // bh
			usuario.ID_SITE = "4"
		case "73": // sede
			usuario.ID_SITE = "3"
		default:
			usuario.ID_SITE = "1"
		}
		idSite = usuario.ID_SITE
	}

	banco.DB.First(&usuario, usuarioIDSalvoNoBanco.IDUSUARIO)
	usuario.NOME = nomeUser
	usuario.CARGO = cargoUser
	usuario.EMAIL = emailUser
	usuario.RE = reFinal
	usuario.ID_SITE = idSite
	banco.DB.Save(&usuario)

	usuarioSalvoNoBanco, erro := repositorio.BuscarPorLogin(usuario.LOGIN_NT)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	token, erro := autenticacao.CriarToken(usuarioSalvoNoBanco.IDUSUARIO, usuarioSalvoNoBanco.LOGIN_NT, usuarioSalvoNoBanco.NOME)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	//logger db - inclui um log no banco ao realizar o login
	var logs modelos.Logs
	logs.Usuario.IDUSUARIO = usuarioSalvoNoBanco.IDUSUARIO
	logs.Usuario.LOGIN_NT = usuarioSalvoNoBanco.LOGIN_NT
	logs.Usuario.NOME = usuarioSalvoNoBanco.NOME
	logs.DATA = time.Now()
	logs.ACTION = "Login Efetuado através do SSO AZURE"

	repositorioLogs := repositorios.NovoRepositorioDeLogs(db)
	logs.Usuario.IDUSUARIO, erro = repositorioLogs.LoggerDB(logs)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	// Session db - inclui uma sessão no banco ao realizar o login
	var session modelos.Session
	session.Usuario.IDUSUARIO = usuarioSalvoNoBanco.IDUSUARIO
	session.DadosAutenticacao.Token = token
	session.DATA = time.Now()

	repositorioSession := repositorios.NovoRepositorioDeSessions(db)
	logs.Usuario.IDUSUARIO, erro = repositorioSession.SessionCreate(session)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	// retorna dados do usuário como JSON
	usuarioID := strconv.FormatUint(usuarioSalvoNoBanco.IDUSUARIO, 10)
	login_nt := usuarioSalvoNoBanco.LOGIN_NT
	re := usuarioSalvoNoBanco.RE
	cargo := usuarioSalvoNoBanco.CARGO
	nome := usuarioSalvoNoBanco.NOME
	email := usuarioSalvoNoBanco.EMAIL
	site := usuarioSalvoNoBanco.Site
	v_usuarios := usuarioSalvoNoBanco.V_USUARIOS
	v_bdc_posts := usuarioSalvoNoBanco.V_BDC_POSTS
	v_bdc_adm := usuarioSalvoNoBanco.V_BDC_ADM
	v_imdb := usuarioSalvoNoBanco.V_IMDB
	v_gsa := usuarioSalvoNoBanco.V_GSA
	v_mapa_operacional := usuarioSalvoNoBanco.V_MAPA_OPERACIONAL
	v_mapa_operacional_adm := usuarioSalvoNoBanco.V_MAPA_OPERACIONAL_ADM
	status := usuarioSalvoNoBanco.STATUS

	respostas.JSON(w, http.StatusOK, modelos.DadosAutenticacao{ID: usuarioID, LOGIN_NT: login_nt, RE: re, CARGO: cargo, NOME: nome, EMAIL: email, Site: site, V_USUARIOS: v_usuarios, V_BDC_POSTS: v_bdc_posts, V_BDC_ADM: v_bdc_adm, V_IMDB: v_imdb, V_GSA: v_gsa, V_MAPA_OPERACIONAL: v_mapa_operacional, V_MAPA_OPERACIONAL_ADM: v_mapa_operacional_adm, STATUS: status, Token: token})
}
