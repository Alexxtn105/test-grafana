# Запуск сервиса с выводом логов в файл
```bash
go run main.go > sample_app.log 2>&1
```
или
```bash
go run main.go > C:/ProgramData/GrafanaLabs/Alloy/data/sample_app.log 2>&1
```

или
```bash
go run main.go > d:/logs/sample_app.log 2>&1

```



Запуск контейнера Alloy:
```bash
docker run -v /d/Projects/Grafana/Alloy/config.alloy:/etc/alloy/config.alloy -v /d/logs:/tmp/logs -p 12345:12345 grafana/alloy:latest run --server.http.listen-addr=0.0.0.0:12345 --storage.path="d:\Projects\Grafana\Alloy" /etc/alloy/config.alloy
```