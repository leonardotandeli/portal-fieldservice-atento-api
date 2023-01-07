package banco

import (
	"api/src/config"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // Driver
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//variaveis ORM
var (
	DB  *gorm.DB
	err error
)

// Conectar realiza a conexão com o banco de dados utilizando os dados informados no arquivo de variaveis de ambiente (.env)
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil

}

// ConectarComORM realiza a conexão com o banco de dados utilizando o GORM (ideia é migrar todas a comunicação com o banco via ORM no futuro)
func ConectarComORM() {
	DB, err = gorm.Open(mysql.Open(config.StringConexaoBanco))
	if err != nil {
		log.Panic("erro ao conectar no db")
	}
}
