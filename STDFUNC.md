### 표준 함수 (STANDARD FUNCTIONS)
Pepper는 기본적인 몇가지 표준 함수를 제공합니다. 표준 함수는 일반적인 함수와 동일하게 호출할 수 있습니다.

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