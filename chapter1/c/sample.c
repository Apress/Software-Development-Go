#include<stdio.h>
#include<string.h>
#include<sys/socket.h>
#include<arpa/inet.h>
#include<netdb.h>

int main(int argc, char * argv[]) {
  int socket_desc;
  struct sockaddr_in server;
  char * message, server_reply[2000];
  struct hostent * host;
  const char * hostname = "httpbin.org";
  //Create socket
  socket_desc = socket(AF_INET, SOCK_STREAM, 0);
  if (socket_desc == -1) {
    printf("Could not create socket");
  }

  if ((server.sin_addr.s_addr = inet_addr(hostname)) == 0xffffffff) {
    if ((host = gethostbyname(hostname)) == NULL) {
      return -1;
    }

    memcpy( & server.sin_addr, host -> h_addr, host -> h_length);
  }

  // server.sin_addr.s_addr = inet_addr("54.221.78.73");
  server.sin_family = AF_INET;
  server.sin_port = htons(80);

  //Connect to remote server
  if (connect(socket_desc, (struct sockaddr * ) & server, sizeof(server)) < 0) {
    puts("connect error");
    return 1;
  }
  puts("Connected\n");
  //Send some data
  message = "GET / HTTP/1.0\n\n";
  if (send(socket_desc, message, strlen(message), 0) < 0) {
    puts("Send failed");
    return 1;
  }
  puts("Data Send\n");
  //Receive a reply from the server
  if (recv(socket_desc, server_reply, 2000, 0) < 0) {
    puts("recv failed");
  }
  puts("Reply received\n");
  puts(server_reply);
  return 0;
}