#!/bin/bash

#Napisz skrypt w Bashu, który co sekundę prezentuje informacje o systemie operacyjnym. Dane pobierz wykorzystując pseudo system plików proc w
#Linuksie domyślnie zamontowanym w katalogu /proc (patrz man 5 proc) oraz sysfs (patrz man 5 sysfs). Skrypt powinien prezentować następujące informacje:

#Aktualną oraz średnią prędkość wysyłania i odbierania danych z sieci (odczytaj i zinterpretuj 
#/proc/net/dev oraz [[[[[wyświetl w B, KB lub MB w zależność od aktualnej prędkości]]]]]

#Aktualne wykorzystanie rdzeni procesora dla każdego rdzenia osobno w procentach (patrz /proc/stat - man 5 proc) wraz z aktualną częstotliwością (patrz np. /sys/devices/system/cpu/cpu0/cpufreq/scaling_cur_freq dla cpu0) pracy rdzenia procesora (podobnie jak htop) i
#Jak długo system jest uruchomiony w dniach, godzinach, minutach i sekundach (/proc/uptime) i/lub

#Aktualny stan baterii w procentach (/sys/class/power_supply/BAT0/uevent) i
#Obciążenie systemu /proc/loadavg oraz
#Aktualne wykorzystanie pamięci /proc/meminfo (przeanalizuj co oznaczają 3 początkowe wiersze).

while true; do

    if [ "$(cat /proc/net/dev | grep wls1 | awk '{print $10}')" -gt 1000000 ]; then
        echo "U SPEED: $(cat /proc/net/dev | grep wls1 | awk '{print $10/1000000}') MB"
    elif [ "$(cat /proc/net/dev | grep wls1 | awk '{print $10}')" -gt 1000 ]; then
        echo "U SPEED: $(cat /proc/net/dev | grep wls1 | awk '{print $10/1024}') KB"
    else
        echo "U SPEED: $(cat /proc/net/dev | grep wls1 | awk '{print $10}') B"
    fi

    if [ "$(cat /proc/net/dev | grep wls1 | awk '{print $2}')" -gt 1000000 ]; then
        echo "D SPEED: $(cat /proc/net/dev | grep wls1 | awk '{print $2/1000000}') MB"
    elif [ "$(cat /proc/net/dev | grep wls1 | awk '{print $2}')" -gt 1000 ]; then
        echo "D SPEED: $(cat /proc/net/dev | grep wls1 | awk '{print $2/1024}') KB"
    else
        echo "D SPEED: $(cat /proc/net/dev | grep wls1 | awk '{print $2}') B"
    fi

    echo "PC nie ma baterii (404)"
    echo "Obciążenie systemu: $(cat /proc/loadavg | awk {'print $2}')"
    memTotal=$(cat /proc/meminfo | grep MemTotal: | awk '{print $2}')
    memAvailible=$(cat /proc/meminfo | grep MemAvailable: | awk '{print $2}')
    memUsed=$(($memTotal - $memAvailible))
    echo "Wykorzystanie pamięci: $memUsed/$memTotal"
    

    sleep 1
done
