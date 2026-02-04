# server.py
import socket

HOST = "0.0.0.0"   # Escuchar en todas las interfaces
PORT = 8080

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.bind((HOST, PORT))
    s.listen()
    print(f"Servidor TCP escuchando en puerto {PORT}...")
    while True:
        conn, addr = s.accept()
        with conn:
            data = conn.recv(1024)
            if not data:
                continue
            conn.sendall(b"respuesta")
            print(f"Recibido: {data.decode()}")
