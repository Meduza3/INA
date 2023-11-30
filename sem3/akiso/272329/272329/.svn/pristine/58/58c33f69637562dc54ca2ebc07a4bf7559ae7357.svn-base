#!/bin/bash

declare -A hasze

for file in $(find $1 -type f); do
    hasz=$(sha256sum $file | awk '{print $1}')
    rozmiar=$(stat -c%s "$file")
    nazwa=$(sha256sum $file | awk '{print $2}')

    hasze[$hasz]="${hasze[$hasz]} $nazwa"
    #echo ${hasze[$hasz]}
done

for i in "${!hasze[@]}"; do
    hasz="${hasze[$i]}"
    ilosc_slow=$(echo "$hasz" | wc -w)

    if [ $ilosc_slow -ge 2 ]; then
        echo $hasz
    fi
done | sort -n -k1
