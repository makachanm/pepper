### 표준 함수 (STANDARD FUNCTIONS)
Pepper는 기본적인 몇가지 표준 함수를 제공합니다. 표준 함수는 일반적인 함수와 동일하게 호출할 수 있습니다.

### 시스템 (System)

#### quit
`quit` 함수는 프로그램을 종료합니다.
```
quit[]
```

### 타입 변환 (Type Casting)

#### to_int
`to_int` 함수는 인자를 정수(integer)로 변환합니다.
```
dim i = to_int[3.14] /* 3 */
```

#### to_real
`to_real` 함수는 인자를 실수(real)로 변환합니다.
```
dim r = to_real[3] /* 3.0 */
```

#### to_str
`to_str` 함수는 인자를 문자열(string)로 변환합니다.
```
dim s = to_str[123] /* "123" */
```

### 입출력 (IO)

#### print
`print` 함수는 하나의 인자를 받아 표준 출력으로 출력합니다.
```
print[`hello world`]
```

#### println
`println` 함수는 하나의 인자를 받아 표준 출력으로 출력하고 개행 문자를 추가합니다.
```
println[`hello world`]
```

#### readln
`readln` 함수는 표준 입력에서 한 줄을 읽어와 문자열로 반환합니다.
```
dim line = readln[]
```

#### read_file
`read_file` 함수는 파일 경로를 인자로 받아 파일의 내용을 문자열로 반환합니다.
```
dim content = read_file[`path/to/file.txt`]
```

#### write_file
`write_file` 함수는 파일 경로와 내용을 인자로 받아 파일에 내용을 씁니다. 성공 시 `true`를, 실패 시 `false`를 반환합니다.
```
write_file[`path/to/file.txt` `hello world`]
```

### 수학 (Math)

#### sin
`sin` 함수는 숫자를 인자로 받아 해당 값의 사인(sine) 값을 반환합니다.
```
dim result = sin[90]
```

#### cos
`cos` 함수는 숫자를 인자로 받아 해당 값의 코사인(cosine) 값을 반환합니다.
```
dim result = cos[0]
```

#### tan
`tan` 함수는 숫자를 인자로 받아 해당 값의 탄젠트(tangent) 값을 반환합니다.
```
dim result = tan[45]
```

#### asin
`asin` 함수는 숫자를 인자로 받아 해당 값의 아크사인(arc-sine) 값을 반환합니다.
```
dim result = asin[1]
```

#### acos
`acos` 함수는 숫자를 인자로 받아 해당 값의 아크코사인(arc-cosine) 값을 반환합니다.
```
dim result = acos[1]
```

#### atan
`atan` 함수는 숫자를 인자로 받아 해당 값의 아크탄젠트(arc-tangent) 값을 반환합니다.
```
dim result = atan[1]
```

#### 2atan
`2atan` 함수는 y와 x 두 개의 숫자를 인자로 받아 y/x의 아크탄젠트 값을 반환합니다.
```
dim result = 2atan[10 5]
```

#### sqrt
`sqrt` 함수는 숫자를 인자로 받아 해당 값의 제곱근을 반환합니다.
```
dim result = sqrt[16]
```

#### pow
`pow` 함수는 밑과 지수, 두 개의 숫자를 인자로 받아 거듭제곱 값을 반환합니다.
```
dim result = pow[2 8] /* 256 */
```

#### log
`log` 함수는 숫자를 인자로 받아 자연로그(natural logarithm) 값을 반환합니다.
```
dim result = log[10]
```

#### exp
`exp` 함수는 숫자를 인자로 받아 e의 거듭제곱(e^x) 값을 반환합니다.
```
dim result = exp[2]
```

#### abs
`abs` 함수는 숫자를 인자로 받아 절댓값을 반환합니다.
```
dim result = abs[-10] /* 10 */
```

#### len
`len` 함수는 팩(pack)을 인자로 받아 팩의 요소 개수(길이)를 정수로 반환합니다.
```
dim my_pack = [ `a`: 1, `b`: 2, `c`: 3 ]
dim length = len[my_pack] /* 3 */
```

### 문자열 (String)

#### str_len
`str_len` 함수는 문자열을 인자로 받아 길이를 반환합니다.
```
dim length = str_len[`hello`]
```

#### str_sub
`str_sub` 함수는 문자열, 시작 위치, 끝 위치를 인자로 받아 부분 문자열을 반환합니다.
```
dim sub = str_sub[`hello world` 0 5]
```

#### str_replace
`str_replace` 함수는 원본 문자열, 바꿀 문자열, 새 문자열을 인자로 받아 바뀐 문자열을 반환합니다.
```
dim replaced = str_replace[`hello world` `world` `pepper`]
```

#### str_split
`str_split` 함수는 문자열과 구분자를 인자로 받아 문자열을 구분자로 나눠 팩(pack)으로 반환합니다.
```
dim parts = str_split[`hello,world` `,`]
```

#### str_join
`str_join` 함수는 팩(pack)과 구분자를 인자로 받아 팩의 요소들을 구분자로 이어붙여 문자열로 반환합니다.
```
dim p = [ 1: `hello`, 2: `world` ]
dim joined = str_join[p `,`]
```

#### str_contains
`str_contains` 함수는 문자열과 부분 문자열을 인자로 받아 부분 문자열이 포함되어 있는지 여부를 불리언(boolean) 값으로 반환합니다.
```
dim found = str_contains[`hello world` `world`]
```

#### str_has_prefix
`str_has_prefix` 함수는 문자열과 접두사를 인자로 받아 문자열이 접두사로 시작하는지 여부를 불리언(boolean) 값으로 반환합니다.
```
dim starts = str_has_prefix[`hello world` `hello`]
```

#### str_has_suffix
`str_has_suffix` 함수는 문자열과 접미사를 인자로 받아 문자열이 접미사로 끝나는지 여부를 불리언(boolean) 값으로 반환합니다.
```
dim ends = str_has_suffix[`hello world` `world`]
```

#### str_to_lower
`str_to_lower` 함수는 문자열을 인자로 받아 모든 문자를 소문자로 바꾼 문자열을 반환합니다.
```
dim lower = str_to_lower[`HELLO WORLD`]
```

#### str_to_upper
`str_to_upper` 함수는 문자열을 인자로 받아 모든 문자를 대문자로 바꾼 문자열을 반환합니다.
```
dim upper = str_to_upper[`hello world`]
```

#### str_trim
`str_trim` 함수는 문자열을 인자로 받아 양쪽 공백을 제거한 문자열을 반환합니다.
```
dim trimmed = str_trim[`  hello world  `]
```

#### str_index_of
`str_index_of` 함수는 문자열과 부분 문자열을 인자로 받아 부분 문자열이 처음 나타나는 위치를 반환합니다. 부분 문자열이 없으면 -1을 반환합니다.
```
dim index = str_index_of[`hello world` `world`]
```

### JSON

#### json_encode
`json_encode` 함수는 팩(pack)을 인자로 받아 JSON 문자열로 변환하여 반환합니다.
```
dim p = [ `name`: `pepper`, `version`: 1 ]
dim json_str = json_encode[p]
```

#### json_decode
`json_decode` 함수는 JSON 문자열을 인자로 받아 팩(pack)으로 변환하여 반환합니다.
```
dim json_str = `{ "name": "pepper", "version": 1 }`
dim p = json_decode[json_str]
```

### 시간 (Time)

#### sleep
`sleep` 함수는 밀리초(ms) 단위의 시간을 인자로 받아 해당 시간만큼 프로그램 실행을 멈춥니다.
```
sleep[1000] /* 1초 동안 대기 */
```

### 그래픽스 (Graphics)

#### gfx_clear
`gfx_clear` 함수는 그래픽 화면을 지웁니다.
```
gfx_clear[]
```

#### gfx_set_source_rgb
`gfx_set_source_rgb` 함수는 R, G, B 세 개의 정수 인자를 받아 드로잉 색상을 설정합니다. 각 값은 0과 255 사이여야 합니다.
```
gfx_set_source_rgb[255 0 0] /* Red */
```

#### gfx_draw_rect
`gfx_draw_rect` 함수는 x, y, width, height 네 개의 정수 인자를 받아 사각형을 그립니다.
```
gfx_draw_rect[10 10 100 50]
```

#### gfx_draw_circle
`gfx_draw_circle` 함수는 x, y, radius 세 개의 정수 인자를 받아 원을 그립니다.
```
gfx_draw_circle[100 100 50]
```

#### gfx_draw_line
`gfx_draw_line` 함수는 x1, y1, x2, y2 네 개의 정수 인자를 받아 선을 그립니다.
```
gfx_draw_line[0 0 200 200]
```

#### gfx_draw_triangle
`gfx_draw_triangle` 함수는 x1, y1, x2, y2, x3, y3 여섯 개의 정수 인자를 받아 삼각형을 그립니다.
```
gfx_draw_triangle[100 50 50 150 150 150]
```

#### gfx_draw_bezier
`gfx_draw_bezier` 함수는 x1, y1, x2, y2, x3, y3, x4, y4 여덟 개의 정수 인자를 받아 3차 베지에 곡선을 그립니다.
```
gfx_draw_bezier[50 50 100 150 150 50 200 150]
```

#### gfx_set_font_face
`gfx_set_font_face` 함수는 폰트 이름이나 경로를 문자열 인자로 받아 텍스트 렌더링에 사용할 폰트를 지정합니다.
```
gfx_set_font_face[`Arial`]
```

#### gfx_set_font_size
`gfx_set_font_size` 함수는 숫자를 인자로 받아 폰트 크기를 지정합니다.
```
gfx_set_font_size[24]
```

#### gfx_draw_text
`gfx_draw_text` 함수는 x, y 좌표와 출력할 문자열을 인자로 받아 화면에 텍스트를 그립니다.
```
gfx_draw_text[10 20 `Hello Pepper!`]
```

#### gfx_save_to_file
`gfx_save_to_file` 함수는 파일 이름을 문자열 인자로 받아 현재 그래픽 화면을 PNG 파일로 저장합니다.
```
gfx_save_to_file[`output.png`]
```

#### gfx_finish
`gfx_finish` 함수는 그래픽 처리를 완료하고, 필요한 경우 창을 닫습니다.
```
gfx_finish[]
```

#### gfx_wait_event
`gfx_wait_event` 함수는 다음 그래픽스 이벤트가 발생할 때까지 기다렸다가 해당 이벤트를 팩으로 반환합니다.
```
dim event = gfx_wait_event[]
```

#### gfx_set_line_width
`gfx_set_line_width` 함수는 선의 두께를 설정합니다.
```
gfx_set_line_width[2.0]
```

#### gfx_stroke
`gfx_stroke` 함수는 현재 경로를 따라 선을 그립니다.
```
gfx_stroke[]
```

#### gfx_fill
`gfx_fill` 함수는 현재 경로의 내부를 채웁니다.
```
gfx_fill[]
```

#### gfx_path_rectangle
`gfx_path_rectangle` 함수는 경로에 사각형을 추가합니다.
```
gfx_path_rectangle[10 10 100 50]
```

#### gfx_path_circle
`gfx_path_circle` 함수는 경로에 원을 추가합니다.
```
gfx_path_circle[100 100 50]
```

#### gfx_path_move_to
`gfx_path_move_to` 함수는 새로운 하위 경로를 시작합니다.
```
gfx_path_move_to[10 10]
```

#### gfx_path_line_to
`gfx_path_line_to` 함수는 현재 지점에서 지정된 지점까지 선을 추가합니다.
```
gfx_path_line_to[100 100]
```

#### gfx_path_close
`gfx_path_close` 함수는 현재 하위 경로의 시작점과 끝점을 연결하여 경로를 닫습니다.
```
gfx_path_close[]
```

#### gfx_resize
`gfx_resize` 함수는 창의 크기를 조절합니다.
```
gfx_resize[800 600]
```

#### gfx_get_width
`gfx_get_width` 함수는 창의 너비를 반환합니다.
```
dim width = gfx_get_width[]
```

#### gfx_get_height
`gfx_get_height` 함수는 창의 높이를 반환합니다.
```
dim height = gfx_get_height[]
```

#### gfx_set_window_title
`gfx_set_window_title` 함수는 창의 제목을 설정합니다.
```
gfx_set_window_title[`My Graphics Window`]
```

#### 그래픽스: 스프라이트 (Graphics: Sprite)

#### gfx_load_sprite
`gfx_load_sprite` 함수는 이미지 파일 경로를 인자로 받아 스프라이트를 로드하고 ID를 반환합니다.
```
dim sprite_id = gfx_load_sprite[`player.png`]
```

#### gfx_create_sprite
`gfx_create_sprite` 함수는 너비와 높이를 인자로 받아 빈 스프라이트를 생성하고 ID를 반환합니다.
```
dim sprite_id = gfx_create_sprite[64 64]
```

#### gfx_destroy_sprite
`gfx_destroy_sprite` 함수는 스프라이트 ID를 인자로 받아 해당 스프라이트를 메모리에서 해제합니다.
```
gfx_destroy_sprite[sprite_id]
```

#### gfx_draw_sprite
`gfx_draw_sprite` 함수는 스프라이트 ID와 x, y 좌표를 인자로 받아 화면에 스프라이트를 그립니다.
```
gfx_draw_sprite[sprite_id 100 150]
```

#### gfx_set_sprite_rotation
`gfx_set_sprite_rotation` 함수는 스프라이트 ID와 회전 각도(라디안)를 인자로 받아 스프라이트의 회전을 설정합니다.
```
gfx_set_sprite_rotation[sprite_id 0.5]
```

#### gfx_set_sprite_scale
`gfx_set_sprite_scale` 함수는 스프라이트 ID와 x, y 스케일 값을 인자로 받아 스프라이트의 크기를 조절합니다.
```
gfx_set_sprite_scale[sprite_id 2.0 2.0]
```

### HTTP

#### http_get
`http_get` 함수는 URL을 인자로 받아 HTTP GET 요청을 보내고, 응답 본문을 문자열로 반환합니다.
```
dim response = http_get[`https://example.com`]
```

#### http_post
`http_post` 함수는 URL과 요청 본문을 인자로 받아 HTTP POST 요청을 보내고, 응답 본문을 문자열로 반환합니다.
```
dim response = http_post[`https://example.com/api` `{ "key": "value" }`]
```

#### http_get_json
`http_get_json` 함수는 URL을 인자로 받아 HTTP GET 요청을 보내고, JSON 응답을 팩(pack)으로 변환하여 반환합니다.
```
dim data = http_get_json[`https://api.example.com/data`]
```
