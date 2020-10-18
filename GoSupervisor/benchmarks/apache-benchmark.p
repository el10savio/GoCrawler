# output as png image
set terminal png

# save file to "benchmarks/benchmark.png"
set output "benchmarks/benchmark.png"

# graph title
set title "/crawler/parse"

#nicer aspect ratio for image size
set size 1,0.8

# y-axis grid
set grid y

#x-axis label
set xlabel "request"

#y-axis label
set ylabel "response time (ms)"

#plot data from "out.data" using column 9 with smooth sbezier lines
plot "benchmarks/out.data" using 9 smooth sbezier with lines title "parse"