#include "echo.h"
#include <unistd.h>

int main() {
  int server_fd = createTcpSocket(PORT);
  struct sockaddr_in address;
  int addrlen = sizeof(address);

  printf("Server is listening on port %d\n", PORT);

  char buffer[BUFFER_SIZE];

  while (1) {
    int client_fd =
        accept(server_fd, (struct sockaddr *)&address, (socklen_t *)&addrlen);
    if (client_fd < 0) {
      perror("accept failed");
      close(server_fd);
      exit(EXIT_FAILURE);
    }

    handleConnection(client_fd);
  }

  close(server_fd);
  return 0;
}

void handleConnection(int client_fd) {
  char buffer[BUFFER_SIZE];

  ssize_t bytes_read = read(client_fd, buffer, BUFFER_SIZE - 1);
  if (bytes_read < 0) {
    perror("read failed");
    close(client_fd);
    return;
  }

  ssize_t bytes_written = write(client_fd, buffer, bytes_read);
  if (bytes_written < 0) {
    perror("write failed");
  }

  close(client_fd);
}
