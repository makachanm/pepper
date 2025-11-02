include[`demo/lib/gui.pep`]

dim state = [
    `slider_value`: 0.5,
    `button_clicks`: 0
]

loop [true] then
    clear[]
    set_color[200 200 200]
    draw_rect[0 0 640 480]

    dim e = wait_event[]
    gui_update[e]

    if [gui_button[1 `Click me!` ((640 / 2) - 50) ((480 / 2) - 15) 100 30]] then
        state->button_clicks = state->button_clicks + 1
        println[`Clicked!`]
    end

    gui_label[`Clicks: ` + string[state->button_clicks] ((640 / 2) - 60)  ((480 / 2) - 30)]

    render[]
    sleep[16]
end
