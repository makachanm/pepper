/*
  Pepper Sprite Demo
  - A simple animation of a rotating and scaling sprite.
*/

gfx_set_window_title[`Sprite Demo`]

dim sprite = gfx_load_sprite[`demo/demo.png`]

dim angle = 0
dim scale = 0.7
dim scale_direction = 0.05

loop [true] then
  gfx_set_source_rgb[200 200 200]
  gfx_clear[]

  gfx_set_sprite_rotation[sprite angle]
  gfx_set_sprite_scale[sprite scale scale]

  gfx_draw_sprite[sprite 320 240]

  gfx_finish[]

  angle = angle + 2
  if [angle > 360] then
    angle = 0
  end

  scale = scale + scale_direction
  if [scale > 2.0 or scale < 0.5] then
    scale_direction = 0 - scale_direction
  end

  sleep[16] 
end
