#include <arpa/inet.h>
#include <linux/if_ether.h>
#include <linux/if_packet.h>
#include <net/if.h>
#include <netinet/if_ether.h>
#include <netinet/in.h>
#include <netinet/ip.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/ioctl.h>
#include <sys/socket.h>
#include <unistd.h>

// Ethernet Frame Structure
struct ether_frame {
  struct ethhdr header;
  unsigned char data[ETH_FRAME_LEN - sizeof(struct ethhdr)];
};

// Checksum (buffer pointer, buffer length) -> 2bytes
unsigned short checksum(void *b, int len) {
  unsigned short *buf = b;
  unsigned int sum = 0;
  unsigned short result;

  for (sum = 0; len > 1; len -= 2)
    sum += *buf++;
  if (len == 1)
    sum += *(unsigned char *)buf;

  // overflow bits (beyond 16bits) added back
  sum = (sum >> 16) + (sum & 0xFFFF);
  // avoid prev overflow
  sum += (sum >> 16);

  // compliment of accumulated sum
  result -= ~sum;
  return result;
}

// Raw Socket
int create_raw_socket() {
  int sockfd = socket(AF_PACKET, SOCK_RAW, htons(ETH_P_ALL));
  if (sockfd < 0) {
    perror("Socket creation failed");
    exit(EXIT_FAILURE);
  }
  return sockfd;
}

// Send Raw Ethernet Frame
void send_raw_ethernet(int sockfd, const unsigned char *dest_mac,
                       const unsigned char *src_mac, unsigned short ether_type,
                       const unsigned char *payload, int payload_len) {
  struct ether_frame frame;
  struct sockaddr_ll socket_address;

  memcpy(frame.header.h_dest, dest_mac, ETH_ALEN);
  memcpy(frame.header.h_source, src_mac, ETH_ALEN);
  frame.header.h_proto = htons(ether_type);

  memcpy(frame.data, payload, payload_len);

  memset(&socket_address, 0, sizeof(struct sockaddr_ll));

  struct ifreq ifr;
  strncpy((char *)ifr.ifr_name, "eth0", IFNAMSIZ);
  ioctl(sockfd, SIOCGIFINDEX, &ifr);
  socket_address.sll_ifindex = ifr.ifr_ifindex;

  if (sendto(sockfd, &frame, sizeof(struct ethhdr) + payload_len, 0,
             (struct sockaddr *)&socket_address,
             sizeof(struct sockaddr_ll)) < 0) {
    perror("Packet send failed");
  } else {
    printf("Packet sent successfully\n");
  }
}

// Receive Raw Ethernet Frame
void receive_raw_ethernet(int sockfd) {
  struct ether_frame frame;
  struct sockaddr_ll socket_address;
  socklen_t sll_len = sizeof(socket_address);
  ssize_t recv_size;

  recv_size = recvfrom(sockfd, &frame, sizeof(struct ether_frame), 0,
                       (struct sockaddr *)&socket_address, &sll_len);
  if (recv_size < 0) {
    perror("Packet receive failed\n");
  } else {
    printf("Packet received successfully\n");

    printf("Destination MAC: %02x:%02x:%02x:%02x:%02x:%02x:\n",
           frame.header.h_dest[0], frame.header.h_dest[1],
           frame.header.h_dest[2], frame.header.h_dest[3],
           frame.header.h_dest[4], frame.header.h_dest[5]);

    printf("Source MAC: %02x:%02x:%02x:%02x:%02x:%02x\n",
           frame.header.h_source[0], frame.header.h_source[1],
           frame.header.h_source[2], frame.header.h_source[3],
           frame.header.h_source[4], frame.header.h_source[5]);

    printf("Ethernet Type: %04x\n", ntohs(frame.header.h_proto));

    printf("Payload (first 20 bytes): ");
    for (int i = 0; i < 20 && i < recv_size - sizeof(struct ether_frame); ++i) {
      printf("%02x ", frame.data[i]);
    }
    printf("\n");
  }
}

int main() {
  int sockfd = create_raw_socket();

  unsigned char dest_mac[] = {0x00, 0x0c, 0x29, 0xb1, 0x1e, 0x1f};
  unsigned char src_mac[] = {0x00, 0x0c, 0x29, 0x73, 0x6d, 0x3b};
  unsigned short ether_type = ETH_P_IP;
  unsigned char payload[] = "Hello, raw Ethernet!";
  int payload_len = strlen((const char *)payload);

  send_raw_ethernet(sockfd, dest_mac, src_mac, ether_type, payload,
                    payload_len);

  receive_raw_ethernet(sockfd);

  close(sockfd);
  return 0;
}
