import socket

client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

addr = ('localhost', 8000)

client_socket.connect(addr)

request = b'Hello this is a sample document'
client_socket.send(request)

response = client_socket.recv(1024)
print(response.decode())

client_socket.close()
