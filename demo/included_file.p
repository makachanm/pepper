dim included_var = `Hello from included file!`

func included_func [arg] then
    return included_var + ` included_func called with: ` + arg
end
