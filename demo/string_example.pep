
/*
  Pepper String Functions Demo
*/

dim original_str = `  Hello, Pepper World!  `
println[`Original String: '` + original_str + `'`]

dim len = strlen[original_str]
println[`Length: ` + len]

dim trimmed_str = trim[original_str ` `]
println[`Trimmed: '` + trimmed_str + `'`]

dim upper_str = upper[trimmed_str]
println[`Uppercase: ` + upper_str]

dim lower_str = lower[upper_str]
println[`Lowercase: ` + lower_str]

dim sub_str = substr[trimmed_str 7 13]
println[`Substring (7-13): '` + sub_str + `'`]

dim replaced_str = replace[trimmed_str `Pepper` `Awesome`]
println[`Replaced 'Pepper' with 'Awesome': ` + replaced_str]

dim contains = contains[trimmed_str `Pepper`]
print[`Contains 'Pepper': `]
println[contains]

dim prefix = prefix[trimmed_str `Hello`]
print[`Starts with 'Hello': `]
println[prefix]

dim suffix = suffix[trimmed_str `World!`]
print[`Ends with 'World!': `] 
println[suffix]

dim index = index[trimmed_str `World`]
println[`Index of 'World': ` + index]

dim parts = split[`apple,banana,orange` `,`]
println[parts]
println[`Splitting 'apple,banana,orange' by ','`]
print[`Part 1: `]
println[parts->0]
print[`Part 2: `]
println[parts->1]
print[`Part 3: `]
println[parts->2]


dim p = [ 0: `one`, 1: `two`, 2: `three` ]
dim joined_str = join[p ` - `]
println[`Joining a pack with ' - ': ` + joined_str]

quit[]
