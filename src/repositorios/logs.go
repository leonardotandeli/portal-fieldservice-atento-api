package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

// Logs representa um repositório de Logs
type Logs struct {
	db *sql.DB
}

// NovoRepositorioDeLogs inicia um repositório de Logs
func NovoRepositorioDeLogs(db *sql.DB) *Logs {
	return &Logs{db}
}

//LoggerOnLogin insere um novo registro no banco ao receber o login
func (repositorio Logs) LoggerDB(logs modelos.Logs) (uint64, error) {
	statment, erro := repositorio.db.Prepare(
		"INSERT INTO LOGS(IDUSUARIO, NOME, LOGIN_NT, ACTION, DATA) VALUES(?, ?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statment.Close()

	resultado, erro := statment.Exec(logs.Usuario.IDUSUARIO, logs.Usuario.NOME, logs.Usuario.LOGIN_NT, logs.ACTION, logs.DATA)
	if erro != nil {
		return 0, erro
	}
	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil

}

//LoggerOnLogin insere um novo registro no banco ao receber o login
func (repositorio Logs) SessionCreate(session modelos.Session) (uint64, error) {
	statment, erro := repositorio.db.Prepare(
		"INSERT INTO SESSIONS(ID_USUARIO, TOKEN, DATA) VALUES(?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statment.Close()

	resultado, erro := statment.Exec(session.Usuario.IDUSUARIO, session.DadosAutenticacao.Token, session.DATA)
	if erro != nil {
		return 0, erro
	}
	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil

}

// BuscarPorID traz um usuário do banco de dados filtrado pelo id
func (repositorio Logs) BuscarPorID(ID uint64) (modelos.Session, error) {
	fmt.Println("testeID")
	fmt.Println(ID)
	linhas, erro := repositorio.db.Query(
		"SELECT ID, ID_USUARIO, TOKEN, DATA FROM SESSIONS WHERE ID_USUARIO = ?", ID)
	if erro != nil {
		return modelos.Session{}, erro
	}
	defer linhas.Close()

	var session modelos.Session

	if linhas.Next() {
		if erro = linhas.Scan(
			&session.ID,
			&session.Usuario.IDUSUARIO,
			&session.DadosAutenticacao.Token,
			&session.DATA,
		); erro != nil {
			return modelos.Session{}, erro
		}
	}

	return session, nil
}
