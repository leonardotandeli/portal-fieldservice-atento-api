package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Sessions representa um repositório de sessões
type Sessions struct {
	db *sql.DB
}

// NovoRepositorioDeSessions inicia um repositório de Sessions
func NovoRepositorioDeSessions(db *sql.DB) *Sessions {
	return &Sessions{db}
}

// BuscarPorID retorna os dados de uma sessão armazenada no banco de dados filtrando pelo id do usuário
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

//SessionCreate insere um novo registro de sessão no banco ao receber o login
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

// DeletarSession exclui uma sessão do banco de dados
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

// CronDeletarSessionApos12Horas deleta os registros que estão no banco há mais de 12 horas
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
