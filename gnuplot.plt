set title "Bitcoin transaction data"
set datafile separator ","
set xdata time
set timefmt "%Y-%m-%d"
set format x "%m-%d"
set key autotitle columnhead
set style line 1 linecolor rgb '#0060ad' linetype 1 linewidth 1.5 pointtype 7 pointsize 0.6
set style line 2 linecolor rgb '#5cc863' linetype 2 linewidth 1.5 pointtype 7 pointsize 0.6
set style line 3 linecolor rgb '#6000ad' linetype 3 linewidth 1.5 pointtype 7 pointsize 0.6
set xlabel 'Period (MM-DD)'                                                                    
set ytics nomirror
set ylabel  'bytes' textcolor rgb "#0060ad"
plot 'stats.csv' \
         using 1:2  with linespoints linestyle 1,  \
     ''  using 1:3  with linespoints linestyle 2,  \
     ''  using 1:4  with linespoints linestyle 3  
