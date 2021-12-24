#include<stdlib.h>
#include<stdio.h>
#include<string.h>
#include<sys/types.h>
#include<netinet/in.h>
#include<netdb.h>
#include<sys/socket.h>
#include<sys/wait.h>
#include<arpa/inet.h>
#include<errno.h>
#include <unistd.h>
#include <time.h>
int main(){
	int sockListen;
	if((sockListen = socket(AF_INET, SOCK_DGRAM, 0)) == -1){
		printf("socket fail\n");
		return -1;
	}
	int set = 1;
	setsockopt(sockListen, SOL_SOCKET, SO_REUSEADDR, &set, sizeof(int));
	struct sockaddr_in recvAddr;
	memset(&recvAddr, 0, sizeof(struct sockaddr_in));
	recvAddr.sin_family = AF_INET;
	recvAddr.sin_port = htons(4001);
	recvAddr.sin_addr.s_addr = INADDR_ANY;
	// 必须绑定，否则无法监听
	if(bind(sockListen, (struct sockaddr *)&recvAddr, sizeof(struct sockaddr)) == -1){
		printf("bind fail\n");
		return -1;
	}
    int sendBytes;
	int recvbytes;
    char msg[128] = "I am broadCast message from client!";
	char recvbuf[128];
	int addrLen = sizeof(struct sockaddr_in);
    while (1)
    {
        if((recvbytes = recvfrom(sockListen, recvbuf, 128, 0,(struct sockaddr *)&recvAddr, &addrLen)) != -1){
		    recvbuf[recvbytes] = '\0';
		    printf("receive a broadCast messgse:%s\n", recvbuf);
	    }else{
	    	printf("recvfrom fail\n");
    	}
        if((sendBytes = sendto(sockListen, msg, strlen(msg), 0,	(struct sockaddr *)&recvAddr, sizeof(struct sockaddr))) == -1){
			printf("sendto fail, errno=%d\n", errno);
			break;
		}
		printf("msg=%s, msgLen=%d, sendBytes=%d\n", msg, strlen(msg), sendBytes);
        sleep(1);
    }
	close(sockListen);
	return 0;
}
