package seguranca

import (
	"golang.org/x/crypto/bcrypt"
)

// Função que recebe uma string e transforma em um hash
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// Função que compara uma senha com o hash gerado e retorna sucesso se forem iguais
func VerificarSenha(senhaComHash, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
}
