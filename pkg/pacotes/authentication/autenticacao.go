package authentication

import (
	"GoTaskManager/pkg/pacotes/logger"
	tipo "GoTaskManager/pkg/pacotes/tipos"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	secretKeyJWT []byte
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
	token := jwt.New(jwt.SigningMethodHS256)

	permissoes := token.Claims.(jwt.MapClaims)
	for chave, valor := range tb.Claims {
		permissoes[chave] = valor
	}
	return token.SignedString(secretKeyJWT)
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
func ExtrairInformacao(tokenString string, chave string) (interface{}, error) {
	parametros, err := ExtrairTodasInformacoes(tokenString)
	if err != nil {
		return nil, err
	}

	return parametros[chave], nil
}

// ExtrairTodasInformacoes retorna todas as informações do token
func ExtrairTodasInformacoes(tokenString string) (map[string]interface{}, error) {
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		if erro.Error() == "Token is expired" {
			return nil, fmt.Errorf("o token expirou")
		}
		return nil, erro
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
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

	return secretKeyJWT, nil
}

// DefinirSecretKey: define a chave de segurança da autenticacao TokenJWT
func DefinirSecretKey(secretKey string) {
	if len(secretKey) > 0 {
		secretKeyJWT = []byte(secretKey)
	}
}

// RetornarSecretKey: retorna a chave de segurança da autenticacao TokenJWT
func RetornarSecretKey() string {
	return string(secretKeyJWT)
}

// IsExpirado: verifica se o token expirou
func IsExpirado(r http.Request) error {
	tokenInvalido := fmt.Errorf("o token expirou")

	if err := ValidarToken(r); err != nil {
		return tokenInvalido
	}

	token := ExtrairToken(r)

	expiracao, err := ExtrairInformacao(token, "exp")
	if err != nil {
		if err != tokenInvalido {
			logger.Logger().Error("Ocorreu um erro ao executar a função ExtrairInformacao", err, token)
		}
		return err
	}

	if tipo.Converter(expiracao).Int(10, 64) < time.Now().Unix() {
		return tokenInvalido
	}
	return nil
}
