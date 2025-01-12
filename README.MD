# GO-QUEUE

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
