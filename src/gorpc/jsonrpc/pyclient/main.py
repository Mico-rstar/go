import socket 
import json

request = {
    "id": 0,
    "params": ["python greet"],
    "method": "HelloService.Hello"
}

client = socket.create_connection(("localhost", 10010))
client.sendall(json.dumps(request).encode())

rsp = client.recv(1024)
rsp = json.loads(rsp.decode())

print(rsp)