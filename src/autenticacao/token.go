package autenticacao

import (
	"api/src/banco"
	"api/src/config"
	"api/src/repositorios"

	"api/src/modelos"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

//var hmacSampleSecret []byte

// CriarToken retorna um token assinado com os dados do usuário
func CriarToken(usuarioID uint64, LOGIN_NT string, NOME string) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 12).Unix() //numero de horas para expirar o token e solicitar novo login
	permissoes["usuarioId"] = usuarioID
	permissoes["login_nt"] = LOGIN_NT
	permissoes["nome"] = NOME
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte(config.SecretKey))
}

// ValidarToken verifica se o token passado na requisição é valido
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)

	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("TOKEN NÃO ESTÁ VÁLIDO")
}

//SessionDB escreve informações da sessão no banco de dados.
func SessionDB(r *http.Request) error {
	var session modelos.Session
	session.Usuario.IDUSUARIO = ExtrairDadosUsuario(r).Usuario.IDUSUARIO

	db, erro := banco.Conectar()
	if erro != nil {
		return erro
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeSessions(db)
	sessionArmazenadaDB, erro := repositorio.BuscarPorID(session.Usuario.IDUSUARIO)
	if erro != nil {
		return erro
	}

	if session.Usuario.IDUSUARIO == sessionArmazenadaDB.Usuario.IDUSUARIO {
		return erro
	}
	return errors.New("TOKEN NÃO ESTÁ VÁLIDO")
}

// ExtrairUsuarioID extrai o dados do usuário e valida o token.
func ExtrairUsuarioID(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return 0, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}

		return usuarioID, nil
	}

	return 0, errors.New("TOKEN NÃO ESTÁ VÁLIDO")
}

// extrairToken extrai o dados do token.
func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

// retornarChaveDeVerificacao retorna a chave de verificação do token
func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("MÉTODO DE ASSINATURA INESPERADO. %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}

// ExtrairUsuarioUsuario extrai o dados do usuário e valida o token.
func ExtrairDadosUsuario(r *http.Request) (user modelos.Logs) {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return user
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user.Usuario.IDUSUARIO, erro = strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioId"]), 10, 64)
		if erro != nil {
			return user
		}
		user.Usuario.LOGIN_NT = fmt.Sprintf("%s", permissoes["login_nt"])
		user.Usuario.NOME = fmt.Sprintf("%s", permissoes["nome"])

		return user
	}

	return user
}
