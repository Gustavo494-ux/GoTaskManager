package manipuladorDeArquivos

import (
	"GoTaskManager/pkg/pacotes/logger"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	DiretorioRaiz string
)

func DefinirDiretorioRaiz(diretorio string) {
	DiretorioRaiz = diretorio
}

func RetornarDiretorioRaiz() string {
	return DiretorioRaiz
}

// CriarArquivo cria um novo arquivo com o nome especificado no diretório fornecido.
func CriarArquivo(caminhoDiretorio string, nomeArquivo string) (arquivo *os.File, err error) {
	arquivo, err = os.Create(filepath.Join(FormatarCaminho(caminhoDiretorio), nomeArquivo))
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao executar a função CriarArquivo", err)
		return nil, err
	}
	defer arquivo.Close()
	return arquivo, nil
}

// AbrirArquivo abre um arquivo existente para leitura.
func AbrirArquivo(caminhoDiretorio string, nomeArquivo string) (string, error) {
	data, err := os.ReadFile(filepath.Join(FormatarCaminho(caminhoDiretorio), nomeArquivo))
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao executar a função AbrirArquivo", err)
		return "", err
	}
	return string(data), nil
}

// CarregarArquivo abre um arquivo existente para leitura.
func CarregarArquivo(nomeArquivo string) (file *os.File, err error) {
	file, err = os.OpenFile(nomeArquivo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao executar a função CarregarArquivo", err)
	}
	return
}

// EscreverArquivo escreve o conteúdo fornecido em um arquivo.
func EscreverArquivo(caminhoDiretorio string, nomeArquivo string, conteudo string) error {
	err := os.WriteFile(filepath.Join(FormatarCaminho(caminhoDiretorio), nomeArquivo), []byte(conteudo), 0600)
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao executar a função EscreverrArquivo", err)
		return err
	}
	return nil
}

// AdicionarAoArquivo anexa o conteúdo fornecido a um arquivo existente.
func AdicionarAoArquivo(caminhoDiretorio string, nomeArquivo string, conteudo string) error {
	file, err := os.OpenFile(
		FormatarCaminho(filepath.Join(caminhoDiretorio, nomeArquivo)),
		os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao executar a função AdicionarAoArquivo", err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString(conteudo)
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao executar a função AdicionarAoArquivo", err)
		return err
	}
	return nil
}

// ExcluirArquivo exclui um arquivo.
func ExcluirArquivo(caminhoDiretorio string, nomeArquivo string) error {
	err := os.Remove(filepath.Join(FormatarCaminho(caminhoDiretorio), nomeArquivo))
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao executar a função EscluirArquivo", err)
		return err
	}
	return nil
}

// RenomearArquivo renomeia um arquivo.
func RenomearArquivo(caminhoDiretorio string, nomeArquivoAntigo string, nomeArquivoNovo string) error {
	err := os.Rename(filepath.Join(FormatarCaminho(caminhoDiretorio), nomeArquivoAntigo), filepath.Join(caminhoDiretorio, nomeArquivoNovo))
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao executar a função RenomearArquivo", err)
		return err
	}
	return nil
}

// ObterListaArquivos retorna a lista de arquivos em um diretório especificado.
func ObterListaArquivos(diretorio string) ([]string, error) {
	fileList := []string{}

	files, err := os.ReadDir(diretorio)
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao executar a função ObterListaArquivos", err)
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() {
			fileList = append(fileList, file.Name())
		}
	}

	return fileList, nil
}

// CriarDiretorio cria um diretório no caminho especificado.
func CriarDiretorio(caminho string) error {
	err := os.MkdirAll(caminho, 0750)
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao executar a função CriarDiretorio", err)
		return fmt.Errorf("erro ao criar o diretório: %v", err)
	}
	return nil
}

// ObterInformacoesArquivo retorna informações sobre um arquivo ou diretório especificado pelo caminho fornecido.
func ObterInformacoesArquivo(caminho string) (os.FileInfo, error) {
	// Converte o caminho para um caminho absoluto se for relativo
	caminhoAbs, err := filepath.Abs(caminho)
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao executar a função ObterInformacoesArquivo", err)
		return nil, fmt.Errorf("erro ao resolver o caminho absoluto: %v", err)
	}

	// Verifica se o arquivo ou diretório existe
	_, err = os.Stat(caminhoAbs)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil // Retorna nil quando o diretório não é encontrado
		}
		logger.Logger().Error("Ocorreu um erro ao executar a função ObterInformacoesArquivo", err)
		return nil, fmt.Errorf("erro ao obter as informações do arquivo: %v", err)
	}

	// Recupera as informações do arquivo
	infoArquivo, err := os.Stat(caminhoAbs)
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao executar a função ObterInformacoesArquivo", err)
		return nil, fmt.Errorf("erro ao obter as informações do arquivo: %v", err)
	}

	return infoArquivo, nil
}

// CriarArquivoSeNaoExistir verifica se um arquivo existe, caso não exista, o mesmo será criado.
func CriarArquivoSeNaoExistir(caminho string) (err error) {
	dir, nomeArquivo := filepath.Split(caminho)
	infoArquivo, err := ObterInformacoesArquivo(caminho)
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao executar a função CriarArquivoSeNaoExistir", err)
		err = fmt.Errorf("erro ao obter as informações do arquivo: %s", err)
	}
	if infoArquivo == nil {
		_, err = CriarArquivo(dir, nomeArquivo)
		if err != nil {
			logger.Logger().Error("Ocorreu um erro ao executar a função CriarArquivoSeNaoExistir", err)
			err = fmt.Errorf("erro ao criar o arquivo: %s", err)
		}
	}
	return
}

// ObterCaminhoDiretorio recebe o caminho de um arquivo e extrai o caminho do diretório onde este arquivo será criado.
func ObterCaminhoUltimoDiretorio(caminho string) string {
	dirPath := strings.Split(FormatarCaminho(caminho), "/")
	dirPath = append(dirPath[:len(dirPath)-1], dirPath[len(dirPath):]...)
	dirPathCriar := ""
	for i, dir := range dirPath {
		if i > 0 {
			dirPathCriar += "/"
		}
		dirPathCriar += dir
	}
	return dirPathCriar
}

// CriarDiretorioOuArquivoSeNaoExistir recebe o caminho de um arquivo e, se não existir, cria todos os diretórios necessários e o próprio arquivo.
func CriarDiretorioOuArquivoSeNaoExistir(caminho string) (err error) {
	err = CriarDiretorio(ObterCaminhoUltimoDiretorio(caminho))
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao executar a função CriarDiretorioOuArquivoSeNaoExistir", err)
		err = fmt.Errorf("erro CriarDiretorioSeNaoExistir : %s", err)
		return
	}

	err = CriarArquivoSeNaoExistir(caminho)
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao executar a função CriarDiretorioOuArquivoSeNaoExistir", err)
		err = fmt.Errorf("erro CriarArquivoSeNaoExistir : %s", err)
		return
	}
	return
}

// ObterCaminhoAtePasta retorna o caminho até o nível da pasta desejada em um caminho completo.
func ObterCaminhoAtePasta(caminho string, nomePasta string) (string, error) {
	// Procurar a posição da última ocorrência do nome da pasta no caminho
	index := strings.LastIndex(strings.ToLower(caminho), nomePasta)
	if index == -1 {
		return "", fmt.Errorf("a pasta '%s' não foi encontrada no caminho '%s'", nomePasta, caminho)
	}

	// Obter o caminho até a última ocorrência do nome da pasta
	caminhoAtePasta := caminho[:index+len(nomePasta)]
	return caminhoAtePasta, nil
}

// ObterCaminhoAbsolutoOuConcatenadoComRaiz retorna o caminho absoluto ou o ultimo diretorio raiz + ultimo diretorio do parametro caminho + nome do arquivo
func ObterCaminhoAbsolutoOuConcatenadoComRaiz(caminho string) (string, error) {
	caminho = filepath.FromSlash(caminho)
	if filepath.IsAbs(caminho) {
		return caminho, nil
	}

	// Obter o último diretório e nome do arquivo do caminho
	dir, nomeArquivo := filepath.Split(caminho)
	dir = strings.TrimSuffix(dir, string(filepath.Separator))
	nomeArquivo = strings.TrimPrefix(nomeArquivo, string(filepath.Separator))

	// Concatenar o último diretório e nome do arquivo com o caminho raiz
	caminhoAbsoluto := filepath.Join(DiretorioRaiz, dir, nomeArquivo)
	return caminhoAbsoluto, nil
}

// FormatarCaminho: aplica um replace substituindo "\\" por "/"
func FormatarCaminho(caminho string) string {
	return strings.ReplaceAll(caminho, "\\", "/")
}

// ObterDiretorioDoArquivo otimiza a busca do diretório do arquivo na hierarquia
func ObterDiretorioDoArquivo(caminho, arquivo string) (string, error) {
	var diretorioEncontrado string
	for {
		err := filepath.Walk(caminho, func(currentDir string, info os.FileInfo, err error) error {
			// currentDir = ObterCaminhoUltimoDiretorio(currentDir)
			if err != nil {
				logger.Logger().Error("Erro ao percorrer diretório", err)
				return err
			}

			// Verifica se o arquivo está presente no diretório atual
			filePath := filepath.Join(currentDir, arquivo)
			_, err = os.Stat(filePath)
			if err == nil {
				// O arquivo foi encontrado, armazenamos o diretório atual
				diretorioEncontrado = currentDir
				return filepath.SkipDir // Pula a verificação em subdiretórios após encontrar o arquivo
			}

			return nil
		})
		if err != nil {
			logger.Logger().Error("Erro ao percorrer diretórios", err)
			return "", err
		}

		if diretorioEncontrado == "" {
			return ObterDiretorioDoArquivo(ObterCaminhoUltimoDiretorio(caminho), arquivo)
		} else {
			break
		}

	}

	return diretorioEncontrado, nil
}
