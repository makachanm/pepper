/*
  Pepper Sprite Demo
  - A simple animation of a rotating and scaling sprite.
*/

set_title[`Sprite Demo`]

dim sprite = load_sprite[`demo/res/demo.png`]

dim angle = 0
dim scale = 0.7
dim scale_direction = 0.05

loop [true] then
  set_color[200 200 200]
  clear[]

  set_sprite_rotation[sprite angle]
  set_sprite_scale[sprite scale scale]

  draw_sprite[sprite 320 240]

  render[]

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
