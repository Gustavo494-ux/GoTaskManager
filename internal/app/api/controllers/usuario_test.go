package controllers_test

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/Gustavo494-ux/PacotesGolang/authentication"
	"github.com/Gustavo494-ux/PacotesGolang/clienteHttp"
	"github.com/Gustavo494-ux/PacotesGolang/configuracoes"
	"github.com/Gustavo494-ux/PacotesGolang/database"
	"github.com/Gustavo494-ux/PacotesGolang/inicializarpkg"
	"github.com/Gustavo494-ux/PacotesGolang/logger"

	"GoTaskManager/internal/app/inicializar"
	"GoTaskManager/internal/app/inicializar/inicializarinternal"
	"GoTaskManager/internal/app/models"
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

func TestCriarUsuario(t *testing.T) {
	type args struct {
		Usuarios           []models.Usuario
		funcoesInicializar []func()
		funcoesFinalizar   []func()
	}

	testes := []struct {
		nome               string
		args               args
		StatusCodeEsperado int
	}{
		{
			nome: "CriarUsuarioComSucesso",
			args: args{
				Usuarios:           Usuarios,
				funcoesInicializar: []func(){DeletarTodosUsuarios},
			},
			StatusCodeEsperado: http.StatusCreated,
		},
		{
			nome: "CriarUsuarioExistente",
			args: args{
				Usuarios:           Usuarios,
				funcoesInicializar: []func(){DeletarTodosUsuarios, PopularUsuarios},
			},
			StatusCodeEsperado: http.StatusConflict,
		},
		{
			nome: "CriarUsuarioCorpoInvalido",
			args: args{
				Usuarios:           LimparCampoAleatorio(Usuarios),
				funcoesInicializar: []func(){DeletarTodosUsuarios},
			},
			StatusCodeEsperado: http.StatusBadRequest,
		},
		{
			nome: "CriarUsuarioBancoDeDadosIndisponivel",
			args: args{
				Usuarios: Usuarios,
				funcoesInicializar: []func(){
					func() {
						database.DesconectarGorm(configuracoes.BancoPrincipalGORM)
					},
				},
				funcoesFinalizar: []func(){inicializarpkg.InicializarBancoDeDadosTeste},
			},
			StatusCodeEsperado: http.StatusServiceUnavailable,
		},
	}

	for _, teste := range testes {
		for _, funcao := range teste.args.funcoesInicializar {
			funcao()
		}

		t.Run(teste.nome, func(t *testing.T) {
			var statusCode int
			var err error
			var corpo string

			for _, usuario := range teste.args.Usuarios {
				statusCode, corpo, err = CriarUsuario(usuario)
				if err != nil {
					t.Errorf("teste %s: error = %v", t.Name(), err)
				}

				if statusCode != teste.StatusCodeEsperado {
					t.Errorf("teste %s: o status code recebido %d, é diferente do status code esperado %d. body da requisição %s",
						t.Name(), statusCode, teste.StatusCodeEsperado, corpo)
				}
			}
		})

		for _, funcao := range teste.args.funcoesFinalizar {
			funcao()
		}
	}
}

func CriarUsuario(usuario models.Usuario) (statusCode int, corpo string, err error) {
	URLCriarUsuario := URLbase + "usuario"
	if VerificarSeUsuarioVazio(usuario) {
		err = fmt.Errorf("nenhum usuário pode estar vazio ou possuir valor nil")
		return
	}

	requisicao := clienteHttp.Requisicao("POST", URLCriarUsuario, usuario, nil)
	statusCode = requisicao.GetStatusCode()

	body, err := requisicao.GetBody()
	if err != nil {
		err = fmt.Errorf("ao recuperar o corpo da requisição ocorreu o erro: %v", err)
	}

	corpo = string(body)
	return
}

//Refatorar daqui para baixo

// criar funcoes para tudo que puder ser reaproveitado, separar em outro pacote de util tudo que for generico
func TestBuscarTodosUsuarios(t *testing.T) {
	URl := fmt.Sprintf("%s%s", URLbase, "usuario")

	t.Run("BuscarTodosUsuariosSucesso", func(t *testing.T) {
		DeletarTodosUsuarios()
		PopularUsuarios()
		BuscarTodosUsuariosSucesso(t, URl)
	})

	t.Run("BuscarTodosUsuariosAutirizacaoExpirada", func(t *testing.T) {
		DeletarTodosUsuarios()
		PopularUsuarios()
		BuscarTodosUsuariosNaoAutorizado(t, URl)
	})

	// VerificarSeUsuarioVazio(t, Usuarios...)
	logger.Logger().Info(fmt.Sprintf("Teste %s:	Executado com sucesso!", t.Name()))
}

//SubTestes de buscar Todos Usuários

func BuscarTodosUsuariosSucesso(t *testing.T, URl string) {
	var usuarios []models.Usuario
	cabecalho := montarCabecalhoToken(t, time.Minute, Usuarios[0])
	requisicao := clienteHttp.Requisicao("GET", URl, nil, cabecalho)

	clienteHttp.ValidarStatusCodeRequisicaoTesting(t, requisicao, http.StatusOK)
	requisicao.GetBodyStructTesting(t, &usuarios)

	// VerificarSeUsuarioVazio(t, Usuarios...)
}

func BuscarTodosUsuariosNaoAutorizado(t *testing.T, URl string) {
	var usuarios []models.Usuario
	cabecalho := montarCabecalhoToken(t, time.Second*0, Usuarios[0])
	requisicao := clienteHttp.Requisicao("GET", URl, nil, cabecalho)

	clienteHttp.ValidarStatusCodeRequisicaoTesting(t, requisicao, http.StatusOK)
	requisicao.GetBodyStructTesting(t, &usuarios)

	// VerificarSeUsuarioVazio(t, Usuarios...)
}

// Funções utilitarias

func VerificarSeUsuarioVazio(usuarios ...models.Usuario) (existe bool) {
	if usuarios == nil {
		return true
	}

	if len(usuarios) == 0 {
		return true
	}

	return
}

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

func LimparCampoAleatorio(usuarios []models.Usuario) []models.Usuario {
	// Inicializa o gerador de números aleatórios
	rand.Seed(time.Now().UnixNano())

	// Criar uma cópia local dos usuários para evitar modificar os valores originais
	usuariosCopia := make([]models.Usuario, len(usuarios))
	copy(usuariosCopia, usuarios)

	campos := []string{"Nome", "CPF", "Email", "Senha"}
	var campo string

	for i := range usuariosCopia {
		campo = campos[rand.Intn(len(campos))]
		val := reflect.ValueOf(&usuariosCopia[i]).Elem()
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
	return usuariosCopia
}

func PopularUsuarios() {
	configuracoes.BancoPrincipalGORM.Create(&Usuarios)
}

func DeletarTodosUsuarios() {
	configuracoes.BancoPrincipalGORM.Unscoped().Where("1=1").Delete(&models.Usuario{})
}

func montarCabecalhoToken(t *testing.T, Expiraem time.Duration, usuario models.Usuario) (cabecalho map[string]string) {
	// VerificarSeUsuarioVazio(t, usuario)
	defer func() {
		if cabecalho == nil {
			logger.Logger().Error(fmt.Sprintf("Teste %s: o cabecalho da requisição não pode ser nulo", t.Name()),
				nil)
			t.FailNow()
		}
	}()

	token, err :=
		authentication.
			NovoToken(true, time.Now().Add(Expiraem).Unix()).
			AdicionarParametro("idUsuario", usuario.ID).
			Criar()
	if err != nil {
		logger.Logger().Error(fmt.Sprintf("Teste %s: ocorreu um erro ao criar um token", t.Name()),
			nil)
		t.FailNow()
	}

	if token == "" {
		logger.Logger().Error(fmt.Sprintf("Teste %s: o token não pode ser vazio", t.Name()),
			nil)
		t.FailNow()
	}

	cabecalho = map[string]string{
		"Authorization": "Bearer " + token,
	}
	return
}
