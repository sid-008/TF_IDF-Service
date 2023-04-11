import socket

client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

addr = ('localhost', 8000)

client_socket.connect(addr)
print("Connected to server...")

with open('file1.txt', 'r') as file:
    contents = file.read()

client_socket.sendall(contents.encode())

response = client_socket.recv(1024)
print(response.decode())

client_socket.close()
