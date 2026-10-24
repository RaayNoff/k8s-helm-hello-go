# Сборка приложения
FROM golang:1.27.2-alpine AS builder

WORKDIR /app

# Кэшируем зависимости отдельно от исходников
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Минимальный образ для запуска
FROM scratch

COPY --from=builder /app/server /server

ENTRYPOINT ["/server"]