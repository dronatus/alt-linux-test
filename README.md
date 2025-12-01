# Simple HTTP Service on ALT Linux

### Description
A simple HTTP service that shows funny quiz.

### Описание
Простой HTTP-сервис, запустив который можно пройти забавный тест.

### License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Public Image

This image is available at:
- **GitHub Container Registry**: `ghcr.io/dronatus/simple-quiz-test-alt:latest`
- **Pull command**: `podman pull ghcr.io/dronatus/simple-quiz-test-alt:latest`

## Build Locally

### Requirements
* Git
* Podman

### Clone
```bash
git clone https://github.com/dronatus/alt-linux-test && \
cd alt-linux-test
```

### Build
```bash
podman build -t simple-quiz-test-alt:myself -f Containerfile .
```

## Run using Podman (Rootless)
```bash
podman run -d --name test -p 8000:8080 simple-quiz-test-alt
```

## Kubernetes Deployment

### Requirements
* Git
* k8s (Minikube)
* kubectl
* make (for Using Makefile)
* Utilities: curl, ss

### Clone
```bash
git clone https://github.com/dronatus/alt-linux-test && \
cd alt-linux-test
```

### Deploy Application
```bash
kubectl apply -f k8s/
```

### Test
```bash
kubectl get pods # STATUS: Running, READY: 1/1
kubectl port-forward service/test-service 8000:61111 &
curl http://localhost:8000
```

### Using Makefile

```bash
make apply       # развернуть
make delete      # удалить
make describe    # детальная информация
make status      # проверить статус развертывания
make test        # протестировать
make run         # протестировать, затем запустить проброс портов для браузера
make stop-run    # остановить проброс портов
make status-run  # проверить статус проброса портов
make clean       # очистить временные файлы
```

## Tests && Founded Bugs:

- В Makefile добавлены зависимости, чтобы тесты завершались корректно.
- Протестировано на ALT Linux P11 и Ubuntu 22.04


### Anyway, have fun!