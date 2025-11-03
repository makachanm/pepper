dim start = now[]

println[`1. Fibonacci Sequence (Recursive)`]

func fib[n] then
    if [n < 2] then
        return n
    else then
        return fib[n - 2] + fib[n - 1]
    end
end

fib[5000]

println[`2. Looping and Arithmetic`]

dim total = 0
dim i = 0
loop [i < 100000000] then
    total = total + i
    i = i + 1
end

println[`3. Pack Manipulation`]

dim p = []
i = 0
loop [i < 10000] then
    p|i| = i * 2
    i = i + 1
end

dim sum = 0
i = 0
loop [i < 10000] then
    sum = sum + p|i|
    i = i + 1
end

println[`4. Random Square Drawing`]

resize_window[400 400]
set_title[`Benchmark`]

i = 0
loop [i < 10000] then
    set_color[rand_int[0 255] rand_int[0 255] rand_int[0 255]]
    draw_rect[rand_int[0 400] rand_int[0 400] rand_int[20 80] rand_int[20 80]]
    i = i + 1

    render[]
end

dim end = now[]

dim elapsed = (end - start)

print[`Benchmark Ended. Elapsed time=`]
println[elapsed]

