# Используем официальный образ Go как базовый
FROM golang:1.22.2-alpine as builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем исходники приложения в рабочую директорию
COPY . .
# Скачиваем все зависимости
RUN go mod tidy

# Собираем приложение
WORKDIR /app/cmd
RUN go build -o main

# Начинаем новую стадию сборки на основе минимального образа
FROM alpine:latest
COPY ./grafana/provisioning/dashboards /etc/grafana/provisioning/dashboards
COPY ./grafana/dashboards /var/lib/grafana/dashboards

# Добавляем исполняемый файл из первой стадии в корневую директорию контейнера
COPY --from=builder /app/cmd/main /main


# Открываем порт 8080
EXPOSE 8080

# Запускаем приложение
CMD ["/main"]

