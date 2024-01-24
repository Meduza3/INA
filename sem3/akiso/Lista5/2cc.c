#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

volatile sig_atomic_t received_signals = 0;

void signal_handler(int signal) {
    if (signal == SIGUSR1) {
        received_signals++;
        printf("Odebrano sygnał SIGUSR1, razem: %d\n", received_signals);
    }
}

int main() {
    struct sigaction sa;
    sa.sa_handler = signal_handler;
    sigemptyset(&sa.sa_mask);
    sa.sa_flags = 0;

    if (sigaction(SIGUSR1, &sa, NULL) == -1) {
        perror("Nie można ustawić obsługi sygnału SIGUSR1");
        exit(EXIT_FAILURE);
    }

    printf("Proces gotowy do odbierania sygnałów SIGUSR1. PID: %d\n", getpid());

    // Czekaj na sygnały w nieskończonej pętli
    while (1) {
        pause();
    }

    return 0;
}
