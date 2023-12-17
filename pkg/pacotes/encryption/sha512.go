package encryption

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
)

// GenerateSHA512: gera um hash usando o algoritmo Sha512 para os dados fornecidos
func GerarSHA512(dado string) (string, error) {
	h := sha512.New()
	_, err := h.Write([]byte(dado))
	if err != nil {
		return "", err
	}
	dadoHash := h.Sum(nil)
	hashHex := hex.EncodeToString(dadoHash)
	return hashHex, nil
}

// CompararSHA512: verifica se o texto fornecido corresponde ao hash
func CompararSHA512(dadoEncriptado, dadoTexto string) (err error) {
	decryptedHash, err := GerarSHA512(dadoTexto)
	if err != nil {
		return
	}
	dadoEncriptadoBytes := []byte(dadoEncriptado)

	if fmt.Sprintf("%x", decryptedHash) != fmt.Sprintf("%x", dadoEncriptadoBytes) {
		return errors.New("o dado fornecido não é compativel com o hash")
	}
	return
}
