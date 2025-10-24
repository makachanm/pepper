/*
  Mandelbrot Set Fractal Demo
  This demo renders the Mandelbrot set, corrected and completed.
*/

dim width = get_width[]
dim height = get_height[]
dim x_center = -0.5
dim y_center = 0
dim x_range = 3.4
dim y_range = x_range * (real[height] / real[width])

dim min_x = x_center - (x_range / 2.0)
dim min_y = y_center - (y_range / 2.0)

dim max_iteration = 1000

println[`starting render...`]
set_color[0 0 0]
clear[]

dim i_height = 0
loop [i_height < height] then
  dim i_width = 0
  loop [i_width < width] then
    dim mw = real[i_width] / real[width]
    dim mh = real[i_height] / real[height]

    dim xk = min_x + (mw * x_range)
    dim yk = min_y + (mh * y_range)

    dim real = 0.0
    dim imag = 0.0
    dim iteration = 0

    loop [((real * real) + (imag * imag)) <= 4 and (iteration < max_iteration)] then
      dim sq_x = (real * real) - (imag * imag)
      dim xtemp = sq_x + xk
      dim q = real * imag
      imag = (2 * q) + yk
      real = xtemp
      iteration = iteration + 1
    end

    if [iteration < max_iteration] then
      dim r = ((iteration * 8) % 255)
      dim g = ((iteration * 5) % 255)
      dim b = ((iteration * 12) % 255)
      set_color[r g b]
    else then
      set_color[0 0 0]
    end

    draw_dot[i_width i_height]

    i_width = (i_width + 1)
  end
  i_height = (i_height + 1)
end

println[`finished`]
render[]
save_canvas[`mandelbrot.png`]