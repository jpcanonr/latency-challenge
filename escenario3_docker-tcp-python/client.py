# client.py
import socket
import time

HOST = "127.0.0.1"
PORT = 8080

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.connect((HOST, PORT))
    start = time.time()
    s.sendall(b"estimulo")
    data = s.recv(1024)
    elapsed = (time.time() - start) * 1000  # ms

    print(f"Respuesta: {data.decode()}")
    print(f"Latencia: {elapsed:.3f} ms")
