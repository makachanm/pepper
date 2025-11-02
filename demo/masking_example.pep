/*
  Pepper Masking Demo
  - Demonstrates using a sprite as a mask for drawing another sprite.
*/

set_title[`Masking Demo`]

dim content_sprite = load_sprite[`demo/res/demo.png`]
dim mask_sprite = load_sprite[`demo/res/golfball.png`]

dim mask_x = 0
dim mask_y = 240
dim mask_dir = 2

loop [true] then
  set_color[50 50 50]
  clear[]

  set_mask[mask_sprite mask_x mask_y]

  draw_sprite[content_sprite 100 100]

  reset_mask[]

  render[]

  mask_x = mask_x + mask_dir
  if [mask_x > 600 or mask_x < 0] then
    mask_dir = 0 - mask_dir
  end

  sleep[16]
end
