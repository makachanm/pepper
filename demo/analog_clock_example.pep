/*
 * Analog Clock Example
 */

dim width = 640
dim height = 480
set_title[`Analog Clock Example`]

dim center_x = width / 2
dim center_y = height / 2
dim radius = 200

loop [true] then
    clear[]

    set_color[255 255 255]
    draw_circle[center_x center_y radius]

    dim h = hour[]
    dim m = minute[]
    dim s = second[]

    set_color[255 255 255]
    dim hour_angle = 90 - (h % 12 + m / 60) * 30
    dim hour_len = radius * 0.5
    dim hour_x = center_x + cos[deg2rad[hour_angle]] * hour_len
    dim hour_y = center_y - sin[deg2rad[hour_angle]] * hour_len
    set_color[255 0 0]
    set_line_width[4]
    draw_line[center_x center_y hour_x hour_y]

    set_color[255 255 255]
    dim minute_angle = 90 - (m + s / 60) * 6
    dim minute_len = radius * 0.8
    dim minute_x = center_x + cos[deg2rad[minute_angle]] * minute_len
    dim minute_y = center_y - sin[deg2rad[minute_angle]] * minute_len
    set_color[255 0 0]
    set_line_width[2]
    draw_line[center_x center_y minute_x minute_y]

    dim second_angle = 90 - s * 6
    dim second_len = radius * 0.9
    dim second_x = center_x + cos[deg2rad[second_angle]] * second_len
    dim second_y = center_y - sin[deg2rad[second_angle]] * second_len
    set_color[255 0 0]
    set_line_width[1]
    draw_line[center_x center_y second_x second_y]

    render[]
    sleep[100]
end

