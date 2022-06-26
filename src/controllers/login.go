package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// Login é responsável por autenticar um usuário na API
func Login(w http.ResponseWriter, r *http.Request) {

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
	usuarioSalvoNoBanco, erro := repositorio.BuscarPorLogin(usuario.LOGIN_NT)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = seguranca.VerificarSenha(usuarioSalvoNoBanco.SENHA, usuario.SENHA); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := autenticacao.CriarToken(usuarioSalvoNoBanco.IDUSUARIO, usuarioSalvoNoBanco.LOGIN_NT, usuarioSalvoNoBanco.NOME)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	fmt.Println(usuarioSalvoNoBanco.LOGIN_NT)

	//logger db
	var logs modelos.Logs
	logs.Usuario.IDUSUARIO = usuarioSalvoNoBanco.IDUSUARIO
	logs.Usuario.LOGIN_NT = usuarioSalvoNoBanco.LOGIN_NT
	logs.Usuario.NOME = usuarioSalvoNoBanco.NOME
	logs.DATA = time.Now()
	logs.ACTION = "Login Efetuado"

	repositorioLogs := repositorios.NovoRepositorioDeLogs(db)
	logs.Usuario.IDUSUARIO, erro = repositorioLogs.LoggerDB(logs)
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

	respostas.JSON(w, http.StatusOK, modelos.DadosAutenticacao{ID: usuarioID, LOGIN_NT: login_nt, RE: re, CARGO: cargo, NOME: nome, EMAIL: email, Site: site, V_USUARIOS: v_usuarios, V_BDC_POSTS: v_bdc_posts, V_BDC_ADM: v_bdc_adm, V_IMDB: v_imdb, V_GSA: v_gsa, V_MAPA_OPERACIONAL: v_mapa_operacional, V_MAPA_OPERACIONAL_ADM: v_mapa_operacional_adm, Token: token})
}
