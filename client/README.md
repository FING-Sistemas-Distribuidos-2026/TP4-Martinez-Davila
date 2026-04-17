## Instalación de Dependencias

Librería `pyzmq`.

1. Crear y activar un entorno virtual:

```bash
python3 -m venv zmq_env
source zmq_env/bin/activate
```

2. Instalar las dependencias:

```bash
pip install -r requirements.txt
```

## Uso

El cliente se conectará por defecto al puerto `3000`.

```bash
python client.py <SERVER_IP>
```

**Ejemplo**:

```bash
python client.py 127.0.0.1
```
