package encryption

import (
	// "GoTaskManager/internal/app/inicializar"
	"GoTaskManager/pkg/pacotes/logger"
	"os"
	"testing"
)

// TestMain:Função executada antes das demais
func TestMain(m *testing.M) {
	exitCode := m.Run()
	if exitCode == 0 {
		logger.Logger().Info("Testes do pacote encryption executados com sucesso!")
	} else {
		logger.Logger().Alerta("Ocorreram erros ao executar os testes do pacote encryption")
	}
	os.Exit(exitCode)
}

func TestGerarSHA512(t *testing.T) {
	t.Parallel()
	data := "password123"
	hash, err := GerarSHA512(data)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao executar a função GerarSHA512", err)
		t.FailNow()
	}

	if hash == "" {
		logger.Logger().Error("Teste "+t.Name()+":	A função GerarSHA512 não retornou qualquer valor", err)
		t.FailNow()
	}
	logger.Logger().Info("Teste " + t.Name() + ":	Executado com sucesso!")
}

func TestCompararSHA512_CredencialValida(t *testing.T) {
	t.Parallel()
	hash := "bed4efa1d4fdbd954bd3705d6a2a78270ec9a52ecfbfb010c61862af5c76af1761ffeb1aef6aca1bf5d02b3781aa854fabd2b69c790de74e17ecfec3cb6ac4bf"
	decryptedPassword := "password123"

	err := CompararSHA512(hash, decryptedPassword)
	if err != nil {
		logger.Logger().Alerta("Teste " + t.Name() + ":	A função CompararSHA512 identificou que a credencial é inválida, no entanto deveria ser valida")
		t.FailNow()
	}
	logger.Logger().Info("Teste " + t.Name() + ":	Executado com sucesso!")
}

func TestCompararSHA512_CredencialInvalida(t *testing.T) {
	t.Parallel()
	hash := "bed4efa1d4fdbd954bd3705d6a2a78270ec9a52ecfbfb010c61862af5c76af1761ffyb1aef6aca1bf5d02b3781aa854fabd2b69c790de74e17ecfec3cb6ac4bf"
	decryptedPassword := "password123"

	err := CompararSHA512(hash, decryptedPassword)
	if err == nil {
		logger.Logger().Alerta("Teste " + t.Name() + ":	A função CompararSHA512 identificou que a credencial é valida, no entanto deveria ser inválida")
		t.FailNow()
	}
	logger.Logger().Info("Teste " + t.Name() + ":	Executado com sucesso!")
}
