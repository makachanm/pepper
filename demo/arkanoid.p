/*
  Simple Arkanoid-like demo for Pepper
  - 마우스로 패들을 움직이고 스페이스(또는 마우스 클릭)로 공을 발사.
  - 벽/패들/벽돌 충돌 처리, 점수와 목숨.
*/

dim SCREEN_W = 640
dim SCREEN_H = 480

/* Paddle */
dim PADDLE_W = 100
dim PADDLE_H = 16
dim paddle_x = SCREEN_W / 2 - PADDLE_W / 2
dim paddle_y = SCREEN_H - 40

/* Ball */
dim ball_x = SCREEN_W / 2
dim ball_y = paddle_y - 12
dim ball_r = 8
dim ball_vx = 4
dim ball_vy = -6
dim ball_launched = false

/* Bricks layout */
dim BRICK_COLS = 10
dim BRICK_ROWS = 5
dim BRICK_GAP = 4
dim BRICK_TOP = 40
dim BRICK_W = (SCREEN_W - (BRICK_COLS + 1) * BRICK_GAP) / BRICK_COLS
dim BRICK_H = 18

/* Brick alive map: numeric keys 1..(BRICK_ROWS*BRICK_COLS) */
dim bricks = [
  1: 1, 2: 1, 3: 1, 4: 1, 5: 1, 6: 1, 7: 1, 8: 1, 9: 1, 10: 1,
  11: 1,12: 1,13: 1,14: 1,15: 1,16: 1,17: 1,18: 1,19: 1,20: 1,
  21: 1,22: 1,23: 1,24: 1,25: 1,26: 1,27: 1,28: 1,29: 1,30: 1,
  31: 1,32: 1,33: 1,34: 1,35: 1,36: 1,37: 1,38: 1,39: 1,40: 1
]

dim score = 0
dim lives = 3

/* helper: clamp using ifs */
func clamp [v lo hi] then
  if [ v < lo ] then
    return lo
  else then
    if [ v > hi ] then
      return hi
    end
  end
  return v
end

/* main loop */
loop [true] then
  /* draw background */
  gfx_set_source_rgb[0 0 0]
  gfx_clear[]

  /* draw bricks */
  dim row = 0
  loop [row < BRICK_ROWS] then
    dim col = 0
    loop [col < BRICK_COLS] then
      dim idx = row * BRICK_COLS + col + 1
      if [ bricks|idx| == 1 ] then
        dim bx = BRICK_GAP + col * (BRICK_W + BRICK_GAP)
        dim by = BRICK_TOP + row * (BRICK_H + BRICK_GAP)
        /* color by row */
        if [ row == 0 ] then gfx_set_source_rgb[200 80 80] end
        if [ row == 1 ] then gfx_set_source_rgb[220 140 60] end
        if [ row == 2 ] then gfx_set_source_rgb[200 200 80] end
        if [ row == 3 ] then gfx_set_source_rgb[120 200 120] end
        if [ row == 4 ] then gfx_set_source_rgb[120 160 220] end
        gfx_draw_rect[bx by BRICK_W BRICK_H]
      end
      col = col + 1
    end
    row = row + 1
  end

  /* draw paddle */
  gfx_set_source_rgb[0 0 255]
  gfx_draw_rect[paddle_x paddle_y PADDLE_W PADDLE_H]

  /* draw ball */
  gfx_set_source_rgb[255 220 80]
  gfx_draw_circle[ball_x ball_y ball_r]

  /* HUD */
  gfx_set_source_rgb[255 255 255]
  gfx_set_font_size[14]
  gfx_draw_text[8 18 `Score: ` + (score)]
  gfx_draw_text[520 18 `Lives: ` + (lives)]

  gfx_finish[]

  /* input / events */
  dim e = gfx_wait_event[]
  if [ e->type == `mouse_motion` ] then
    /* center paddle on mouse x */
    paddle_x = e->x - PADDLE_W / 2
    /* clamp */
    paddle_x = clamp[paddle_x 0 SCREEN_W - PADDLE_W]
    /* if ball not launched, follow paddle */
    if [ not ball_launched ] then
      ball_x = paddle_x + PADDLE_W / 2
      ball_y = paddle_y - ball_r - 2
    end
  end

  if [ e->type == `mouse_down` ] then
    /* launch if not launched */
    if [ not ball_launched ] then
      ball_launched = true
      /* slight randomization: flip vx direction if click on left/right */
      if [ e->x < paddle_x + PADDLE_W / 2 ] then
        ball_vx = -abs[ball_vx]
      else then
        ball_vx = abs[ball_vx]
      end
      ball_vy = -6
    end
  end

  if [ e->type == `key_down` ] then
    /* space to launch/reset */
      if [ not ball_launched ] then
        ball_launched = true
        ball_vx = 4
        ball_vy = -6
      end
  end

  if [ e->type == `quit` ] then
    break
  end

  /* physics update (fixed small step) */
  if [ ball_launched ] then
    /* move */
    ball_x = ball_x + ball_vx
    ball_y = ball_y + ball_vy

    /* wall collisions */
    if [ ball_x - ball_r < 0 ] then
      ball_x = ball_r
      ball_vx = 0 - ball_vx
    end
    if [ ball_x + ball_r > SCREEN_W ] then
      ball_x = SCREEN_W - ball_r
      ball_vx = 0 - ball_vx
    end
    if [ ball_y - ball_r < 0 ] then
      ball_y = ball_r
      ball_vy = 0 - ball_vy
    end

    /* paddle collision (AABB approx) */
    if [ ball_y + ball_r >= paddle_y and ball_y - ball_r <= paddle_y + PADDLE_H and ball_x + ball_r >= paddle_x and ball_x - ball_r <= paddle_x + PADDLE_W ] then
      /* place ball above paddle */
      ball_y = paddle_y - ball_r - 1
      /* reflect Y */
      ball_vy = 0 - abs[ball_vy]
      /* change X based on hit position */
      dim hit_pos = (ball_x - (paddle_x + PADDLE_W / 2)) / (PADDLE_W / 2) /* -1..1 */
      ball_vx = ball_vx + hit_pos * 3
      /* clamp vx so it's not zero */
      if [ abs[ball_vx] < 1 ] then
        if [ ball_vx < 0 ] then ball_vx = -1 else ball_vx = 1 end
      end
    end

    /* 속도 상한선 추가 */
    dim MAX_SPEED = 0.3
    if [ abs[ball_vx] > MAX_SPEED ] then
      ball_vx = MAX_SPEED * (ball_vx / abs[ball_vx]) /* 방향 유지 */
    end
    if [ abs[ball_vy] > MAX_SPEED ] then
      ball_vy = MAX_SPEED * (ball_vy / abs[ball_vy]) /* 방향 유지 */
    end

    /* brick collisions: check all bricks */
    dim r2 = 0
    dim c2 = 0

    loop [ r2 < BRICK_ROWS ] then
        loop [ c2 < BRICK_COLS ] then

            dim idx2 = (r2 * BRICK_COLS) + (c2 + 1)

            if [ bricks|idx2| == 1 ] then
                dim bx = (BRICK_GAP + c2) * (BRICK_W + BRICK_GAP)
                dim by = (BRICK_TOP + r2) * (BRICK_H + BRICK_GAP)
      /* nearest point on rect to ball */
                dim nearestX = ball_x
                if [ nearestX < bx ] then nearestX = bx end
                if [ nearestX > bx + BRICK_W ] then nearestX = bx + BRICK_W end
      
                dim nearestY = ball_y
                if [ nearestY < by ] then nearestY = by end
                if [ nearestY > by + BRICK_H ] then nearestY = by + BRICK_H end
                dim dx = ball_x - nearestX
                dim dy = ball_y - nearestY
                if [ dx*dx + dy*dy <= ball_r*ball_r ] then
        /* hit */
                bricks|idx2| = 0
                score = score + 10
        /* simple reflect: prefer vertical invert if impact x-dist larger */
                if [ abs[dx] > abs[dy] ] then
                    ball_vx = 0 - ball_vx
                else then
                    ball_vy = 0 - ball_vy
                end
            end
        end
        c2 = c2 + 1
        end
        r2 = r2 + 1
    end

    /* fell below screen */
    if [ ball_y - ball_r > SCREEN_H ] then
      lives = lives - 1
      ball_launched = false
      /* reset ball on paddle */
      ball_x = paddle_x + PADDLE_W / 2
      ball_y = paddle_y - ball_r - 2
      ball_vx = 4
      ball_vy = -6
      if [ lives <= 0 ] then
        /* reset game */
        lives = 3
        score = 0
        /* restore all bricks */
        dim k = 1
        loop [ k <= BRICK_ROWS * BRICK_COLS ] then
          bricks->k = 1
          k = k + 1
        end
      end
    end
  else then
    /* if not launched, keep ball on paddle */
    ball_x = paddle_x + PADDLE_W / 2
    ball_y = paddle_y - ball_r - 2
  end

  /* win check */
  dim all_dead = true
  dim kk = 1
  loop [ kk <= BRICK_ROWS * BRICK_COLS ] then
    if [ bricks|kk| == 1 ] then
      all_dead = false
    end
    kk = kk + 1
  end
  if [ all_dead ] then
    /* simple win: reset bricks and increase difficulty slightly */
    score = score + 100
    ball_vx = ball_vx * 1.05 /* 속도 증가 제한 */
    ball_vy = ball_vy * 1.05 /* 속도 증가 제한 */
    if [ abs[ball_vx] > MAX_SPEED ] then
      ball_vx = MAX_SPEED * (ball_vx / abs[ball_vx]) /* 방향 유지 */
    end
    if [ abs[ball_vy] > MAX_SPEED ] then
      ball_vy = MAX_SPEED * (ball_vy / abs[ball_vy]) /* 방향 유지 */
    end
    dim t = 1
    loop [ t <= BRICK_ROWS * BRICK_COLS ] then
      bricks->t = 1
      t = t + 1
    end
    ball_launched = false
  end

  sleep[16] /* ~60fps */
end