/*
  Pepper Bouncing Ball Demo
  - 공이 중력에 의해 떨어지고 바닥/벽에 튕깁니다.
*/

dim SCREEN_W = 640
dim SCREEN_H = 480

dim x = SCREEN_W / 2
dim y = SCREEN_H / 4
dim vx = 5
dim vy = 0
dim r = 20

dim gravity = 0.5
dim bounce = 1.0 /* 튕김 계수 (1.0 = 완전탄성) */
dim friction = 0.995 /* 바닥에서 수평 감쇠 */

gfx_set_window_title[`Bouncing Ball`]

loop [true] then
  /* 배경 지우기 */
  gfx_set_source_rgb[0 0 0]
  gfx_clear[]

  /* 공 그리기 */
  gfx_set_source_rgb[50 200 255]
  gfx_draw_circle[x y r]

  /* 물리 업데이트 */
  vy = vy + gravity
  x = x + vx
  y = y + vy

  /* 좌우 충돌 */
  if [ x - r < 0 ] then
    x = r
    vx = 0 - vx * bounce
  else then
    if [ x + r > SCREEN_W ] then
      x = SCREEN_W - r
      vx = 0 - vx * bounce
    end
  end

  /* 천장 충돌 */
  if [ y - r < 0 ] then
    y = r
    vy = 0 - vy * bounce
  end

  /* 바닥 충돌 */
  if [ y + r > SCREEN_H ] then
    y = SCREEN_H - r
    vy = 0 - vy * bounce
    /* 바닥에서 수평 속도 감쇠 */
    vx = vx * friction
    /* 아주 작으면 정지시켜 흔들림 방지 */
    if [ abs[vx] < 0.05 ] then
      vx = 0
    end
    if [ abs[vy] < 0.05 ] then
      vy = 0
    end
  end

  gfx_finish[]
  sleep[16] /* ~60fps */
end