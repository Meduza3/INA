#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>
#include <linux/limits.h>

#include "myscanf.h"

int myscanf(const char* pattern, ...) {
	char* p = (char *) &pattern + sizeof pattern; // Pointer  to the start of the argument list
	
	char* input = calloc(1024, sizeof(char)); //Allocates a space and fills it with next line
	int size = read(0, input, 1024); //Read 1024 characters from standard input
	if (input[size - 1] == '\n') {
        input[size-1] = '\0'; //Remove new line at end line
        } 
	input[size] = '\0'; //Make sure it is null terminated
	
	if (!strcmp(pattern, "%d")) { 
		int* in = (int*)(*(int*)p); 
		p += sizeof(int*);
		*in = StringToInt(input, 10);
	} else if (!strcmp(pattern, "%s")) {
		char** in = (char **)(*(char**) p);
		*in = input;
		p += sizeof(*in);
	} else if (!strcmp(pattern, "%x")) {
		int* in = (int*)(*(int*)p);
		p += sizeof(int*) = StringToInt(input, 16);
	} else if (!strcmp(pattern, "%b")) {
		int* in = (int*)(*(int*)p);
		p += sizeof(int*);
		*in = StringToInt(input, 2);
	}
	p = NULL;
	return size;
}

int StringToInt(char* string, int base) {
	char numbers[17] = "0123456789ABCDEF";
	int result = 0;
	int negative = 0;
	if (string[0] == '-') {
        negative = 1;
    }

	int length = strlen(string);
	for (int i = length - 1; i > negative - 1; i--) {
		for (int j = 0; j < base + 1; j++) {
			if (string[i] == numbers[j]) {
				result += j*pow(base, length - 1 - i);
			}
		}
	}
	if (negative) {
        return result*-1;
    } else {
        return result;
    }
}