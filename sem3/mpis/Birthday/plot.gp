set terminal pngcairo
set output 'plots/plot_b_n.png'

set xlabel 'n'
set ylabel 'Value'

set style data points
set key below

plot 'data/data.txt' using 1:2 notitle with points pointtype 7 pointsize 1 lc rgb 'gray', \
'data/data.txt' using 1:7 notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data/data.txt' using 1:7 title 'B_n' with points pointtype 7 pointsize 1 lc rgb 'red'

set output 'plots/plot_b_over_n.png'
plot 'data/data.txt' using 1:($2/$1) notitle with points pointtype 7 pointsize 1 lc rgb 'gray', \
'data/data.txt' using 1:($7/$1) notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data/data.txt' using 1:($7/$1) title 'B_n / n' with points pointtype 7 pointsize 1 lc rgb 'red'

set output 'plots/plot_b_over_skrr_n.png'
plot 'data/data.txt' using 1:($2/sqrt($1)) notitle with points pointtype 7 pointsize 1 lc rgb 'gray', \
'data/data.txt' using 1:($7/sqrt($1)) notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data/data.txt' using 1:($7/sqrt($1)) title 'B_n / sqrt(n)' with points pointtype 7 pointsize 1 lc rgb 'red'

set output 'plots/plot_u_n.png'

plot 'data/data.txt' using 1:3 notitle with points pointtype 7 pointsize 1 lc rgb 'gray', \
'data/data.txt' using 1:8 notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data/data.txt' using 1:8 title 'U_n' with points pointtype 7 pointsize 1 lc rgb 'purple'

set output 'plots/plot_u_over_n.png'

plot 'data/data.txt' using 1:($3/$1) notitle with points pointtype 7 pointsize 1 lc rgb 'gray', \
'data/data.txt' using 1:($8/$1) notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data/data.txt' using 1:($8/$1) title 'U_n / n' with points pointtype 7 pointsize 1 lc rgb 'purple'

set output 'plots/plot_c_n.png'

plot 'data/data.txt' using 1:4 notitle with points pointtype 7 pointsize 1 lc rgb 'gray', \
'data/data.txt' using 1:9 notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data/data.txt' using 1:9 title 'C_n' with points pointtype 7 pointsize 1 lc rgb 'green'

set output 'plots/plot_c_over_n.png'

plot 'data/data.txt' using 1:($4/$1) notitle with points pointtype 7 pointsize 1 lc rgb 'gray', \
'data/data.txt' using 1:($9/$1) notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data/data.txt' using 1:($9/$1) title 'C_n over n' with points pointtype 7 pointsize 1 lc rgb 'green'

set output 'plots/plot_c_over_nlogn.png'

plot 'data/data.txt' using 1:($4/($1 * log($1))) notitle with points pointtype 7 pointsize 1 lc rgb 'gray', \
'data/data.txt' using 1:($9/($1 * log($1))) notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data/data.txt' using 1:($9/($1 * log($1))) title 'C_n over n*ln(n)' with points pointtype 7 pointsize 1 lc rgb 'green'

set output 'plots/plot_c_over_nsquared.png'

plot 'data/data.txt' using 1:($4/($1 * $1)) notitle with points pointtype 7 pointsize 1 lc rgb 'gray', \
'data/data.txt' using 1:($9/($1 * $1)) notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data/data.txt' using 1:($9/($1 * $1)) title 'C_n over n*n' with points pointtype 7 pointsize 1 lc rgb 'green'

set output 'plots/plot_d_n.png'

plot 'data/data.txt' using 1:5 notitle with points pointtype 7 pointsize 1 lc rgb 'gray', \
'data/data.txt' using 1:10 notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data/data.txt' using 1:10 title 'D_n' with points pointtype 7 pointsize 1 lc rgb 'cyan'

set output 'plots/plot_d_over_n.png'

plot 'data/data.txt' using 1:($5/$1) notitle with points pointtype 7 pointsize 1 lc rgb 'gray', \
'data/data.txt' using 1:($10/$1) notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data/data.txt' using 1:($10/$1) title 'D_n over n' with points pointtype 7 pointsize 1 lc rgb 'cyan'

set output 'plots/plot_d_over_nlogn.png'

plot 'data/data.txt' using 1:($5/($1 * log($1))) notitle with points pointtype 7 pointsize 1 lc rgb 'gray', \
'data/data.txt' using 1:($10/($1 * log($1))) notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data/data.txt' using 1:($10/($1 * log($1))) title 'D_n over n*log(n)' with points pointtype 7 pointsize 1 lc rgb 'cyan'

set output 'plots/plot_d_over_nsquared.png'

plot 'data/data.txt' using 1:($5/($1 * $1)) notitle with points pointtype 7 pointsize 1 lc rgb 'gray', \
'data/data.txt' using 1:($10/($1 * $1)) notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data/data.txt' using 1:($10/($1 * $1)) title 'D_n over n*n' with points pointtype 7 pointsize 1 lc rgb 'cyan'

set output 'plots/plot_dc_n.png'

plot 'data/data.txt' using 1:6 notitle with points pointtype 7 pointsize 1 lc rgb 'gray', \
'data/data.txt' using 1:11 notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data/data.txt' using 1:11 title 'D_n - C_n' with points pointtype 7 pointsize 1 lc rgb 'orange'

set output 'plots/plot_dc_over_n.png'

plot 'data/data.txt' using 1:($6/$1) notitle with points pointtype 7 pointsize 1 lc rgb 'gray', \
'data/data.txt' using 1:($11/$1) notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data/data.txt' using 1:($11/$1) title 'D_n - C_n over n' with points pointtype 7 pointsize 1 lc rgb 'orange'

set output 'plots/plot_dc_over_nlogn.png'

plot 'data/data.txt' using 1:($6/($1 * log($1))) notitle with points pointtype 7 pointsize 1 lc rgb 'gray', \
'data/data.txt' using 1:($11/($1 * log($1))) notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data/data.txt' using 1:($11/($1 * log($1))) title 'D_n - C_n over n log n' with points pointtype 7 pointsize 1 lc rgb 'orange'

set output 'plots/plot_dc_over_nloglogn.png'

plot 'data/data.txt' using 1:($6/($1 * log(log($1)))) notitle with points pointtype 7 pointsize 1 lc rgb 'gray', \
'data/data.txt' using 1:($11/($1 * log(log($1)))) notitle with points pointtype 7 pointsize 1.5 lc rgb 'black', \
'data/data.txt' using 1:($11/($1 * log(log($1)))) title 'D_n - C_n over n log log n' with points pointtype 7 pointsize 1 lc rgb 'orange'