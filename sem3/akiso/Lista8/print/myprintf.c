#include "myprintf.h"

void myprintf(char* pattern, ...) {
	char *p = (char *) &pattern + sizeof(pattern);
	
	for(int i = 0; i < strlen(pattern); i++) {
		if (pattern[i] == '%' && i < strlen(pattern) - 1) {
			char output[1024];
			switch (pattern[i+1]) {

				case 'd': {
					int* number = ((int *)p); // retrieve the next argument
					p += sizeof(int);
					changeBase(output, *number, 10);
					write(1, output, strlen(output));
					i+=2;
					break;
				}
				case 's': {
					char* text  = *((char **) p);
					p += sizeof(char*);
					while (text[0] != '\0') {
						write(1, text, sizeof(char));
						text++;
					}
					i+=2;
					break;
				}
				case 'x': {
					int number = *((int *)p);
					p += sizeof(int);
					changeBase(output, number, 16);
					write(1, output, strlen(output));
					i+=2;
					break;
				}
				case 'b': {
					int number = *((int *)p);
					p += sizeof(int);
					changeBase(output, number, 2);
					write(1, output, strlen(output));
					i+=2;
					break;
				}
				default: {
					write(1, pattern[i], 1);
					break;
				}
			}
		}
		int a = write(1, &pattern[i], 1);
	}
	p = NULL;
	write(1, "\n\0", 3);
}

void changeBase(char* out, int number, int base) {
    //Convert a number to a string reppresentation of that number in another base - Writes results in out
	int i = 127;
	int j = 0;
	if (number < 0) {
		number = -number;
		j = 1;
		out[0] = '-';
	}
	do {
		out[i] = "0123456789ABCDEF"[number % base];
		i--;
		number = number/base;
	} while (number > 0);
	while (++i < 128) { //reveres order
		out[j++] = out[i];
	}
	out[j] = '\0';
}