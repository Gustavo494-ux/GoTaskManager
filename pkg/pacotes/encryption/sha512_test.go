package encryption

import (
	"GoTaskManager/internal/app/inicializar"
	"GoTaskManager/pkg/pacotes/logger"
	"os"
	"testing"
)

// TestMain:Função executada antes das demais
func TestMain(m *testing.M) {
	inicializar.InicializarParaTestes()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestGerarSHA512(t *testing.T) {
	data := "password123"
	hash, err := GerarSHA512(data)
	if err != nil {
		t.Errorf("Erro ao gerar a hash Sha512: %v", err)
		logger.Logger().Error("Erro ao gerar a hash Sha512", err)
	}

	if hash == "" {
		t.Error("O hash gerado não deve estar vazio")
		logger.Logger().Error("O hash gerado não deve estar vazio", err)
	}
	logger.Logger().Info("Teste TestGerarSHA512 executado com sucesso!")
}

func TestCompararSHA512_CredencialValida(t *testing.T) {
	hash := "bed4efa1d4fdbd954bd3705d6a2a78270ec9a52ecfbfb010c61862af5c76af1761ffeb1aef6aca1bf5d02b3781aa854fabd2b69c790de74e17ecfec3cb6ac4bf"
	decryptedPassword := "password123"

	err := CompararSHA512(hash, decryptedPassword)
	if err != nil {
		t.Errorf("Erro ao comparar hashes SHA512: %v", err)
		logger.Logger().Error("Erro ao comparar hashes SHA512", err)
	}
	logger.Logger().Info("Teste TestCompararSHA512_CredenciailValida executado com sucesso!")
}

func TestCompararSHA512_CredencialInvalida(t *testing.T) {
	hash := "bed4efa1d4fdbd954bd3705d6a2a78270ec9a52ecfbfb010c61862af5c76af1761ffyb1aef6aca1bf5d02b3781aa854fabd2b69c790de74e17ecfec3cb6ac4bf"
	decryptedPassword := "password123"

	err := CompararSHA512(hash, decryptedPassword)
	if err == nil {
		t.Error("Erro esperado ao comparar hashes SHA512, mas não ocorreu")
		logger.Logger().Error("Erro esperado ao comparar hashes SHA512, mas não ocorreu", err)
	}
	logger.Logger().Info("Teste TestCompararSHA512_CredencialInvalida executado com sucesso!")
}
