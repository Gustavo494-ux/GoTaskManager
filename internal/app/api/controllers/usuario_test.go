package controllers_test

import (
	"GoTaskManager/internal/app/inicializar"
	"GoTaskManager/internal/app/inicializar/inicializarinternal"
	"GoTaskManager/internal/app/models"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/Gustavo494-ux/PacotesGolang/clienteHttp"
	"github.com/Gustavo494-ux/PacotesGolang/logger"
)

var (
	Protocolo = "http"
	Host      = "localhost"
	Porta     int
	Usuarios  []models.Usuario

	URLbase         string
	URLCriarUsuario string
)

func TestMain(m *testing.M) {
	mockUsuarios()
	go inicializar.InicializarParaTestes()

	time.Sleep(time.Millisecond * 200)

	Porta = inicializarinternal.RetonarPortaApi()
	URLbase = fmt.Sprintf("%s://%s:%d/", Protocolo, Host, Porta)

	exitCode := m.Run()
	if exitCode == 0 {
		logger.Logger().Info("Testes do pacote controllers_test executados com sucesso!")
	} else {
		logger.Logger().Alerta("Ocorreram erros ao executar os testes do pacote controllers_test")
	}

	os.Exit(exitCode)
}

func TestCriarUsuario(t *testing.T) {
	URLCriarUsuario = URLbase + "usuario"
	t.Parallel()
	if len(Usuarios) == 0 {
		logger.Logger().Error(fmt.Sprintf("Teste %s: Nenhum usuário foi passado para a realização do teste", t.Name()), nil)
		t.FailNow()
	}
	tempo := time.Now()
	for _, usuario := range Usuarios {
		requisicao := clientehttp.POST(URLCriarUsuario, usuario)
		statusCodeRequisicao := requisicao.GetStatusCode()

		if statusCodeRequisicao != http.StatusOK {
			logger.Logger().Error(fmt.Sprintf("Teste %s: retornou o status code %s o status code esperado é %d", t.Name(), strconv.Itoa(statusCodeRequisicao), http.StatusOK), nil)
			t.FailNow()
		}
	}
	logger.Logger().Info(fmt.Sprintf("Teste %s:	Executado com sucesso! tempo decorrido: %s", t.Name(), time.Since(tempo).Abs().String()))
}

func mockUsuarios() {
	Usuarios = []models.Usuario{
		{Nome: "João Silva", CPF: "12345678903", Email: "test@example.com", Senha: "senha123"},
		{Nome: "Maria Oliveira", CPF: "98765432100", Email: "maria@gmail.com", Senha: "outrasenha"},
		{Nome: "maria", CPF: "11122233344", Email: "emailinvalido", Senha: "senha123"},
		{Nome: "Carlos", CPF: "11122233344", Email: "carlos@gmail.com", Senha: "senha123"},
		{Nome: "Ana34", CPF: "55566677788", Email: "ana@gmail.com", Senha: "senha123456"},
		{Nome: "Lucas", CPF: "99988877766", Email: "lucas@gmail.com", Senha: "senha123"},
		{Nome: "Gabriela", CPF: "12345678909", Email: "gabriela@gmail.com", Senha: "senha123456"},
		{Nome: "Fernanda", CPF: "12345678002", Email: "fernanda@gmail.com", Senha: "senha123"},
		{Nome: "Gustavo", CPF: "12345678955", Email: "gustavo@gmail.com", Senha: "senha12345"},
		{Nome: "Roberto", CPF: "11122233344", Email: "roberto@gmail.com", Senha: "senha123"},
	}
}
