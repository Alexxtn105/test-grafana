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


Настройка контейнеров:
I've been struggling for a few days, but it seems like I've solved the problem. The reason for this issue is that the official documentation is terribly bad and hasn't been maintained.

The solution to the problem is to not follow the documentation, but instead manually install the latest version.

Listen! When I wrote this comment, the loki version is 3.2.1.

Step:

Do not follow https://grafana.com/docs/loki/latest/setup/install/, because it use loki 2.9.2, use https://github.com/grafana/loki/blob/main/production/docker-compose.yaml instead.
Find latest loki promtail version, It may be same in https://github.com/grafana/loki/releases/ , and manually change the image version to 3.2.1, do not use :latest.
why? see loki-docker-driver plugin version always shows loki:latest #3155
Do not follow https://grafana.com/docs/loki/latest/send-data/docker-driver/, Install latest docker-driver, find https://hub.docker.com/r/grafana/loki-docker-driver/tags, I found 3.2.1.
I dont find grafana has version issue.
docker plugin install grafana/loki-docker-driver:3.2.1 --alias loki --grant-all-permissions
In sum.

I use grafana/loki:3.2.1 promtail:3.2.1 grafana:latest loki-docker-driver:3.2.1

When you install grafana/loki above 3.0.0, you must read this.

https://grafana.com/docs/loki/latest/setup/upgrade/

#12588

The official documentation is really outdated and wasted a lot of my time. Time is valuable, and I hate this documentation. I despise it. I've never seen the worst documentation.

I really don't know why they haven't updated the documentation. Why are their employees in such chaos?

#12506

When your product has good documentation, it saves everyone time and benefits everyone.

I have already started a company, and I really wouldn't purchase a product without proper documentation.

---------------------------------------------------------


Запуск контейнера Alloy:
```bash
docker run -v "d:\Projects\Grafana\Alloy\alloy.cfg":/etc/alloy/config.alloy -v /d/logs:/tmp/logs -p 12345:12345 grafana/alloy:v1.7.4 run --server.http.listen-addr=0.0.0.0:12345 --storage.path="d:\Projects\Grafana\Alloy" /etc/alloy/config.alloy
```

Запуск контейнера loki
```bash
docker run -d --name=loki --mount type=bind,source="d:\Projects\Grafana\Loki\loki-config.yml",target=/etc/loki/local-config.yaml -p 3100:3100 grafana/loki:3.4.2
```