#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <fcntl.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <signal.h>

#define STRSZ 4096
#define XORSZ 32
#define ADDR "0.0.0.0"
#define PORT 6868
#define CONNQUEUE 128

int cat_flag();
void print_banner();
int read_len();
int read_str(char *str, int len);
int crypt_str(char *str, int len);
void loop();

int client;
FILE *fdclient;

int cat_flag() {
	FILE *fd;
	char buf[STRSZ];

	fd = fopen("flag.txt", "r");
	if (fd == NULL) {
		fprintf(fdclient, "error: cannot open flag.txt\n");
		fflush(fdclient);
		return -1;
	}
	fgets(buf, STRSZ, fd);
	fprintf(fdclient, "%s", buf);
	fflush(fdclient);
	fclose(fd);

	return 0;
}

void print_banner() {
	fprintf(fdclient,
			"#############################################################\n"
			"#                 Welcome to crypt.nsa.gov                  #\n"
			"#        All connections are monitored and recorded         #\n"
			"# Disconnect IMMEDIATELY if you are not an authorized user! #\n"
			"#############################################################\n\n");
	fflush(fdclient);
}

int read_len() {
	int len = 0;

	if (read(client, &len, sizeof(len)) != sizeof(len))
		return -1;

	if (len >= STRSZ) {
		fprintf(fdclient, "error: length >= %i\n", STRSZ);
		fflush(fdclient);
		return -1;
	}
	return len;
}

int read_str(char *str, int len) {
	if (read(client, str, len) != len)
		return -1;
	return 0;
}

int crypt_str(char *str, int len) {
	unsigned char *ptr;
	int i;

	static int first = 1;
	static unsigned char key[XORSZ];

	if (first) {
		int fd = open("/dev/urandom", O_RDONLY);
		if (read(fd, key, XORSZ) != XORSZ) {
			fprintf(fdclient, "error: cannot open /dev/urandom\n");
			fflush(fdclient);
			return -1;
		}
		close(fd);
		first = 0;
	}

	ptr = (unsigned char*)(str);
	for (i = 0; i < len; i++) {
		ptr[i] ^= key[i % XORSZ]; 
	}
	ptr[i] = '\00';

	write(client, &len, sizeof(len));
	write(client, str, len);
	fprintf(fdclient, "[DEBUG] puts@GOT=0x804a374 ; cat_flag=0x804882b\n");
	fprintf(fdclient, "[DEBUG] Msg = ");
	fprintf(fdclient, str);
	fprintf(fdclient, "\n");
	fprintf(fdclient, "\nMessage successfully encrypted.\n");
	fflush(fdclient);
	printf("[DEBUG] End of crypt_str()\n");
	return 0;
}

void loop() {
	char str[STRSZ];
	int len;

	print_banner();
	for (;;) {
		if ((len = read_len()) == -1)
			break;
		if (read_str(str, len) == -1)
			break;
		if (crypt_str(str, len) == -1)
			break;
	}
	fprintf(fdclient, "Bye\n");
	fflush(fdclient);
}

int main(int argc, char **argv) {
	struct sockaddr_in srvaddr, cliaddr;
	int sock, pid;
	socklen_t cliaddrlen = sizeof(cliaddr);

	memset(&srvaddr, '0', sizeof(srvaddr));
	srvaddr.sin_family = AF_INET;
	srvaddr.sin_port = htons(PORT);
	srvaddr.sin_addr.s_addr = inet_addr(ADDR);

	if ((sock = socket(AF_INET, SOCK_STREAM, 0)) == -1) {
		perror("socket");
		return 1;
	}

	if (bind(sock, (const struct sockaddr*)&srvaddr, sizeof(srvaddr)) == -1) {
		perror("bind");
		return 1;
	}

	if (listen(sock, CONNQUEUE) == -1) {
		perror("listen");
		return 1;
	}

	for(;;) {
		client = accept(sock, (struct sockaddr *)&cliaddr, &cliaddrlen);
		pid = fork();

		if(pid == 0) {
			alarm(60);
			if ((fdclient = fdopen(client, "r+")) == NULL) {
				perror("fdopen");
				close(client);
				break;
			}
			loop();
			fclose(fdclient);
			close(client);
			break;
		} else {
			close(client);
			continue;
		}
	}
	close(sock);
	return 0;
}
