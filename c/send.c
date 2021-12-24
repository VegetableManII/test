#include<stdio.h>
#include<stdlib.h>
#include<string.h>
#include<sys/types.h>
#include<sys/socket.h>
#include<sys/wait.h>
#include<netinet/in.h>
#include<arpa/inet.h>
#include<errno.h>
#include <unistd.h>
#include <time.h>
int main(){
	char msg[128] = "I am broadCast message from server!";
	int brdcFd;
	if((brdcFd = socket(AF_INET, SOCK_DGRAM, 0)) == -1){
		printf("socket fail\n");
		return -1;
	}
	int optval = 1;//这个值一定要设置，否则可能导致sendto()失败
	setsockopt(brdcFd, SOL_SOCKET, SO_BROADCAST, &optval, sizeof(int));
	setsockopt(brdcFd, SOL_SOCKET, SO_REUSEADDR, &optval, sizeof(int));

	struct sockaddr_in theirAddr;
	memset(&theirAddr, 0, sizeof(struct sockaddr_in));
	theirAddr.sin_family = AF_INET;
	theirAddr.sin_addr.s_addr = inet_addr("255.255.255.255");
	theirAddr.sin_port = htons(4001);
	int sendBytes;
	int recvbytes;
	char recvbuf[128];
	while (1)
	{
		if((sendBytes = sendto(brdcFd, msg, strlen(msg), 0,	(struct sockaddr *)&theirAddr, sizeof(struct sockaddr))) == -1){
			printf("sendto fail, errno=%d\n", errno);
			break;
		}
		printf("msg=%s, msgLen=%d, sendBytes=%d\n", msg, strlen(msg), sendBytes);
		int addrLen = sizeof(struct sockaddr_in);
		if((recvbytes = recvfrom(brdcFd, recvbuf, 128, 0,(struct sockaddr *)&theirAddr, &addrLen)) != -1){
			recvbuf[recvbytes] = '\0';
			printf("receive a response messgse:%s\n", recvbuf);
		}else{
			printf("recvfrom fail\n");
		}
		sleep(1);
	}
	close(brdcFd);
	return 0;
}
