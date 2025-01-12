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

### Отправить сообщение в очередь:
```curl
localhost:8080/v1/queues/queue2/messages
```

### Добавить подписчика в очередь:
```curl
localhost:8080/v1/queues/queue2/subscriptions
```

## MAINTAINER
* Blckvia