set terminal pngcairo
set output 'plots/plot_maxload.png'

set xlabel 'n'
set ylabel 'L_n^(d)'


set style data points
set key below
plot 'data.txt' using 1:($2 == 1 ? $3 : 1/0) notitle with points pointtype 7 pointsize 1 lc rgb 'light-pink', \
'data.txt' using 1:($2 == 1 ? $4 : 1/0) notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data.txt' using 1:($2 == 1 ? $4 : 1/0) title "d = 1" with points pointtype 7 pointsize 1 lc rgb 'red', \
'data.txt' using 1:($2 == 2 ? $3 : 1/0) notitle with points pointtype 7 pointsize 1 lc rgb 'light-cyan', \
'data.txt' using 1:($2 == 2 ? $4 : 1/0) notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data.txt' using 1:($2 == 2 ? $4 : 1/0) title "d = 2" with points pointtype 7 pointsize 1 lc rgb 'blue'

set output 'plots/plot_maxload_extra_1.png'

set xlabel 'n'
set ylabel 'l_n^1/f_1(n)'

set style data points
set key below
plot 'data.txt' using 1:($2 == 1 ? $4/((log($1))/(log(log($1)))) : 1/0) notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data.txt' using 1:($2 == 1 ? $4/((log($1))/(log(log($1)))) : 1/0) notitle with points pointtype 7 pointsize 1.5 lc rgb 'red'

set output 'plots/plot_maxload_extra_2.png'

set xlabel 'n'
set ylabel 'l_n^2/f_1(n)'

set style data points
set key below
plot 'data.txt' using 1:($2 == 2 ? $4/((log(log($1)))/log(2)) : 1/0) notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data.txt' using 1:($2 == 2 ? $4/((log(log($1)))/log(2)) : 1/0) notitle with points pointtype 7 pointsize 1.5 lc rgb 'blue'