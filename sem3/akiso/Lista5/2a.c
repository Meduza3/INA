#include <signal.h>
#include <stdio.h>
#include <unistd.h>

//Program, który obsługuje wszystkie sygnały

void signal_handler(int signal) {
    printf("Signal %d\n", signal);
}

int main() {
    struct sigaction sa;
    sa.sa_handler = signal_handler;
    sigemptyset(&sa.sa_mask);
    sa.sa_flags = 0;

    for (int i = 1; i < _NSIG; i++) {
        if(sigaction(i, &sa, NULL) == -1) {
            perror("sigaction");
            printf("Nie można obsłużyć sygnału: %d\n", i);
        }
    }

    while(1) {
        pause();
    }

    return 0;
}