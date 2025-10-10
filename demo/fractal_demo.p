/*
  Mandelbrot Set Fractal Demo
  This demo renders the Mandelbrot set, corrected and completed.
*/

dim width = gfx_get_width[]
dim height = gfx_get_height[]
dim x_center = -0.5
dim y_center = 0
dim x_range = 3.4
dim y_range = x_range * (to_real[height] / to_real[width])

dim min_x = x_center - (x_range / 2.0)
dim min_y = y_center - (y_range / 2.0)

dim max_iteration = 1000

println[`starting render...`]
gfx_set_source_rgb[0 0 0]
gfx_clear[]

dim i_height = 0
loop [i_height < height] then
  dim i_width = 0
  loop [i_width < width] then
    dim mw = to_real[i_width] / to_real[width]
    dim mh = to_real[i_height] / to_real[height]

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
      gfx_set_source_rgb[r g b]
    else then
      gfx_set_source_rgb[0 0 0]
    end

    gfx_draw_rect[i_width i_height 1 1]

    i_width = (i_width + 1)
  end
  i_height = (i_height + 1)
end

println[`finished`]
gfx_finish[]