import sys
import zmq

def run(server_host):
    context = zmq.Context()
    sock = context.socket(zmq.PULL)

    sock.connect(f"tcp://{server_host}:3000")
    print(f"Worker conectado (?) al server {server_host}")

    while True:
        msg = sock.recv_string()
        print(f'Procesado "{msg}"')

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Uso: python client.py <SERVER_IP>")
        sys.exit(1)
    
    host = sys.argv[1]
    run(host)
