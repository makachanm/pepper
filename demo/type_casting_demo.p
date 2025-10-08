include `demo/type_casting_lib.p`

println[`--- Calling functions from included file ---`]
dim int_val = get_int[]
dim real_val = get_real[]
dim str_val = get_string[]

println[`int_val: ` + to_str[int_val]]
println[`real_val: ` + to_str[real_val]]
println[`str_val: ` + str_val]

println[`--- Type Casting Demo ---`]

dim real_to_int = to_int[real_val]
println[`to_int(3.14) -> ` + to_str[real_to_int]]

dim str_to_int = to_int[str_val]
println[`to_int('456') -> ` + to_str[str_to_int]]

dim int_to_real = to_real[int_val]
println[`to_real(123) -> ` + to_str[int_to_real]]

dim int_to_str = to_str[int_val]
println[`to_str(123) -> ` + int_to_str]

println[`--- Arithmetic with Casted Types ---`]
dim result = 100 + str_to_int
println[`100 + to_int('456') = ` + to_str[result]]

quit[]