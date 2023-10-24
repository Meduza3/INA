set terminal pngcairo enhanced font 'Verdana,10'
set output 'plots/plot_g.png'

set xlabel 'n'
set ylabel 'Value'
set title 'Integral from 0 to π of g(x)dx'

set style data points

set key bottom right

plot 'data/data_g.txt' using 1:2 title 'Experiment Results' with points lc rgb 'gray', \
    'data/data_g.txt' using 1:3 title 'Average' with points pt 7 lc rgb 'red', \
    2 with lines lt 2 lw 2 lc rgb 'blue' title 'I'