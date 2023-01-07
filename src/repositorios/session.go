package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Sessions struct {
	db *sql.DB
}

func NovoRepositorioDeSessions(db *sql.DB) *Sessions {
	return &Sessions{db}
}

func (repositorio Sessions) BuscarPorID(ID uint64) (modelos.Session, error) {

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

func (repositorio Sessions) SessionCreate(session modelos.Session) (uint64, error) {
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

func (repositorio Sessions) DeletarSession(ID uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE FROM SESSIONS WHERE ID = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Sessions) CronDeletarSessionApos12Horas() error {

	statement, erro := repositorio.db.Prepare("DELETE FROM SESSIONS WHERE DATA < (NOW() - INTERVAL 12 HOUR)")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Sessions) DeletarSessionByUserID(ID_USUARIO uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE FROM SESSIONS WHERE ID_USUARIO = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID_USUARIO); erro != nil {
		return erro
	}

	return nil
}
