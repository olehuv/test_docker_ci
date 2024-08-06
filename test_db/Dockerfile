# syntax=docker/dockerfile:1
FROM golang:1.21

# Встановлення робочої директорії
WORKDIR /app

# Копіювання go модулів та коду
COPY go.mod ./
#COPY go.sum ./
RUN go mod download

COPY . .

# Скомпілюємо додаток
RUN go build -o /main

# Команда для запуску додатку
CMD ["/main"]
