/* graph.pep - A simple graphing library for Pepper */

/*
drawAxes draws the X and Y axes for the graph.
width, height: The dimensions of the canvas.
xLabel, yLabel: The labels for the axes.
*/
func drawAxes [width height xLabel yLabel] then
    /* Draw Y axis */
    draw_line[50 50 50 height - 50]
    /* Draw X axis */
    draw_line[50 height - 50 width - 50 height - 50]

    /* Draw labels */
    draw_text[width - 100 height - 30 xLabel]
    draw_text[60 60 yLabel]
end

/*
plotLine draws a simple line graph from a pack of data points.
data: A pack of numbers to plot, with integer keys starting from 0.
width, height: The dimensions of the canvas.
*/
func plotLine [data width height] then
    dim numPoints = len[data]
    if [numPoints < 2] then
        return nil
    end

    dim graphWidth = width - 100
    dim graphHeight = height - 100
    dim xStep = graphWidth / (numPoints - 1)

    /* Find max value in data to scale the graph */
    dim maxVal = 0
    dim i = 0
    loop [i < numPoints] then
        if [data|i| > maxVal] then
            maxVal = data|i|
        end
        i = i + 1
    end

    /* Draw lines between data points */
    i = 0
    loop [i < numPoints - 1] then
        dim x1 = 50 + (i * xStep)
        dim val1 = data|i|
        dim y1 = (height - 50) - (val1 / maxVal * graphHeight)
        
        dim next_i = i + 1
        dim x2 = 50 + (next_i * xStep)
        dim val2 = data|next_i|
        dim y2 = (height - 50) - (val2 / maxVal * graphHeight)

        draw_line[x1 y1 x2 y2]
        i = i + 1
    end
end