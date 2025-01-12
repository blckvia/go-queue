#!/bin/sh -e

echo "Запущен скрипт с параметрами: $@"

REPO_NAME="go-queue"

CGO_ENABLED=0

PPROF_UI_PORT=77345
PPROF_DEFAULT_PORT=8080
PPROF_DEFAULT_CPU_DURATION=5

# Подтянуть зависимости
deps(){
  go get ./...
}

# Собрать исполняемый файл
build(){
  deps
  go build ./cmd/order
}

# Собрать docker образ
build_docker() {
  build
  docker build -t "$REPO_NAME:local" .
  rm ./"$REPO_NAME"
}

# Cписок команд
using(){
  echo "Укажите команду при запуске: ./run.sh [command]"
  echo "Список команд:"
  echo "  deps - подтянуть зависимости"
  echo "  build - собрать приложение"
  echo "  build_docker - собрать локальный docker образ"
  echo "  fmt - форматирование кода при помощи 'go fmt'"
  echo "  vet - проверка правильности форматирования кода"
  echo "  pprof_cpu HOST [SECONDS] - сбор метрик нагрузки на cpu из pprof (НЕ НАСТРОЕНО)"
  echo "  pprof_heap HOST - запустить сбор метрик памяти из pprof (НЕ НАСТРОЕНО)"
}

# Запустить контейнер с приложением
compose(){
  docker build -t go-queue:develop .
}

fmt() {
  echo "run go fmt"
  go fmt ./...
}

vet() {
  echo "run go vet"
  go vet ./...
}

# Запустить сбор метрик нагрузки на cpu из pprof
pprof_cpu(){
  local SECS=${3:-$PPROF_DEFAULT_CPU_DURATION}
  local HOST=$2

  go tool pprof -http :$PPROF_UI_PORT $HOST/debug/pprof/profile?seconds=$SECS
}

# Запустить сбор метрик памяти из pprof
pprof_heap(){
  local HOST=$2

  go tool pprof -http :$PPROF_UI_PORT $HOST/debug/pprof/heap
}

############### НЕ МЕНЯЙТЕ КОД НИЖЕ ЭТОЙ СТРОКИ #################

command="$1"
if [ -z "$command" ]
then
 using
 exit 0;
else
 $command $@
fi

