# client.py
import requests
import time

url = "http://127.0.0.1:8080"

start = time.time()
resp = requests.get(url)
elapsed = (time.time() - start) * 1000  # ms

print(f"Respuesta: {resp.text}")
print(f"Latencia: {elapsed:.3f} ms")
