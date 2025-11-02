### 표준 함수 (STANDARD FUNCTIONS)
Pepper는 기본적인 몇가지 표준 함수를 제공합니다. 표준 함수는 일반적인 함수와 동일하게 호출할 수 있습니다.

### 시스템 (System)

#### quit
`quit` 함수는 프로그램을 종료합니다.
```
quit[]
```

### 타입 변환 (Type Casting)

#### int
`int` 함수는 인자를 정수(integer)로 변환합니다.
```
dim i = int[3.14] /* 3 */
```

#### real
`real` 함수는 인자를 실수(real)로 변환합니다.
```
dim r = real[3] /* 3.0 */
```

#### string
`string` 함수는 인자를 문자열(string)로 변환합니다.
```
dim s = string[123] /* "123" */
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

#### atan2
`atan2` 함수는 y와 x 두 개의 숫자를 인자로 받아 y/x의 아크탄젠트 값을 반환합니다.
```
dim result = atan2[10 5]
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

#### deg2rad
`deg2rad` 함수는 각도(degree)를 라디안(radian)으로 변환합니다.
```
dim radians = deg2rad[180] /* 3.14159... */
```

#### rand_int
`rand_int` 함수는 두 개의 정수 `min`, `max`를 인자로 받아 `min`과 `max` 사이의 임의의 정수를 반환합니다.
```
dim random_number = rand_int[1 100]
```

#### rand_real
`rand_real` 함수는 두 개의 실수 `min`, `max`를 인자로 받아 `min`과 `max` 사이의 임의의 실수를 반환합니다.
```
dim random_number = rand_real[0.0 1.0]
```

### 문자열 (String)

#### strlen
`strlen` 함수는 문자열을 인자로 받아 길이를 반환합니다.
```
dim length = strlen[`hello`]
```

#### substr
`substr` 함수는 문자열, 시작 위치, 끝 위치를 인자로 받아 부분 문자열을 반환합니다.
```
dim sub = substr[`hello world` 0 5]
```

#### replace
`replace` 함수는 원본 문자열, 바꿀 문자열, 새 문자열을 인자로 받아 바뀐 문자열을 반환합니다.
```
dim replaced = replace[`hello world` `world` `pepper`]
```

#### split
`split` 함수는 문자열과 구분자를 인자로 받아 문자열을 구분자로 나눠 팩(pack)으로 반환합니다.
```
dim parts = split[`hello,world` `,`]
```

#### join
`join` 함수는 팩(pack)과 구분자를 인자로 받아 팩의 요소들을 구분자로 이어붙여 문자열로 반환합니다.
```
dim p = [ 1: `hello`, 2: `world` ]
dim joined = join[p `,`]
```

#### contains
`contains` 함수는 문자열과 부분 문자열을 인자로 받아 부분 문자열이 포함되어 있는지 여부를 불리언(boolean) 값으로 반환합니다.
```
dim found = contains[`hello world` `world`]
```

#### prefix
`prefix` 함수는 문자열과 접두사를 인자로 받아 문자열이 접두사로 시작하는지 여부를 불리언(boolean) 값으로 반환합니다.
```
dim starts = prefix[`hello world` `hello`]
```

#### suffix
`suffix` 함수는 문자열과 접미사를 인자로 받아 문자열이 접미사로 끝나는지 여부를 불리언(boolean) 값으로 반환합니다.
```
dim ends = suffix[`hello world` `world`]
```

#### upper
`upper` 함수는 문자열을 인자로 받아 모든 문자를 대문자로 바꾼 문자열을 반환합니다.
```
dim upper = upper[`hello world`]
```

#### lower
`lower` 함수는 문자열을 인자로 받아 모든 문자를 소문자로 바꾼 문자열을 반환합니다.
```
dim lower = lower[`HELLO WORLD`]
```

#### trim
`trim` 함수는 문자열을 인자로 받아 양쪽 공백을 제거한 문자열을 반환합니다.
```
dim trimmed = trim[`  hello world  `]
```

#### index
`index` 함수는 문자열과 부분 문자열을 인자로 받아 부분 문자열이 처음 나타나는 위치를 반환합니다. 부분 문자열이 없으면 -1을 반환합니다.
```
dim index = index[`hello world` `world`]
```

### JSON

#### json_stringify
`json_stringify` 함수는 팩(pack)을 인자로 받아 JSON 문자열로 변환하여 반환합니다.
```
dim p = [ `name`: `pepper`, `version`: 1 ]
dim json_str = json_stringify[p]
```

#### json_parse
`json_parse` 함수는 JSON 문자열을 인자로 받아 팩(pack)으로 변환하여 반환합니다.
```
dim json_str = `{ "name": "pepper", "version": 1 }`
dim p = json_parse[json_str]
```

### 시간 (Time)

#### sleep
`sleep` 함수는 밀리초(ms) 단위의 시간을 인자로 받아 해당 시간만큼 프로그램 실행을 멈춥니다.
```
sleep[1000] /* 1초 동안 대기 */
```

#### now
`now` 함수는 현재 시간을 유닉스 타임스탬프 (초) 로 반환합니다.
```
dim current_time = now[]
```

#### year
`year` 함수는 현재 년도를 반환합니다.
```
dim current_year = year[]
```

#### month
`month` 함수는 현재 월을 반환합니다.
```
dim current_month = month[]
```

#### day
`day` 함수는 현재 일을 반환합니다.
```
dim current_day = day[]
```

#### hour
`hour` 함수는 현재 시간을 반환합니다.
```
dim current_hour = hour[]
```

#### minute
`minute` 함수는 현재 분을 반환합니다.
```
dim current_minute = minute[]
```

#### second
`second` 함수는 현재 초를 반환합니다.
```
dim current_second = second[]
```

### 그래픽스 (Graphics)

#### clear
`clear` 함수는 그래픽 화면을 지웁니다.
```
clear[]
```

#### set_color
`set_color` 함수는 R, G, B 세 개의 정수 인자를 받아 드로잉 색상을 설정합니다. 각 값은 0과 255 사이여야 합니다.
```
set_color[255 0 0] /* Red */
```

#### set_rgba
`set_rgba` 함수는 R, G, B, A 네 개의 정수 인자를 받아 드로잉 색상과 투명도를 설정합니다. 각 값은 0과 255 사이여야 합니다.
```
set_rgba[255 0 0 128] /* 50% 투명도의 빨간색 */
```

#### draw_rect
`draw_rect` 함수는 x, y, width, height 네 개의 정수 인자를 받아 사각형을 그립니다.
```
draw_rect[10 10 100 50]
```

#### draw_dot
`draw_dot` 함수는 x, y 두 개의 정수 인자를 받아 점을 그립니다.
```
draw_dot[50 50]
```

#### draw_circle
`draw_circle` 함수는 x, y, radius 세 개의 정수 인자를 받아 원을 그립니다.
```
draw_circle[100 100 50]
```

#### draw_line
`draw_line` 함수는 x1, y1, x2, y2 네 개의 정수 인자를 받아 선을 그립니다.
```
draw_line[0 0 200 200]
```

#### draw_triangle
`draw_triangle` 함수는 x1, y1, x2, y2, x3, y3 여섯 개의 정수 인자를 받아 삼각형을 그립니다.
```
draw_triangle[100 50 50 150 150 150]
```

#### draw_bezier
`draw_bezier` 함수는 x1, y1, x2, y2, x3, y3, x4, y4 여덟 개의 정수 인자를 받아 3차 베지에 곡선을 그립니다.
```
draw_bezier[50 50 100 150 150 50 200 150]
```

#### set_font
`set_font` 함수는 폰트 이름이나 경로를 문자열 인자로 받아 텍스트 렌더링에 사용할 폰트를 지정합니다.
```
set_font[`Arial`]
```

#### set_font_size
`set_font_size` 함수는 숫자를 인자로 받아 폰트 크기를 지정합니다.
```
set_font_size[24]
```

#### draw_text
`draw_text` 함수는 x, y 좌표와 출력할 문자열을 인자로 받아 화면에 텍스트를 그립니다.
```
draw_text[10 20 `Hello Pepper!`]
```

#### save_canvas
`save_canvas` 함수는 파일 이름을 문자열 인자로 받아 현재 그래픽 화면을 PNG 파일로 저장합니다.
```
save_canvas[`output.png`]
```

#### render
`render` 함수는 그래픽 처리를 완료하고, 필요한 경우 창을 닫습니다.
```
render[]
```

#### wait_event
`wait_event` 함수는 다음 그래픽스 이벤트가 발생할 때까지 기다렸다가 해당 이벤트를 팩으로 반환합니다.
```
dim event = wait_event[]
```

#### set_line_width
`set_line_width` 함수는 선의 두께를 설정합니다.
```
set_line_width[2.0]
```

#### stroke
`stroke` 함수는 현재 경로를 따라 선을 그립니다.
```
stroke[]
```

#### fill
`fill` 함수는 현재 경로의 내부를 채웁니다.
```
fill[]
```

#### path_rect
`path_rect` 함수는 경로에 사각형을 추가합니다.
```
path_rect[10 10 100 50]
```

#### path_circle
`path_circle` 함수는 경로에 원을 추가합니다.
```
path_circle[100 100 50]
```

#### move_to
`move_to` 함수는 새로운 하위 경로를 시작합니다.
```
move_to[10 10]
```

#### line_to
`line_to` 함수는 현재 지점에서 지정된 지점까지 선을 추가합니다.
```
line_to[100 100]
```

#### close_path
`close_path` 함수는 현재 하위 경로의 시작점과 끝점을 연결하여 경로를 닫습니다.
```
close_path[]
```

#### resize_window
`resize_window` 함수는 창의 크기를 조절합니다.
```
resize_window[800 600]
```

#### get_width
`get_width` 함수는 창의 너비를 반환합니다.
```
dim width = get_width[]
```

#### get_height
`get_height` 함수는 창의 높이를 반환합니다.
```
dim height = get_height[]
```

#### set_title
`set_title` 함수는 창의 제목을 설정합니다.
```
set_title[`My Graphics Window`]
```

#### 그래픽스: 스프라이트 (Graphics: Sprite)

#### load_sprite
`load_sprite` 함수는 이미지 파일 경로를 인자로 받아 스프라이트를 로드하고 ID를 반환합니다.
```
dim sprite_id = load_sprite[`player.png`]
```

#### create_sprite
`create_sprite` 함수는 너비와 높이를 인자로 받아 빈 스프라이트를 생성하고 ID를 반환합니다.
```
dim sprite_id = create_sprite[64 64]
```

#### destroy_sprite
`destroy_sprite` 함수는 스프라이트 ID를 인자로 받아 해당 스프라이트를 메모리에서 해제합니다.
```
destroy_sprite[sprite_id]
```

#### draw_sprite
`draw_sprite` 함수는 스프라이트 ID와 x, y 좌표를 인자로 받아 화면에 스프라이트를 그립니다.
```
draw_sprite[sprite_id 100 150]
```

#### set_sprite_rotation
`set_sprite_rotation` 함수는 스프라이트 ID와 회전 각도(라디안)를 인자로 받아 스프라이트의 회전을 설정합니다.
```
set_sprite_rotation[sprite_id 0.5]
```

#### set_sprite_scale
`set_sprite_scale` 함수는 스프라이트 ID와 x, y 스케일 값을 인자로 받아 스프라이트의 크기를 조절합니다.
```
set_sprite_scale[sprite_id 2.0 2.0]
```

### 그래픽스: 마스킹 (Graphics: Masking)

#### set_mask
`set_mask` 함수는 마스크로 사용할 스프라이트 ID와 마스크를 적용할 x, y 좌표를 인자로 받습니다. 이 함수 호출 이후부터 `reset_mask`가 호출되기 전까지 그려지는 모든 내용은 해당 스프라이트의 알파 채널에 의해 마스킹됩니다.
```
set_mask[mask_sprite_id 0 0]
```

#### reset_mask
`reset_mask` 함수는 이전에 `set_mask`로 설정된 마스크를 해제합니다.
```
reset_mask[]
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