# Используем официальный образ Grafana
FROM grafana/grafana:latest

# Копируем файлы конфигурации Grafana
COPY ./proxy/grafana/provisioning/dashboards /etc/grafana/provisioning/dashboards
COPY ./proxy/grafana/dashboards /var/lib/grafana/dashboards

# Открываем порт 3000
EXPOSE 3000
