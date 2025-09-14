/* Pepper Language Example - Comprehensive Syntax Showcase */

/* --- Variable Declaration & Basic Types --- */
dim an_integer = 42
dim a_real = 3.14
dim a_string = `hello pepper`
dim a_boolean = true
dim a_nil_value = nil

/* --- Pack (Key-Value Map) --- */
dim my_pack = [
    `name`: `pepper`,
    `version`: 1.0,
    1: `first`,
    2: `second`
]

/* Access and assign to pack elements */
dim original_name = my_pack|`name`|
my_pack|`name`| = `peppermint`
dim new_name = my_pack|`name`|

/* --- Operators Showcase --- */
dim a = 10
dim b = 20
dim c = 10

/* Arithmetic */
dim sum = a + b
dim difference = b - a
dim product = a * 2
dim quotient = b / a
dim remainder = 21 % a /* 21 % 10 = 1 */

/* Comparison and Logical */
dim are_equal = (a == c) and true
dim are_not_equal = (a != b) or false
dim is_greater = (b > a)
dim is_less_or_equal = (a <= c)

if [(not a_boolean) or (are_equal and is_greater)] then
    /* This block should not execute */
    dim result = `if branch`
else then
    dim result = `else branch`
end

/* --- Loop Showcase --- */
dim counter = 0
loop [counter < 5] then
    counter = counter + 1
    if [counter == 2] then
        continue /* Skip iteration 2 */
    end
    if [counter == 4] then
        break /* Exit loop at 4 */
    end
end
/* counter will be 4 */

/* --- Repeat Loop --- */
repeat 2 then
    dim dummy = 0
end


/* --- Function Definition & Calls --- */

/* Standard function */
func multiply[x y] then
    return x * y
end

dim product_result = multiply[a c] /* Call with space-separated arguments */


/* --- Final Output --- */
/* The final expression's value is implicitly returned by the script. */
/* Let's create a final pack to hold interesting results. */
dim final_results = [
    `new_name`: new_name,
    `product`: product_result,
    `loop_counter`: counter
]

print[product_result]