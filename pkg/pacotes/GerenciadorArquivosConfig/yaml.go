package GerenciadorArquivosConfig

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type ArquivoConfig struct {
	caminhoArquivo string
	dados          map[string]interface{}
}

type valorParametro struct {
	valor interface{}
}

const permissaoArquivo = 0644

// NovoArquivoConfig cria um novo ArquivoConfig.
func NovoArquivoConfig(caminhoArquivo string) *ArquivoConfig {
	return &ArquivoConfig{
		caminhoArquivo: caminhoArquivo,
		dados:          make(map[string]interface{}),
	}
}

// Ler lê o arquivo YAML e preenche os dados do struct.
func (a *ArquivoConfig) Ler() *ArquivoConfig {
	arquivoYAML, err := os.ReadFile(a.caminhoArquivo)
	if err != nil {
		log.Fatalf("falha ao ler o arquivo %s: %v", a.caminhoArquivo, err)
	}

	err = yaml.Unmarshal(arquivoYAML, &a.dados)
	if err != nil {
		log.Fatalf("falha ao unmarshallar os dados do YAML: %v", err)
	}

	return a
}

// ObterValorParametro retorna o valor do parâmetro específico do arquivo YAML.
func (a *ArquivoConfig) ObterValorParametro(parametro string) *valorParametro {
	valor, encontrado := a.dados[parametro]
	if !encontrado {
		log.Fatalf("parâmetro não encontrado")
	}
	return &valorParametro{valor: valor}
}

// AdicionarParametro atualiza ou adiciona um parâmetro no arquivo YAML.
func (a *ArquivoConfig) AdicionarParametro(parametro string, valor interface{}) *ArquivoConfig {
	a.dados[parametro] = valor
	return a
}

// Escrever escreve os dados do struct em um arquivo YAML.
func (a *ArquivoConfig) Escrever() error {
	dadosYAML, err := yaml.Marshal(a.dados)
	if err != nil {
		return fmt.Errorf("falha ao marshallar os dados para YAML: %w", err)
	}

	err = os.WriteFile(a.caminhoArquivo, dadosYAML, permissaoArquivo)
	if err != nil {
		return fmt.Errorf("falha ao escrever no arquivo %s: %w", a.caminhoArquivo, err)
	}

	return nil
}

// Int retorna o valor do parâmetro como um int.
func (v *valorParametro) Int() (int, error) {
	i, ok := v.valor.(int)
	if !ok {
		return 0, errors.New("valor não é um int")
	}

	return i, nil
}

// String retorna o valor do parâmetro como uma string.
func (v *valorParametro) String() (string, error) {
	str, ok := v.valor.(string)
	if !ok {
		return "", errors.New("valor não é uma string")
	}

	return str, nil
}

// Float retorna o valor do parâmetro como um float64.
func (v *valorParametro) Float() (float64, error) {
	f, ok := v.valor.(float64)
	if !ok {
		return 0, errors.New("valor não é um float64")
	}

	return f, nil
}

// Bool retorna o valor do parâmetro como um bool.
func (v *valorParametro) Bool() (bool, error) {
	b, ok := v.valor.(bool)
	if !ok {
		return false, errors.New("valor não é um bool")
	}

	return b, nil
}
