include[`gui.pep`]

dim state = [
    `slider_value`: 0.5,
    `button_clicks`: 0
]

loop [true] then
    clear[1 1 1]

    dim e = wait_event[]
    gui_update[e]

    if [gui_button[1 `Click me!` 50 50 100 30]] then
        state->button_clicks = state->button_clicks + 1
        println[`Clicked!`]
    end

    gui_label[`Clicks: ` + string[state->button_clicks] 50 100]

    state->slider_value = gui_slider[2 state->slider_value 50 150 200 20]
    gui_label[`Slider: ` + string[state->slider_value] 50 180]

    render[]
    sleep[16]
end
