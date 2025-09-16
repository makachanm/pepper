### STANDARD FUNCTIONS
Pepper는 몇가지 표준 함수를 제공한다. 표준 함수는 일반적인 함수와 동일하게 호출할 수 있다.

#### print
`print` 함수는 하나의 인자를 받아서 표준 출력으로 출력한다.
```
print["hello world"]
```

#### println
`println` 함수는 하나의 인자를 받아서 표준 출력으로 출력하고 개행 문자를 추가한다.
```
println["hello world"]
```

#### screen_clear
`screen_clear` 함수는 그래픽 화면을 지운다. 인자는 필요 없다.
```
screen_clear[]
```

#### set_color
`set_source_rgb` 함수는 R, G, B 세 개의 정수 인자를 받아서 드로잉 색상을 설정한다. 각 값은 0과 255 사이여야 한다.
```
set_source_rgb[255, 0, 0] /* Red */
```

#### draw_rect
`draw_rect` 함수는 x, y, width, height 네 개의 정수 인자를 받아서 사각형을 그린다.
```
draw_rect[10, 10, 100, 50]
```

#### draw_circle
`draw_circle` 함수는 x, y, radius 세 개의 정수 인자를 받아서 원을 그린다.
```
draw_circle[100, 100, 50]
```

#### draw_line
`draw_line` 함수는 x1, y1, x2, y2 네 개의 정수 인자를 받아서 선을 그린다.
```
draw_line[0, 0, 200, 200]
```

#### screen_save
`screen_save` 함수는 파일 이름을 문자열 인자로 받아서 현재 그래픽 화면을 PNG 파일로 저장한다.
```
screen_save["output.png"]
```

#### finish
`finish` 함수는 프로그램을 종료한다. 인자는 필요 없다.
```
finish[]
```
