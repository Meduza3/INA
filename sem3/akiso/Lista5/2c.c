#include <stdio.h>
#include <stdlib.h>
#include <signal.h>
#include <unistd.h>

int main(int argc, char *argv[]) {
    if (argc != 2) {
        fprintf(stderr, "Użycie: %s <PID>\n", argv[0]);
        exit(EXIT_FAILURE);
    }

    pid_t pid = atoi(argv[1]);

    for (int i = 0; i < 10; i++) {
        kill(pid, SIGUSR1);
    }

    printf("Wysłano 10 sygnałów SIGUSR1 do procesu %d.\n", pid);
    return 0;
}
