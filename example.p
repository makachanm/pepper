dim x = 10
dim y = 20

func add[a, b] then
    repeat 3 then
        a = a + 1
    end

    return a + b
end

if [x < y] then
    dim z = add[x, y]
elif [x > y] then
    dim z = add[x, y]
else then
    dim z = 0
end