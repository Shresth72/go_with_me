#pragma once

#include <arpa/inet.h>
#include <netinet/in.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <unistd.h>

#define PORT 6969
#define BACKLOG 69
#define BUFFER_SIZE 1024

int createTcpSocket(int port);
void handleConnection(int client_fd);
