# Server (Go)

## Instalación de Dependencias

Librería `goczmq` (requiere `libczmq-dev`).

1. Instalar la dependencia del sistema:

```bash
sudo apt install libczmq-dev
```

2. Instalar las dependencias de Go:

```bash
go get gopkg.in/zeromq/goczmq.v4
```

## Uso

El servidor hará bind en el puerto `3000` y enviará 100 mensajes

```bash
go run server.go
```
