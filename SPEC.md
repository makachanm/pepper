### TYPES
다음과 같은 타입들은 기본 타입이다.

```
42                 (int)
3.14               (real)
`foo`              (str)
true/false         (bool)
func [foo] ... end (function)
[ i: 34, j: 6.34 ] (pack)
nil                (nil)
```

### COMMENT
`/* */`로 감싸진 모든 코드는 무시된다.

### PACK
Pack은 일종의 Key-Value 맵과 같다. 모든 값은 키를 가지고 있다.
int, real, str만 키로써 사용될 수 있으며, 키는 중복될수 없고 모두 unique해야 한다. 

value는 오직 int, real, str, bool, pack만 허용되며, function과 nil은 허용되지 않는다.
```
example: simple number array in pack
[ 1: 5, 2: 7, 3: 9, 4: 11 ]

exampe: simple string K-V map in pack
[ `foo`: `bar`, `fizz`, `buzz` ]
```

### EXPRESSIONS
키워드를 제외한 모든 코드는 줄바꿈 문자를 기준으로 하나의 표현식으로써 취급된다. <br />
표현식은 몇가지 형태를 가질 수 있다. <br />
```
(name) (assign operation) (right value)
assign expression

(left value) (compare operation) (right value) 
compare expressions

(left value) (logical operation) (right value)
(logical operation) (right value)
logical expressions

(name)[(arguments) ....]
call expressions

(keyword) (right value)
keyword expression
```

각 표현식에서는 하위-표현식을 가질수 있다. 하위 표현식은 표현식 내부 특정 자리에 들어가는 하나의 표현식으로써, 상위 부모 표현식보다 실행 순서를 높게 갖게 된다.
```
(name) (assign operation) (sub expression)
assign expression

(sub expression) (compare operation) (sub expression) 
compare expressions

(sub expression) (logical operation) (sub expression)
(logical operation) (sub expression)
logical expressions

(name)[(sub expression) ....]
call expressions

(keyword) (sub expression)
keyword expression
```

### ARITHMETIC OPERATIONS
모든 산술 연산은 산술 연산자에 의해 이루어진다. 산술 연산자는 항상 우항과 좌항 사이에만 위치하며, 우항과 좌항 모두 상수값, 변수값, 함수의 호출값이 들어갈 수 있다.
```
+ (add)
- (subtract)
* (multiply)
/ (divide)
% (remainder)
```

### COMPARE OPERATIONS
모든 비교 연산은 비교 연산자에 의해 이루어진다. 비교 연산자는 항상 우항과 좌항 사이에만 위치하며, 우항과 좌항 모두 상수값, 변수값, 함수의 호출값이 들어갈 수 있다.
```
==  (equal)
!=  (not equal)
<   (less than)
>   (greater than)
<=  (less than or equal)
>=  (greater than or equal)
```

### LOGICAL OPERATIONS
모든 논리 연산은 논리 연산자에 의해 이루어진다. 논리 연산자는 하나의 값 또는 두개의 값에 대해서 수행되며, 논리 연산자의 양변은 무조건 상수값, 변수값, 함수의 호출값이여만 한다.
```
and  (and operation)
or   (or operation)
not  (not operation)
```

### ASSIGNMENT OPERATIONS
모든 대입 연산은 대입 연산자에 의해 이루어진다. 대입 연산자는 항상 우항과 좌항 사이에만 위치하며, 우항과 좌항 모두 상수값, 변수값, 함수의 호출값이 들어갈 수 있다.
```
= (insert right value to left) 
```

### CONDITIONAL BRANCH CONTROL STATEMENTS
if-elif-else 구조의 흐름은 키워드의 다음으로 주어진 표현식의 값을 통해 정해지며, 표현식의 값은 무조건 bool 값이여야만 한다. 표현식은 대괄호`[]`로 감싸야 하며, 대괄호 안에 감싸진 표현식에 따라 if와 elif의 흐름이 결정된다.
```
if [a > b] then
    (some code...)
elif [a == b] then
    (some code...)
else then
    (some code...)
end
```

### LOOP CONTROL STATEMENTS
loop 구조의 흐름은 키워드의 다음으로 주어진 표현식의 값을 통해 정해지며, 표현식의 값은 무조건 bool 값이여야 한다. 표현식은 대괄호`[]`로 감싸야 하며, 대괄호 안에 감싸진 표현식에 따라 loop의 흐름이 결정된다.
```
loop [a > b] then
    (some code...)
end
```

repeat 구조의 흐름은 키워드의 다음으로 주어진 수의 값을 통해 정해지며, 다음으로 주어진 수만큼 해당 코드를 반복하게 한다. 주어지는 수는 상수, 변수, 함수의 호출값이여야 한다.
```
repeat 3 then
    (some code...)
end
``` 

break와 continue는 loop와 repeat를 제어할 수 있으며, break는 반복문을 종료시킨 후 스코프에서 탈출시키며, continue는 다음 표현식의 실행을 중지하고 다시 반복문의 첫 표현식으로 돌아간다.

### FUNCTIONS
함수는 표현식들의 집합이다. 함수는 이름과 인자를 가질 수 있으며, 반환값을 가질 수 있다. 인자는 `[]`대괄호로 감싸야 하며, 인자들은 모두 해당 스코프 내에서만 접근할 수 있게 된다. 만일 함수의 마지막 반환값이 없다면 함수는 `nil`을 반환하게 된다.
```
func foo [fizz buzz] then
    return fizz + buzz
end
```

### VARIABLES
변수는 값을 저장하는 저장소이다. 변수는 위와 같은 타입을 지닐 수 있으며, 무조건 첫 선언 시에 초기값을 지정해 줘야 한다. 하나의 스코프 내에서 선언된 변수는 해당 스코프 외부에서 접근할 수 없다.
```
dim foo = `hello world!"
dim number = 42
dim examplepack = [ `foo`: `bar`, `fizz`, `buzz` ]
```

Pack의 특정한 Key에 접근할 때엔 특수한 문법을 사용한다.
```
dim examplepack = [ `foo`: `bar`, `fizz`, `buzz` ]
examplepack|`foo`|
```
이때, 특수 문법으로 접근한 Key는 일종의 독립적인 Variable로 취급된다.

### STANDARD FUNCTIONS
Pepper는 몇가지 표준 함수를 제공한다. 표준 함수는 일반적인 함수와 동일하게 호출할 수 있다.

#### print
`print` 함수는 하나의 인자를 받아서 표준 출력으로 출력한다.
```
print["hello world"]
```