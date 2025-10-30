/*
  1-Player Ping-Pong Demo for Pepper
  - A compromise between bounce_example.pep and eyes_example.pep
  - Move the paddle with the mouse to keep the ball in play.
  - Syntax strictly follows SPEC.md
*/

dim SCREEN_W = 640
dim SCREEN_H = 480

dim paddle = [
  `w`: 120,
  `h`: 16,
  `x`: 0,
  `y`: (SCREEN_H - 30)
]
paddle->x = (SCREEN_W / 2) - (paddle->w / 2)

dim ball = [
  `r`: 10,
  `x`: 0,
  `y`: 0,
  `vx`: 5,
  `vy`: -5
]

dim game = [
  `score`: 0,
  `game_over`: 0 
]


func reset_ball[] then
  ball->x = (SCREEN_W / 2)
  ball->y = (SCREEN_H / 2)
  ball->vx = 5
  ball->vy = -5
end

func init_game[] then
  println[`Initializing Game`]
  game->score = 0
  game->game_over = 0
  reset_ball[]
end

func handle_input[e] then
  if [e == nil] then
    return nil
  end

  if [e->type == `mouse_motion` and game->game_over == 0] then
    paddle->x = e->x - (paddle->w / 2)
    if [paddle->x < 0] then
      paddle->x = 0
    end
    if [paddle->x > (SCREEN_W - paddle->w)] then
      paddle->x = (SCREEN_W - paddle->w)
    end
  end
  
  if [game->game_over == 1 and (e->type == `key_down`)] then 
    init_game[]
  end
end

func update_game[] then
  if[game->game_over == 1] then
    return nil
  end

  ball->x = ball->x + ball->vx
  ball->y = ball->y + ball->vy

  if [(ball->x - ball->r) < 0] then
    ball->vx = abs[ball->vx]
  end
  if [(ball->x + ball->r) > SCREEN_W] then
    ball->vx = 0 - abs[ball->vx]
  end
  if [(ball->y - ball->r) < 0] then
    ball->vy = abs[ball->vy]
  end

  if [(ball->y + ball->r) > paddle->y and (ball->y - ball->r) < (paddle->y + paddle->h)] then
    if [(ball->x > paddle->x) and (ball->x < (paddle->x + paddle->w))] then
      ball->vy = 0 - abs[ball->vy]
      game->score = game->score + 1
    end
  end

  if [(ball->y - ball->r) >= SCREEN_H] then
    game->game_over = 1 
    println[`Game Over`]
    println[`Final Score: ` + game->score]
  end
end

func draw_game[] then
  set_color[0 0 0]
  clear[]

  if [game->game_over == 1] then
    set_color[255 255 255]
    set_font_size[30]
    draw_text[(SCREEN_W/2 - 100) (SCREEN_H/2 - 50) `GAME OVER`]
    set_font_size[16]
    draw_text[(SCREEN_W/2 - 120) (SCREEN_H/2) (`Final Score: ` + game->score)]
    draw_text[(SCREEN_W/2 - 100) (SCREEN_H/2 + 30) `Press Any Key to Restart`]
  else then
    set_color[0 0 200]
    draw_rect[paddle->x paddle->y paddle->w paddle->h]

    set_color[200 200 0]
    draw_circle[ball->x ball->y ball->r]

    set_color[255 255 255]
    set_font_size[20]
    draw_text[10 20 (`Score: ` + game->score)]
  end
end


init_game[]

loop [true] then
  loop [true] then
    dim e = wait_event[]
    if [e == nil] then
      break
    end
    if [e->type == `quit`] then
      quit[]
    end
    handle_input[e]
  end

  update_game[]
  draw_game[]
  
  render[]
  sleep[16] 
end