# Estágio de build
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copia os arquivos de módulo e baixa as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código-fonte
COPY . .

# Compila a aplicação
# O -o /app/main especifica o nome e o local do arquivo de saída
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main ./cmd/api/main.go

# Estágio final
FROM alpine:latest

WORKDIR /root/

# Copia o binário compilado do estágio de build
COPY --from=builder /app/main .

# Copia as migrações do banco de dados
COPY ./db/migrations ./db/migrations

# Copia o arquivo de exemplo de variáveis de ambiente
COPY .env.example .

# Expõe a porta que a aplicação vai rodar
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./main"]
