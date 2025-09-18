### STANDARD FUNCTIONS
Pepper는 몇가지 표준 함수를 제공한다. 표준 함수는 일반적인 함수와 동일하게 호출할 수 있다.

### IO

#### print
`print` 함수는 하나의 인자를 받아서 표준 출력으로 출력한다.
```
print[`hello world`]
```

#### println
`println` 함수는 하나의 인자를 받아서 표준 출력으로 출력하고 개행 문자를 추가한다.
```
println[`hello world`]
```

#### readln
`io_readln` 함수는 표준 입력에서 한 줄을 읽어와 문자열로 반환한다.
```
dim line = readln[]
```

#### read_file
`io_read_file` 함수는 파일 경로를 인자로 받아서 파일의 내용을 문자열로 반환한다.
```
dim content = read_file[`path/to/file.txt`]
```

#### write_file
`io_write_file` 함수는 파일 경로와 내용을 인자로 받아서 파일에 내용을 쓴다. 성공 시 `true`를, 실패 시 `false`를 반환한다.
```
write_file[`path/to/file.txt`, `hello world`]
```

### Math

#### sin
`sin` 함수는 숫자를 인자로 받아서 해당 값의 사인(sine) 값을 반환한다.
```
dim result = sin[90]
```

#### cos
`cos` 함수는 숫자를 인자로 받아서 해당 값의 코사인(cosine) 값을 반환한다.
```
dim result = cos[0]
```

#### tan
`tan` 함수는 숫자를 인자로 받아서 해당 값의 탄젠트(tangent) 값을 반환한다.
```
dim result = tan[45]
```

#### sqrt
`sqrt` 함수는 숫자를 인자로 받아서 해당 값의 제곱근을 반환한다.
```
dim result = sqrt[16]
```

### String

#### str_len
`str_len` 함수는 문자열을 인자로 받아서 길이를 반환한다.
```
dim length = str_len[`hello`]
```

#### str_sub
`str_sub` 함수는 문자열, 시작 위치, 끝 위치를 인자로 받아서 부분 문자열을 반환한다.
```
dim sub = str_sub[`hello world`, 0, 5]
```

#### str_replace
`str_replace` 함수는 원본 문자열, 바꿀 문자열, 새 문자열을 인자로 받아서 바뀐 문자열을 반환한다.
```
dim replaced = str_replace[`hello world`, `world`, `pepper`]
```

### JSON

#### json_encode
`json_encode` 함수는 팩(pack)을 인자로 받아서 JSON 문자열로 변환하여 반환한다.
```
dim p = [ `name`: `pepper`, `version`: 1 ]
dim json_str = json_encode[p]
```

#### json_decode
`json_decode` 함수는 JSON 문자열을 인자로 받아서 팩(pack)으로 변환하여 반환한다.
```
dim json_str = `{ \`name\`: \`pepper\`, \`version\`: 1 }`
dim p = json_decode[json_str]
```

### Graphics

#### gfx_clear
`gfx_clear` 함수는 그래픽 화면을 지운다.
```
gfx_clear[]
```

#### gfx_set_source_rgb
`gfx_set_source_rgb` 함수는 R, G, B 세 개의 정수 인자를 받아서 드로잉 색상을 설정한다. 각 값은 0과 255 사이여야 한다.
```
gfx_set_source_rgb[255, 0, 0] /* Red */
```

#### gfx_draw_rect
`gfx_draw_rect` 함수는 x, y, width, height 네 개의 정수 인자를 받아서 사각형을 그린다.
```
gfx_draw_rect[10, 10, 100, 50]
```

#### gfx_draw_circle
`gfx_draw_circle` 함수는 x, y, radius 세 개의 정수 인자를 받아서 원을 그린다.
```
gfx_draw_circle[100, 100, 50]
```

#### gfx_draw_line
`gfx_draw_line` 함수는 x1, y1, x2, y2 네 개의 정수 인자를 받아서 선을 그린다.
```
gfx_draw_line[0, 0, 200, 200]
```

#### gfx_draw_triangle
`gfx_draw_triangle` 함수는 x1, y1, x2, y2, x3, y3 여섯 개의 정수 인자를 받아서 삼각형을 그린다.
```
gfx_draw_triangle[100, 50, 50, 150, 150, 150]
```

#### gfx_draw_bezier
`gfx_draw_bezier` 함수는 x1, y1, x2, y2, x3, y3, x4, y4 여덟 개의 정수 인자를 받아서 3차 베지에 곡선을 그린다.
```
#### gfx_draw_bezier
`gfx_draw_bezier` 함수는 x1, y1, x2, y2, x3, y3, x4, y4 여덟 개의 정수 인자를 받아서 3차 베지에 곡선을 그린다.
```
gfx_draw_bezier[50 50 100 150 150 50 200 150]
```

#### gfx_set_font_face
`gfx_set_font_face` 함수는 폰트 이름이나 경로를 문자열 인자로 받아서 텍스트 렌더링에 사용할 폰트를 지정한다.
```
gfx_set_font_face[`Arial`]
```

#### gfx_set_font_size
`gfx_set_font_size` 함수는 숫자를 인자로 받아서 폰트 크기를 지정한다.
```
gfx_set_font_size[24]
```

#### gfx_draw_text
`gfx_draw_text` 함수는 x, y 좌표와 출력할 문자열을 인자로 받아서 화면에 텍스트를 그린다.
```
gfx_draw_text[10 20 `Hello Pepper!`]
```
```

#### gfx_draw_text
`gfx_draw_text` 함수는 x, y 좌표와 출력할 문자열을 인자로 받아서 화면에 텍스트를 그린다.
```
gfx_draw_text[10, 20, `Hello Pepper!`]
```

#### gfx_save_to_file
`gfx_save_to_file` 함수는 파일 이름을 문자열 인자로 받아서 현재 그래픽 화면을 PNG 파일로 저장한다.
```
gfx_save_to_file[`output.png`]
```

#### gfx_finish
`gfx_finish` 함수는 그래픽 처리를 완료하고, 필요한 경우 창을 닫는다.
```
gfx_finish[]
```

### HTTP

#### http_get
`http_get` 함수는 URL을 인자로 받아서 HTTP GET 요청을 보내고, 응답 본문을 문자열로 반환한다.
```
dim response = http_get[`https://example.com`]
```

#### http_post
`http_post` 함수는 URL과 요청 본문을 인자로 받아서 HTTP POST 요청을 보내고, 응답 본문을 문자열로 반환한다.
```
dim response = http_post[`https://example.com/api`, `{ \`key\`: \`value\` }`]
```

#### http_get_json
`http_get_json` 함수는 URL을 인자로 받아서 HTTP GET 요청을 보내고, JSON 응답을 팩(pack)으로 변환하여 반환한다.
```
dim data = http_get_json[`https://api.example.com/data`]
```
