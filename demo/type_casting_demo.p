include `demo/type_casting_lib.p`

println[`--- Calling functions from included file ---`]
dim int_val = get_int[]
dim real_val = get_real[]
dim str_val = get_string[]

println[`int_val: ` + string[int_val]]
println[`real_val: ` + string[real_val]]
println[`str_val: ` + str_val]

println[`--- Type Casting Demo ---`]

dim real_to_int = int[real_val]
println[`int(3.14) -> ` + string[real_to_int]]

dim str_to_int = int[str_val]
println[`int('456') -> ` + string[str_to_int]]

dim int_to_real = real[int_val]
println[`real(123) -> ` + string[int_to_real]]

dim int_to_str = string[int_val]
println[`string(123) -> ` + int_to_str]

println[`--- Arithmetic with Casted Types ---`]
dim result = 100 + str_to_int
println[`100 + int('456') = ` + string[result]]

quit[]