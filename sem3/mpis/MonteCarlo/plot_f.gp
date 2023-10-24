set terminal pngcairo
set output 'plots/plot_f.png'

set xlabel 'n'
set ylabel 'Value'
set title 'Integral from 0 to 8 of f(x)dx'

set style data points

set key bottom right

plot 'data/data_f.txt' using 1:2 title 'Experiment Results' with points lc rgb 'gray', \
    'data/data_f.txt' using 1:3 title 'Average' with points pt 7 lc rgb 'red', \
    12 with lines lt 2 lw 2 lc rgb 'blue' title 'I'