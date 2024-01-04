package manipuladorDeArquivos_test

import (
	"GoTaskManager/pkg/pacotes/logger"
	"GoTaskManager/pkg/pacotes/manipuladorDeArquivos"
	// "GoTaskManager/pkg/routines/configuracoes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestMain:Função executada antes das demais
func TestMain(m *testing.M) {
	exitCode := m.Run()

	if exitCode == 0 {
		logger.Logger().Info("Testes do pacote manipuladorDeArquivos executados com sucesso!")
	} else {
		logger.Logger().Alerta("Ocorreram erros ao executar os testes do pacote manipuladorDeArquivos")
	}

	os.Exit(exitCode)
}

func TestCriarArquivo(t *testing.T) {
	fileName := t.Name()
	t.Parallel()
	dir, _ := os.Getwd()
	_, err := manipuladorDeArquivos.CriarArquivo(dir, fileName)
	if err != nil {
		logger.Logger().Error("Teste: "+t.Name()+":	Ocorreu um erro ao  executar a função CriarArquivo", err)
		t.FailNow()
	}
	os.Remove(fileName)
	logger.Logger().Info("Teste " + t.Name() + ":	Executado com sucesso!")
}

func TestAbrirArquivo(t *testing.T) {
	t.Parallel()

	fileName := t.Name()

	dir, _ := os.Getwd()
	expectedContent := "Conteúdo do arquivo de teste"
	err := os.WriteFile(fileName, []byte(expectedContent), 0644)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao escrever o arquivo de teste", err)
		t.FailNow()
	}
	defer os.Remove(fileName)

	content, err := manipuladorDeArquivos.AbrirArquivo(dir, fileName)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao abrir o arquivo de teste", err)
		t.FailNow()
	}

	if content != expectedContent {
		logger.Logger().Alerta("Teste " + t.Name() + ":	A função AbrirArquivo não retornou o resultado esperado")
		t.FailNow()
	}

	logger.Logger().Info("Teste " + t.Name() + ":	Executado com sucesso!")
}

func TestEscreverArquivo(t *testing.T) {
	fileName := t.Name()

	t.Parallel()
	dir, _ := os.Getwd()
	content := "Conteúdo do arquivo de teste"
	err := manipuladorDeArquivos.EscreverArquivo(dir, fileName, content)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao executar a função EscreverArquivo", err)
		t.FailNow()
	}
	defer os.Remove(fileName)

	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao ler o arquivo de teste", err)
		t.FailNow()
	}

	if string(fileContent) != content {
		logger.Logger().Alerta("Teste " + t.Name() + ":	A função EscreverArquivo não escreveu o conteúdo no arquivo")
		t.FailNow()
	}

	logger.Logger().Info("Teste " + t.Name() + ":	Executado com sucesso!")
}

func TestAdicionarNoArquivo(t *testing.T) {
	fileName := t.Name()

	t.Parallel()
	initialContent := "Conteúdo inicial"
	appendedContent := "Conteúdo anexado"

	directoryPath, _ := os.Getwd()
	fullPath := strings.ReplaceAll(filepath.Join(directoryPath, fileName), "\\", "/")
	directoryPath = directoryPath + "\\"

	err := os.WriteFile(fullPath, []byte(initialContent), 0644)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao escrever o arquivo de teste", err)
		t.FailNow()
	}
	defer os.Remove(fullPath)

	err = manipuladorDeArquivos.AdicionarAoArquivo(directoryPath, fileName, appendedContent)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao anexar ao arquivo teste", err)
		t.FailNow()
	}

	fileContent, err := os.ReadFile(fullPath)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao ler o arquivo de teste", err)
		t.FailNow()
	}

	expectedContent := initialContent + appendedContent
	if string(fileContent) != expectedContent {
		logger.Logger().Error("Teste "+t.Name()+":	A função AdicionarAoArquivo não anexou o conteúdo.", err)
		t.FailNow()
	}

	logger.Logger().Info("Teste " + t.Name() + ":	Executado com sucesso!")
}

func TestDeleteFile(t *testing.T) {
	fileName := t.Name()

	t.Parallel()
	dir, _ := os.Getwd()
	fullPath := filepath.Join(dir, fileName)
	file, err := os.Create(fullPath)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao criar o arquivo de teste", err)
		t.FailNow()
	}
	file.Close()
	err = manipuladorDeArquivos.ExcluirArquivo(dir, fileName)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao executar a função ExcluirArquivo", err)
		t.FailNow()
	}

	_, err = os.Stat(fileName)
	if !os.IsNotExist(err) {
		logger.Logger().Error("Teste "+t.Name()+":	A função ExcluirArquivo não excluiu o arquivo", err)
		t.FailNow()
	}

	logger.Logger().Info("Teste " + t.Name() + ":	Executado com sucesso!")
}

func TestRenameFile(t *testing.T) {
	fileName := t.Name()

	t.Parallel()
	dir, _ := os.Getwd()
	newFileName := "newfile.txt"

	fullPathInitial := strings.ReplaceAll(filepath.Join(dir, fileName), "\\", "/")
	fullPathRename := strings.ReplaceAll(filepath.Join(dir, newFileName), "\\", "/")

	err := os.WriteFile(fullPathInitial, []byte("Conteúdo do arquivo de teste"), 0644)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao escrever o arquivo de teste", err)
		t.FailNow()
	}
	defer os.Remove(newFileName)

	err = manipuladorDeArquivos.RenomearArquivo(dir, fileName, newFileName)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao executar a função RenomearArquivo", err)
		t.FailNow()
	}

	_, err = os.Stat(fullPathInitial)
	if !os.IsNotExist(err) {
		logger.Logger().Error("Teste "+t.Name()+": Ocorreu um erro ao buscar as propriedades do arquivo", err)
		t.FailNow()
	}

	_, err = os.Stat(fullPathRename)
	if os.IsNotExist(err) {
		logger.Logger().Error("Teste "+t.Name()+": A função RenomearArquivo não renomeou o arquivo", err)
		t.FailNow()
	}

	logger.Logger().Info("Teste " + t.Name() + ":	Executado com sucesso!")
}

func TestGetFileList(t *testing.T) {
	t.Parallel()
	// Criar um diretório temporário para fins de teste
	currentDirectory, err := os.Getwd()
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+": Ocorreu um erro ao buscar o caminho do diretório atual", err)
		t.FailNow()
	}

	fullPath := strings.ReplaceAll(filepath.Join(currentDirectory, "test_directory"), "\\", "/")

	err = os.MkdirAll(fullPath, 0750)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+": Ocorreu um erro ao criar o diretório temporario", err)
		t.FailNow()
	}
	defer os.RemoveAll("test_directory")

	// Criar arquivos de teste dentro do diretório
	testFiles := []string{"file1.txt", "file2.txt", "file3.txt"}
	for _, filename := range testFiles {
		filePath := filepath.Join(fullPath, filename)
		file, err := os.Create(filePath)
		if err != nil {
			logger.Logger().Error("Teste "+t.Name()+": Ocorreu um erro ao criar o arquivo de teste", err)
			t.FailNow()
		}
		defer file.Close()
	}

	// Executar a função GetFileList no diretório de teste
	fileList, err := manipuladorDeArquivos.ObterListaArquivos(fullPath)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+": Ocorreu um erro ao executar a função ObterListaArquivos", err)
		t.FailNow()
	}

	// Verificar se todos os arquivos de teste estão presentes na lista de arquivos retornada
	for _, filename := range testFiles {
		found := false
		for _, file := range fileList {
			if file == filename {
				found = true
				break
			}
		}
		if !found {
			logger.Logger().Error("Teste "+t.Name()+": Arquivo esperado ausente na lista de arquivos retornada", err)
			t.FailNow()
		}
	}

	// Verificar se não há arquivos extras na lista de arquivos retornada
	if len(fileList) != len(testFiles) {
		t.Errorf("O número de arquivos retornados é diferente do esperado. Esperado: %d, Retornado: %d", len(testFiles), len(fileList))
		logger.Logger().Error(fmt.Sprintf("O número de arquivos retornados é diferente do esperado. Esperado: %d, Retornado: %d", len(testFiles), len(fileList)), err)

		logger.Logger().Alerta("Teste "+t.Name()+": O número de arquivos retornados é diferente do esperado. Esperado: %d, Retornado: %d", len(testFiles), len(fileList))
		t.FailNow()
	}

	logger.Logger().Info("Teste " + t.Name() + ":	Executado com sucesso!")
}

func TestCreateDirectory(t *testing.T) {
	t.Parallel()
	// Especificar o caminho base para os novos diretórios
	basePath := "./src/utility/teste"

	// Criar uma lista de nomes de pastas
	pastas := []string{"pasta1", "pasta2", "pasta3", "pasta4", "pasta5", "pasta6", "pasta7", "pasta8", "pasta9", "pasta10"}

	// Iterar sobre a lista de pastas
	for _, pasta := range pastas {
		// Construir o caminho completo para cada diretório
		caminho := filepath.Join(basePath, pasta)

		// Chamar a função CreateDirectory
		err := manipuladorDeArquivos.CriarDiretorio(caminho)
		if err != nil {
			logger.Logger().Error("Teste "+t.Name()+": Ocorreu um erro ao executar a função CriarDiretorio", err)
			t.FailNow()
		}
	}

	sliceBasePath := strings.Split(basePath, "/")
	err := os.RemoveAll("./" + sliceBasePath[1])
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+": Ocorreu um erro ao remover os diretorios", err)
		t.FailNow()
	}

	logger.Logger().Info("Teste " + t.Name() + ":	Executado com sucesso!")
}

func TestGetFileInfo(t *testing.T) {
	t.Parallel()
	// Teste para um arquivo
	caminhoDoArquivo := "./src/utility/file.txt"
	infoDoArquivo, err := manipuladorDeArquivos.ObterInformacoesArquivo(caminhoDoArquivo)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+": Ocorreu um erro ao executar a função ObterInformacoesArquivo", err)
		t.FailNow()
	}

	if infoDoArquivo != nil {
		// Verificar se é um arquivo
		if !infoDoArquivo.Mode().IsRegular() {
			logger.Logger().Error("Teste "+t.Name()+": Esperado um arquivo, obtido um diretório", err)
			t.FailNow()
		}
	}

	// Teste para um diretório
	caminhoDoDiretório := "./src/utility"
	infoDoDiretório, err := manipuladorDeArquivos.ObterInformacoesArquivo(caminhoDoDiretório)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+": Ocorreu um erro ao executar a função ObterInformacoesArquivo", err)
		t.FailNow()
	}

	if infoDoDiretório != nil {
		// Verificar se é um diretório
		if !infoDoDiretório.Mode().IsDir() {
			logger.Logger().Error("Teste "+t.Name()+": Esperado um diretório, obtido um arquivo", err)
			t.FailNow()
		}
	}

	logger.Logger().Info("Teste " + t.Name() + ":	Executado com sucesso!")
}
