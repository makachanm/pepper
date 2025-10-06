
/*
  Pepper String Functions Demo
*/

dim original_str = `  Hello, Pepper World!  `
println[`Original String: '` + original_str + `'`]

dim len = str_len[original_str]
println[`Length: ` + len]

dim trimmed_str = str_trim[original_str ` `]
println[`Trimmed: '` + trimmed_str + `'`]

dim upper_str = str_to_upper[trimmed_str]
println[`Uppercase: ` + upper_str]

dim lower_str = str_to_lower[upper_str]
println[`Lowercase: ` + lower_str]

dim sub_str = str_sub[trimmed_str 7 13]
println[`Substring (7-13): '` + sub_str + `'`]

dim = replace_length = str_len[trimmed_str]
dim replaced_str = str_replace[trimmed_str `Pepper` `Awesome` replace_length]
println[`Replaced 'Pepper' with 'Awesome': ` + replaced_str]

dim contains = str_contains[trimmed_str `Pepper`]
print[`Contains 'Pepper': `]
println[contains]

dim prefix = str_has_prefix[trimmed_str `Hello`]
print[`Starts with 'Hello': `]
println[prefix]

dim suffix = str_has_suffix[trimmed_str `World!`]
print[`Ends with 'World!': `] 
println[suffix]

dim index = str_index_of[trimmed_str `World`]
println[`Index of 'World': ` + index]

dim parts = str_split[`apple,banana,orange` `,`]
println[parts]
println[`Splitting 'apple,banana,orange' by ','`]
print[`Part 1: `]
println[parts|0|]
print[`Part 2: `]
println[parts|1|]
print[`Part 3: `]
println[parts|2|]


dim p = [ 0: `one`, 1: `two`, 2: `three` ]
dim joined_str = str_join[p ` - `]
println[`Joining a pack with ' - ': ` + joined_str]
