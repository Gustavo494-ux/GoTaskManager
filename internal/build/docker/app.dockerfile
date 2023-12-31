# internal/build/docker/app.dockerfile

# Use a imagem base desejada
FROM golang:latest

ENV APP_DIR=/go/src/usr/app/GoTaskManager/

# Define o diretório de trabalho dentro do contêiner
WORKDIR "$APP_DIR"

# Copia todos os arquivos do contexto de construção para o diretório de trabalho
COPY go.mod go.sum ./

# Executa o download das dependências
RUN go mod download && go mod verify

COPY . ./

# Baixar as dependências
RUN go get -d -v ./...

# Build do projeto
RUN cd ./internal/app/ && go build -o aplicativo .

VOLUME "$APP_DIR"

# Comando para executar o aplicativo quando o contêiner for iniciado
CMD ["./internal/app/aplicativo"]
