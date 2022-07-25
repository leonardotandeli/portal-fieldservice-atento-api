package banco

import (
	"api/src/config"
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // Driver
	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	erro error
)

// Conectar realiza a conexão com o banco de dados através do gorm
func ConectaComOBanco() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}
	user := os.Getenv("DB_USUARIO")
	password := os.Getenv("DB_SENHA")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORTA")
	dbname := os.Getenv("DB_NOME")

	stringDeConexao := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	DB, erro = gorm.Open(mysql.Open(stringDeConexao))
	if erro != nil {
		log.Panic("erro ao conectar no db")
	}

}

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
