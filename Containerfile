# СТАДИЯ 1: Сборка (builder)
FROM registry.altlinux.org/alt/alt:p11 AS builder

# Подготовка к установке GO v.1.21
RUN mkdir -p /app && \
    apt-get update && \
    apt-get install -y golang

# Копируем исходные коды для сборки
WORKDIR /app
COPY sources/ .

# Собираем бинарник
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o simple-quiz .



# СТАДИЯ 2: Финальный образ
FROM registry.altlinux.org/alt/alt:p11

# Создаем пользователя
RUN groupadd -r appuser && useradd -r -g appuser -u 1001 appuser

# Создаем директорию для данных
RUN mkdir -p /app && chown appuser:appuser /app

# Переключаемся на пользователя
USER appuser

# Копируем бинарник и ресурсы
WORKDIR /app
COPY --from=builder /app/simple-quiz .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["/app/simple-quiz"]