/*
  Pepper Animation Demo
  - A simple animation of a bouncing rectangle.
*/

dim x = 0
dim y = 240
dim w = 50
dim h = 50
dim dx = 5 /* speed and direction */

set_title[`Bouncing Rectangle`]

loop [true] then
  /* Clear screen to black */
  set_color[0 0 0]
  clear[]

  /* Draw rectangle */
  set_color[255 0 0]
  draw_rect[x y w h]

  /* Update screen */
  render[]

  /* Move rectangle */
  x = x + dx

  /* Bounce off edges */
  if [x + w > 640 or x < 0] then
    dx = 0 - dx
  else then
    dx = dx + 2
  end

  /* Wait */
  sleep[16] /* ~60 fps */
end
