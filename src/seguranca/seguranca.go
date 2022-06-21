package seguranca

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash recebe uma string e transforma em um hash
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerificarSenha compara a senha em formato string com o hash amarzenado em banco de dados e retorna sucesso se forem iguais
func VerificarSenha(senhaComHash, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
}
