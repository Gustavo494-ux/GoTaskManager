package controllers_test

import (
	"GoTaskManager/internal/app/inicializar"
	"GoTaskManager/internal/app/inicializar/inicializarinternal"
	"GoTaskManager/internal/app/models"
	"fmt"
	clientehttp "github.com/Gustavo494-ux/PacotesGolang/clienteHttp"
	"github.com/Gustavo494-ux/PacotesGolang/logger"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"
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

// Testes
func TestCriarUsuario(t *testing.T) {
	t.Parallel()
	URLCriarUsuario = URLbase + "usuario"

	if len(Usuarios) == 0 {
		logger.Logger().Error(fmt.Sprintf("Teste %s: Nenhum usuário foi passado para a realização do teste", t.Name()), nil)
		t.FailNow()
	}
	t.Run("CriarUsuarioComSucesso", func(t *testing.T) {
		for _, usuario := range Usuarios {
			CriarUsuarioSucesso(t, usuario)
		}
	})

	t.Run("CriarUsuarioCorpoInvalido", func(t *testing.T) {
		var ponteiroUsuarios *[]models.Usuario = &Usuarios
		LimparCampoAleatorio(ponteiroUsuarios)
		for _, usuario := range *ponteiroUsuarios {
			CriarUsuarioCorpoInvalido(t, usuario)
		}
	})

	logger.Logger().Info(fmt.Sprintf("Teste %s:	Executado com sucesso!", t.Name()))
}

// SubTestes
func CriarUsuarioSucesso(t *testing.T, usuario models.Usuario) {
	requisicao := clientehttp.Requisicao("POST", URLCriarUsuario, usuario, nil)
	if requisicao.GetStatusCode() != http.StatusOK {
		logger.Logger().Error(fmt.Sprintf("Teste %s: retornou o status code %s o status code esperado é %d", t.Name(),
			strconv.Itoa(requisicao.GetStatusCode()), http.StatusOK), nil)
		t.FailNow()
	}
}

func CriarUsuarioCorpoInvalido(t *testing.T, usuario models.Usuario) {
	requisicao := clientehttp.Requisicao("POST", URLCriarUsuario, usuario, nil)
	if requisicao.GetStatusCode() != http.StatusBadRequest {
		logger.Logger().Error(fmt.Sprintf("Teste %s: retornou o status code %s o status code esperado é %d", t.Name(),
			strconv.Itoa(requisicao.GetStatusCode()), http.StatusBadRequest), nil)
		t.FailNow()
	}
}

// Funções utilitarias
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

func LimparCampoAleatorio(usuarios *[]models.Usuario) {
	// Inicializa o gerador de números aleatórios
	rand.Seed(time.Now().UnixNano())

	campos := []string{"Nome", "CPF", "Email", "Senha"}
	var campo string

	for i := range *usuarios {
		campo = campos[rand.Intn(len(campos))]
		val := reflect.ValueOf(&(*usuarios)[i]).Elem()
		fieldVal := val.FieldByName(campo)
		if fieldVal.IsValid() && fieldVal.CanSet() {
			switch fieldVal.Kind() {
			case reflect.String:
				fieldVal.SetString("")
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				fieldVal.SetInt(0)
			case reflect.Float32, reflect.Float64:
				fieldVal.SetFloat(0)
			case reflect.Bool:
				fieldVal.SetBool(false)
			default:
				panic(fmt.Sprintf("Não é possível definir o valor do campo %s", campo))
			}
		} else {
			panic(fmt.Sprintf("Campo %s inválido", campo))
		}
	}
}
