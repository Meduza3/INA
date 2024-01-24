set terminal pngcairo
set output 'plots/plot_is_comparisons.png'

set xlabel 'n'
set ylabel 'cmp(n)'

set style data points
set key below

plot 'data_insertion.txt' using 1:2 notitle with points pointtype 7 pointsize 1 lc rgb 'gray', \
'data_insertion.txt' using 1:4 notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data_insertion.txt' using 1:4 title 'cmp(n)' with points pointtype 7 pointsize 1 lc rgb 'blue'

set output 'plots/plot_is_comparisons_overn.png'

set xlabel 'n'
set ylabel 'cmp(n)/n'

set style data points
set key below

plot 'data_insertion.txt' using 1:($2/$1) notitle with points pointtype 7 pointsize 1 lc rgb 'gray', \
'data_insertion.txt' using 1:($4/$1)  notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data_insertion.txt' using 1:($4/$1)  title 'cmp(n)/n' with points pointtype 7 pointsize 1 lc rgb 'blue'

set output 'plots/plot_is_comparisons_overnn.png'

set xlabel 'n'
set ylabel 'cmp(n)/n*n'

set style data points
set key below

plot 'data_insertion.txt' using 1:($2/($1*$1)) notitle with points pointtype 7 pointsize 1 lc rgb 'gray', \
'data_insertion.txt' using 1:($4/($1*$1))  notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data_insertion.txt' using 1:($4/($1*$1))  title 'cmp(n)/n*n' with points pointtype 7 pointsize 1 lc rgb 'blue'


set output 'plots/plot_is_switches.png'

set xlabel 'n'
set ylabel 's(n)'

set style data points
set key below

plot 'data_insertion.txt' using 1:3 notitle with points pointtype 7 pointsize 1 lc rgb 'gray', \
'data_insertion.txt' using 1:5 notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data_insertion.txt' using 1:5 title 's(n)' with points pointtype 7 pointsize 1 lc rgb 'purple'

set output 'plots/plot_is_switches_overn.png'

set xlabel 'n'
set ylabel 's(n)/n'

set style data points
set key below

plot 'data_insertion.txt' using 1:($3/$1) notitle with points pointtype 7 pointsize 1 lc rgb 'gray', \
'data_insertion.txt' using 1:($5/$1) notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data_insertion.txt' using 1:($5/$1) title 's(n)/n' with points pointtype 7 pointsize 1 lc rgb 'purple'

set output 'plots/plot_is_switches_overnn.png'

set xlabel 'n'
set ylabel 's(n)/n*n'

set style data points
set key below

plot 'data_insertion.txt' using 1:($3/($1*$1)) notitle with points pointtype 7 pointsize 1 lc rgb 'gray', \
'data_insertion.txt' using 1:($5/($1*$1)) notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data_insertion.txt' using 1:($5/($1*$1)) title 's(n)/n*n' with points pointtype 7 pointsize 1 lc rgb 'purple'
