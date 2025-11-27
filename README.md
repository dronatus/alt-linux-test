# Simple HTTP Service on ALT Linux

## Description
A simple HTTP service that shows funny quiz.

## Описание
Простой HTTP-сервис, запустив который можно пройти забавный тест.

## License
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

### Run (Rootless)
```bash
podman run -d --name test -p 8080:8080 simple-quiz-test-alt:myself
```

## Kubernetes Deployment

### Requirements
* Git
* k8s (Minikube)
* kubectl
* curl (or other browser)

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
kubectl port-forward service/test-service 8080:61111 &
curl http://localhost:8080
```
