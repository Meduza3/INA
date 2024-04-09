#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>
#include <linux/limits.h>

#include "myprintf.h"
#include "myscanf.h"

// gcc -m32 -fno-stack-protector *.c -lm

int main(void) {
	char* s;
	myprintf("Enter a string ");
	myscanf("%s", &s);
	int d;
	myprintf("Enter an int ");
	myscanf("%d", &d);
	int b;
	myprintf("Enter a bin ");
	myscanf("%b", &b);
	int x;
	myprintf("Enter a hex ");
	myscanf("%x", &x);

	myprintf("Here you go: %s, %d, %b and %x", s, d, b, x);
	return 0;
}