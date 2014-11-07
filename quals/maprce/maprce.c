#include <stdio.h>
#include <unistd.h>
#include <fcntl.h>
#include <string.h>
#include <sys/mman.h>

int main(int argc, char *argv[]) {
	unsigned char buf[128], *ptr;
	char msg[12];
	int solved, i, fd;

	msg[0] = 'Y';
	msg[1] = 'o';
	msg[2] = 'u';
	msg[3] = ' ';
	msg[4] = 'g';
	msg[5] = 'o';
	msg[6] = 't';
	msg[7] = ' ';
	msg[8] = 'i';
	msg[9] = 't';
	msg[10] = '!';
	msg[11] = '\x00';

	if (argc != 1)
		return 1;

	if (fread(buf, 1, 128, stdin) != 128)
		return 1;
	if ((fd = open(argv[0], O_RDONLY)) == -1)
		return 1;
	if ((ptr = mmap(NULL, 128, PROT_READ, MAP_PRIVATE, fd, 0)) == MAP_FAILED)
		return 1;

	for (i = 0, solved = 1; i < 128; i++)
		solved &= (ptr[i] == buf[i]);
	if (solved)
		puts(msg);

	return 0;
}
