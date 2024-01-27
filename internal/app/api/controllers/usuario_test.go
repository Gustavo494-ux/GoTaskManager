package controllers_test

import (
	"GoTaskManager/internal/app/inicializar"
	"GoTaskManager/internal/app/inicializar/inicializarinternal"
	"GoTaskManager/internal/app/models"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"

	clientehttp "github.com/Gustavo494-ux/PacotesGolang/clienteHttp"
	"github.com/Gustavo494-ux/PacotesGolang/configuracoes"
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

	defer DeletarTodosUsuarios()

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
		DeletarTodosUsuarios()
		for _, usuario := range Usuarios {
			CriarUsuarioSucesso(t, usuario)
		}
	})

	t.Run("CriarUsuarioCorpoInvalido", func(t *testing.T) {
		DeletarTodosUsuarios()
		var ponteiroUsuarios *[]models.Usuario = &Usuarios
		LimparCampoAleatorio(ponteiroUsuarios)

		for _, usuario := range *ponteiroUsuarios {
			CriarUsuarioCorpoInvalido(t, usuario)
		}
	})

	t.Run("CriarUsuarioExistente", func(t *testing.T) {
		DeletarTodosUsuarios()
		for _, usuario := range Usuarios {
			CriarUsuarioExistente(t, usuario)
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

func CriarUsuarioExistente(t *testing.T, usuario models.Usuario) {
	PopularUsuarios()
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
		{Nome: "Maria Oliveira", CPF: "98765432100", Email: "maria@gmail.com", Senha: "senha123"},
		{Nome: "maria", CPF: "11122233344", Email: "emailinvalido", Senha: "senha123"},
		{Nome: "Carlos", CPF: "11125233344", Email: "carlos@gmail.com", Senha: "senha123"},
		{Nome: "Ana34", CPF: "55566477788", Email: "ana@gmail.com", Senha: "senha123"},
		{Nome: "Lucas", CPF: "99988877766", Email: "lucas@gmail.com", Senha: "senha123"},
		{Nome: "Gabriela", CPF: "12345678909", Email: "gabriela@gmail.com", Senha: "senha123"},
		{Nome: "Fernanda", CPF: "12345678002", Email: "fernanda@gmail.com", Senha: "senha123"},
		{Nome: "Gustavo", CPF: "12345678955", Email: "gustavo@gmail.com", Senha: "senha123"},
		{Nome: "Roberto", CPF: "11122533344", Email: "roberto@gmail.com", Senha: "senha123"},
		{Nome: "Carlos Oliveira", CPF: "56789012345", Email: "carlos.oliveira@example.com", Senha: "senha123"},
		{Nome: "Juliana Silva", CPF: "67890123456", Email: "juliana.silva@example.com", Senha: "senhajul"},
		{Nome: "Fernanda Souza", CPF: "78901234567", Email: "fernanda.souza@example.com", Senha: "senhaffts"},
		{Nome: "Ricardo Santos", CPF: "89012345678", Email: "ricardo.santos@example.com", Senha: "senharic"},
		{Nome: "Camila Pereira", CPF: "90123456789", Email: "camila.pereira@example.com", Senha: "senhacam"},
		{Nome: "Gabriel Lima", CPF: "01234567890", Email: "gabriel.lima@example.com", Senha: "senhagab"},
		{Nome: "Aline Oliveira", CPF: "12345098765", Email: "aline.oliveira@example.com", Senha: "senhaagl"},
		{Nome: "Rodrigo Silva", CPF: "23456789098", Email: "rodrigo.silva@example.com", Senha: "senharodg"},
		{Nome: "Patrícia Souza", CPF: "34567890987", Email: "patricia.souza@example.com", Senha: "senhapatt"},
		{Nome: "Lucas Pereira", CPF: "45678909876", Email: "lucas.pereira@example.com", Senha: "senha1234"},
		{Nome: "Larissa Souza", CPF: "76543210987", Email: "larissa.souza@example.com", Senha: "senha12345"},
		{Nome: "Vinícius Pereira", CPF: "65432109876", Email: "vinicius.pereira@example.com", Senha: "senhavin"},
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

func PopularUsuarios() {
	configuracoes.BancoPrincipalGORM.Create(&Usuarios)
}

func DeletarTodosUsuarios() {
	configuracoes.BancoPrincipalGORM.Unscoped().Where("1=1").Delete(&models.Usuario{})
}
