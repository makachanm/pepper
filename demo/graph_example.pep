/*
    Graphing Library Example
    This example demonstrates how to use the graph.pep library to draw a simple line graph.
*/
include[`demo/lib/graph.pep`]

dim width = 800
dim height = 600

resize_window[width height]
set_title[`Graph Example`]

dim data = [ 0: 0 ]
dim i = 0
loop [i < 100] then
    data|i| = (sin[i * 0.1] + 1) * 100
    i = i + 1
end

set_color[32 32 32]
clear[]

set_color[255 255 255]

drawAxes[width height `Time` `Value`]

plotLine[data width height]

render[]
