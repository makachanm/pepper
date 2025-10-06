### 이벤트 (EVENTS)
Pepper는 GUI 상호작용을 위해 여러 이벤트를 제공합니다. 이벤트는 `gfx_poll_event[]` 또는 `gfx_wait_event[]` 함수를 통해 얻을 수 있으며, 이벤트 정보는 `pack` 형태로 반환됩니다.

#### 이벤트 가져오기

##### gfx_has_event
처리할 이벤트가 큐에 있는지 확인합니다. 이벤트가 있으면 `true`, 없으면 `false`를 반환합니다.
```
if [gfx_has_event[]] then
    // 이벤트 처리
end
```

##### gfx_poll_event
이벤트 큐에서 이벤트를 하나 가져옵니다. 만약 처리할 이벤트가 없다면 `nil`을 반환합니다. 이 함수는 프로그램을 멈추지 않습니다.
```
dim event = gfx_poll_event[]
if [event != nil] then
    // 이벤트 처리
end
```

##### gfx_wait_event
이벤트 큐에 이벤트가 들어올 때까지 기다렸다가, 이벤트를 하나 가져옵니다. 이 함수는 이벤트가 발생할 때까지 프로그램의 실행을 멈춥니다.
```
dim event = gfx_wait_event[]
// 이벤트 처리
```

---

### 이벤트 종류 및 데이터
모든 이벤트 `pack`은 이벤트 종류를 나타내는 `type` 키를 가지고 있습니다.

#### quit
- **타입**: `quit`
- **설명**: 사용자가 창을 닫으려고 할 때 발생합니다.
- **데이터**: 없음

#### mouse_motion
- **타입**: `mouse_motion`
- **설명**: 마우스 커서가 창 안에서 움직일 때 발생합니다.
- **데이터**:
    - `x` (int): 마우스 커서의 x 좌표
    - `y` (int): 마우스 커서의 y 좌표

#### mouse_button_down
- **타입**: `mouse_button_down`
- **설명**: 마우스 버튼을 눌렀을 때 발생합니다.
- **데이터**:
    - `x` (int): 마우스 커서의 x 좌표
    - `y` (int): 마우스 커서의 y 좌표
    - `button` (int): 눌린 버튼 번호 (1: 왼쪽, 2: 가운데, 3: 오른쪽)

#### mouse_button_up
- **타입**: `mouse_button_up`
- **설명**: 마우스 버튼에서 손을 뗐을 때 발생합니다.
- **데이터**:
    - `x` (int): 마우스 커서의 x 좌표
    - `y` (int): 마우스 커서의 y 좌표
    - `button` (int): 떼어진 버튼 번호 (1: 왼쪽, 2: 가운데, 3: 오른쪽)

#### key_down
- **타입**: `key_down`
- **설명**: 키보드의 키를 눌렀을 때 발생합니다.
- **데이터**:
    - `key_name` (str): 눌린 키의 이름 (예: "A", "Return", "Space")

#### key_up
- **타입**: `key_up`
- **설명**: 키보드의 키에서 손을 뗐을 때 발생합니다.
- **데이터**:
    - `key_name` (str): 떼어진 키의 이름 (예: "A", "Return", "Space")

---

### 예제 코드
```
dim running = true
loop [running] then
    if [gfx_has_event[]] then
        dim event = gfx_poll_event[]
        dim event_type = event->type

        if [event_type == `quit`] then
            running = false
        elif [event_type == `mouse_motion`] then
            print[`Mouse moved to: `]
            print[event->x]
            print[`, `]
            println[event->y]
        elif [event_type == `key_down`] then
            println[`Key pressed: ` + event->key_name]
        end
    end

    // ... 그리기 코드 ...
end
```
