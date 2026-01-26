import requests

request = {
    "id": 0,
    "params": ["hhh"],
    "method": "HelloService.Hello"
}


rsp = requests.post("http://localhost:10010/jsonrpc", json=request)
print(rsp.text)