#include <stdio.h>
#include <stdlib.h>
#include <signal.h>

//Wysyłanie killa do inita

int main() {
    if(kill(1, SIGKILL) == -1) {
        perror("kill");
        exit(1);
    }

    printf("SIGKILL wysłany do procesu init.\n");
}