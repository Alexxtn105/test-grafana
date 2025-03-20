# Запуск сервиса с выводом логов в файл (запускать в bash!!!)
```bash
go run main.go > d:/logs/sample_app.log 2>&1
```

# Установка необходимых инструментов в Docker:
Установка `Grafana/Loki`

[Официальная документация](https://grafana.com/docs/loki/latest/setup/install/docker/)
```bash
docker pull grafana/loki:3.4.2
```

Установка `Grafana/Alloy`

[Официальная документация](https://grafana.com/docs/alloy/latest/set-up/install/docker/)
```bash
docker pull grafana/alloy:v1.7.4
```

Установка `Grafana`

[Официальная документация](https://grafana.com/docs/grafana/latest/setup-grafana/installation/docker/)
```bash
docker pull grafana/grafana
```

Установка драйвера `Loki` для `Grafana` (необязательно):
```bash
docker plugin install grafana/loki-docker-driver:3.4.2-amd64 --alias loki_driver --grant-all-permissions
```

# Запуск контейнеров Docker
ВНИМАНИЕ! Запускать в `cmd`! Не в `bash`!

Логи пишутся в папку `d:\logs`.

Запуск контейнера `alloy` (запускается только с абсолютными путями):
```bash
docker run -v "d:\Projects\Go\test-grafana\configs\alloy.cfg":/etc/alloy/config.alloy -v "d:\logs":/tmp/logs -p 12345:12345 grafana/alloy:v1.7.4 run --server.http.listen-addr=0.0.0.0:12345 --storage.path="d:\Projects\Grafana\Alloy" /etc/alloy/config.alloy
```

Запуск контейнера `loki`:
```bash
docker run -d --name=loki --mount type=bind,source="d:\Projects\Go\test-grafana\configs\loki-config.yml",target=/etc/loki/local-config.yaml -p 3100:3100 grafana/loki:3.4.2
```

Запуск контейнера `grafana`:
```bash
docker run -d --name=grafana -p 3000:3000 grafana/grafana
```

# Решение проблемы с дедлоками loki (контейнер не останавливается)

Причина этой проблемы в том, что официальная документация ужасно плоха и не обновляется.
Решение проблемы заключается в том, чтобы не следовать документации, а вместо этого вручную установить последнюю версию вместо `latest`.
Актуальная версия `Loki` на момент написания документации была `3.4.2`.

При установке `grafana/loki` выше версии `3.0.0` необходимо прочитать это: https://grafana.com/docs/loki/latest/setup/upgrade/

Установка:

1. Не следуйте инструкциям на https://grafana.com/docs/loki/latest/setup/install/, потому что там используется `Loki` `2.9.2`. 
   Вместо этого используйте https://github.com/grafana/loki/blob/main/production/docker-compose.yaml.
2. Найдите последнюю версию `alloy`, и вручную измените версию образа на `v1.7.4`, не используйте `:latest`.
   Почему? Версия плагина `loki-docker-driver` всегда показывает `loki:latest`
   Не следуйте инструкциям на https://grafana.com/docs/loki/latest/send-data/docker-driver/, установите последнюю версию `docker-driver`, найдите https://hub.docker.com/r/grafana/loki-docker-driver/tags, я нашёл версию `3.4.2`.
   ```bash
   docker plugin install grafana/loki-docker-driver:3.4.2-amd64 --alias loki_driver --grant-all-permissions
   ```
3. Проблем с версией `Grafana` не выявлено. 

В итоге, я использую:
   - `grafana/loki:3.4.2` 
   - `grafana/alloy:v1.7.4` 
   - `grafana:latest` 
   - `loki-docker-driver:3.4.2`

