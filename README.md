# Simple HTTP Service on ALT Linux

## Description
A simple HTTP service that shows funny quiz.

## Описание
Простой HTTP-сервис, запустив который можно пройти забавный тест.

## Build
```bash
podman build -t simple-quiz-test-alt:v1.0 -f Containerfile .
```

## Run (Rootless)
```bash
podman run -d --name test -p 8080:8080 simple-quiz-test-alt:v1.0
```

## Requirements
* Podman
* ALT Linux p11 base image
