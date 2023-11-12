#include <stdio.h>
#include <unistd.h>

int main() {
	for(int i = 0; i < 257; i++) {
        usleep(10000);
		printf("\e[38;5;%dmHello, World!\n", i);
	}
	printf("\e[0m");
	return 0;
}