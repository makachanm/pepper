/* Pepper Language Example */

/* --- Pack (Key-Value Map) --- */
dim my_pack = [
    1: `first`,
    2: `second`
]

my_pack->foo = `eggs`
println[my_pack]

/* --- Loop --- */
dim counter = 0
loop [counter < 10] then
    counter = counter + 1
    if [counter == 2] then
        continue /* Skip iteration 2 */
    end
    println[counter]
end

func multiply[a b] then
    return a * b
end

dim result  = multiply[10 10]
println[result]

draw_rect[400 100 100 100]
draw_circle[300 220 100]
screen_save[`./test.png`]
finish[]



