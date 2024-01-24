// Napisz program w języku C, który uruchomi powłokę (Bash) z prawami roota. Po kompilacji programu można ustawić (z poziomu roota) dowolne atrybuty 
// (np. patrz SUID). Następnie już z poziomu dowolnego użytkownika uruchamiając program uruchamia się konsola administratora, 
// podobnie jak sudo /bin/bash (bez wprowadzania hasła). Oczywiście proszę nie wykorzystywać programu 'sudo' we własnym programie!

#include <stdlib.h>
#include <unistd.h>


int main(void) {

    setuid(0);

    char *bash[2] = {"/bin/bash", NULL};

    execvp("bash", bash);
    perror("execvp");
    exit(1);

    return 0;
}