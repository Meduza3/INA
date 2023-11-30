#!/bin/bash

processes=($(ls /proc | grep "[0-9]"))

for i in ${processes[@]}; do
    
    if ((i == 1)); then
        echo "PPID PID COMM STATE TTY RSS PGID SID OPNDFILS"
    fi
    if [ -d "/proc/$i" ]; then
        ppid=$(cat /proc/$i/status | grep PPid | awk '{print $2}')
        pid=$(cat /proc/$i/status | grep -w Pid | awk '{print $2}')
        comm=$(cat /proc/$i/status | grep -w Name | awk '{print $2}')
        state=$(cat /proc/$i/status | grep -w State | awk '{print $2}')
        tty=$(cat /proc/$i/status | grep -w Tty | awk '{print $2}')
        rss=$(cat /proc/$i/status | grep -w VmRSS | awk '{print $2}')
        pgid=$(cat /proc/$i/status | grep -E "^NSpgid" | awk '{print $2}')
        sid=$(cat /proc/$i/status | grep -E "^NSsid" | awk '{print $2}')
        opndfils=$(find /proc/$i/fd/ -maxdepth 1 -type f | wc -l)

        echo "$ppid $pid $comm $state $tty $rss $pgid $sid $opndfils"
    fi
    

done | column -t