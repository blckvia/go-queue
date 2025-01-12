# GO-QUEUE

## CMD

## Запуск рутинных операций
Для запуска рутинных операций, такие как статические анализаторы, тесты, всякие генерации и т.п. существует скрипт
run.sh, посмотрите его using, чтобы ознакомиться с его возможностями. 
```bash
./run.sh using
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