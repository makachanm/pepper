/*
 * Simple Immediate-Mode GUI Framework for Pepper
 */

dim gui_state = [
    `hot`: 0,
    `active`: 0,
    `mouse_x`: 0,
    `mouse_y`: 0,
    `mouse_down`: false
]

func gui_update [event] then
    if [event == nil] then
        return nil    
    end

    gui_state->hot = 0
    gui_state->mouse_x = event->x
    gui_state->mouse_y = event->y

    if [event->type == `mouse_button_down`] then
            println[`Mouse down: `]
            gui_state->mouse_down = true
    end
    if [event->type == `mouse_button_up`] then
            println[`Mouse up: `]
            gui_state->mouse_down = false
    end
end

func is_hot [id x y width height] then
    if [gui_state->mouse_x >= x and gui_state->mouse_x <= x + width and
       gui_state->mouse_y >= y and gui_state->mouse_y <= y + height] then
        gui_state->hot = id
        println[`Hot: `]
        return true
    end
    return false
end

func gui_button [id text x y width height] then
    dim hot = is_hot[id x y width height]
    dim clicked = false

    if [hot and gui_state->mouse_down] then
        gui_state->active = id
    end

    path_rect[x y width height]
    if [gui_state->active == id] then
        set_color[0.5 0.5 0.5]
    else then
        if [hot] then
            set_color[0.8 0.8 0.8]
        else then
            set_color[0.7 0.7 0.7]
        end
    end
    fill[]

    set_color[0 0 0]
    draw_text[(x + 5) (y + height / 2) text]

    if [hot and not gui_state->mouse_down and gui_state->active == id] then
        clicked = true
    end

    if [not gui_state->mouse_down] then
        gui_state->active = 0
    end

    return clicked
end

func gui_label [text x y] then
    set_color[0 0 0]
    draw_text[x y text]
end

func gui_slider [id value x y width height] then
    dim hot = is_hot[id x y width height]
    dim new_value = value

    if [hot and gui_state->mouse_down] then
        gui_state->active = id
    end

    if [gui_state->active == id] then
        new_value = (gui_state->mouse_x - x) / width
        if [new_value < 0] then new_value = 0 end
        if [new_value > 1] then new_value = 1 end
    end

    path_rect[x y width height]
    set_color[0.7 0.7 0.7]
    fill[]

    path_rect[x + new_value * width - 5 y 10 height]
    if [gui_state->active == id] then
        set_color[0.5 0.5 0.5]
    else then
        if [hot] then
            set_color[0.8 0.8 0.8]
        else then
            set_color[0.7 0.7 0.7]
        end
    end
    fill[]
    
    if [not gui_state->mouse_down] then
        gui_state->active = 0
    end

    return new_value
end
