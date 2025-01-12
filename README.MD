# GO-QUEUE

## Запуск рутинных операций
Для запуска рутинных операций, такие как статические анализаторы, тесты, всякие генерации и т.п. существует скрипт
run.sh, посмотрите его help, чтобы ознакомиться с его возможностями.

## CMD

### Запуск приложение локально
```bash
go build gtihub.com/blckvia/go-queue/cmd/app
```

### Запуск контейнеров для разработки API
```shell script
docker build -t go-queue:develop .
```
```shell script
docker-compose -f deployments/docker-compose.yaml up -d
```
