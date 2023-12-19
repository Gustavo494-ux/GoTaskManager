package authentication

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

var (
	SecretKeyJWT = []byte("secretKey")
)

type TokenBuilder struct {
	Claims jwt.MapClaims
}

// NovoToken: cria uma nova instancia de autentição para criar o  token
func NovoToken(Autorizado bool, expiraEm int64) *TokenBuilder {
	return &TokenBuilder{
		Claims: jwt.MapClaims{
			"authorized": Autorizado,
			"exp":        expiraEm,
		},
	}
}

// AdicionarParametro: adiciona parametro ao token
func (tb *TokenBuilder) AdicionarParametro(key string, value interface{}) *TokenBuilder {
	tb.Claims[key] = value
	return tb
}

// Criar: cria o token a partir dos dados fornecidos anteriormente
func (tb *TokenBuilder) Criar() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tb.Claims)
	return token.SignedString(SecretKeyJWT)
}

// ValidarToken verifica se o token passado na requisição é válido
func ValidarToken(c http.Request) error {
	tokenString := ExtrairToken(c)
	_, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return erro
	}
	return nil
}

// ExtrairInformacao retorna a informação do token que corresponde à chave fornecida
func ExtrairInformacao(tokenString string, chave string) (string, error) {
	parametros, err := ExtrairTodasInformacoes(tokenString)
	if err != nil {
		return "", err
	}
	return parametros[chave], nil
}

// ExtrairTodasInformacoes retorna todas as informações do token
func ExtrairTodasInformacoes(tokenString string) (map[string]string, error) {
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return nil, erro
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claimsMap := make(map[string]string)
		for key, value := range claims {
			claimsMap[key] = fmt.Sprintf("%v", value)
		}
		return claimsMap, nil
	}

	return nil, errors.New("token inválido")
}

// ExtrairToken: Extrai o Token da requisição
func ExtrairToken(c http.Request) string {
	token := c.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado! %v", token.Header["alg"])
	}

	return SecretKeyJWT, nil
}
