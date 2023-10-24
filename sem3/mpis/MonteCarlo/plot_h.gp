set terminal pngcairo enhanced font 'Verdana,10'
set output 'plots/plot_h.png'

set xlabel 'n'
set ylabel 'Value'
set title 'Integral from 0 to 1 of h(x)dx'

set style data points

set key bottom right

plot 'data/data_h.txt' using 1:2 title 'Experiment Results' with points lc rgb 'gray', \
    'data/data_h.txt' using 1:3 title 'Average' with points pt 7 lc rgb 'red', \
    0.2 with lines lt 2 lw 2 lc rgb 'blue' title 'I'