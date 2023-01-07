package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Logs struct {
	db *sql.DB
}

func NovoRepositorioDeLogs(db *sql.DB) *Logs {
	return &Logs{db}
}

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
