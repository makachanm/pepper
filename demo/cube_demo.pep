dim vertices = []
vertices|1| = [ `x`: -1, `y`: -1, `z`: -1 ]
vertices|2| = [ `x`:  1, `y`: -1, `z`: -1 ]
vertices|3| = [ `x`:  1, `y`:  1, `z`: -1 ]
vertices|4| = [ `x`: -1, `y`:  1, `z`: -1 ]
vertices|5| = [ `x`: -1, `y`: -1, `z`:  1 ]
vertices|6| = [ `x`:  1, `y`: -1, `z`:  1 ]
vertices|7| = [ `x`:  1, `y`:  1, `z`:  1 ]
vertices|8| = [ `x`: -1, `y`:  1, `z`:  1 ]

dim edges = []
edges|1|  = [ 1: 1, 2: 2 ]
edges|2|  = [ 1: 2, 2: 3 ]
edges|3|  = [ 1: 3, 2: 4 ]
edges|4|  = [ 1: 4, 2: 1 ]
edges|5|  = [ 1: 5, 2: 6 ]
edges|6|  = [ 1: 6, 2: 7 ]
edges|7|  = [ 1: 7, 2: 8 ]
edges|8|  = [ 1: 8, 2: 5 ]
edges|9|  = [ 1: 1, 2: 5 ]
edges|10| = [ 1: 2, 2: 6 ]
edges|11| = [ 1: 3, 2: 7 ]
edges|12| = [ 1: 4, 2: 8 ]

dim screen_width = 800
dim screen_height = 600

dim angle_x = 0.0
dim angle_y = 0.0

func rotate_point [point sin_x cos_x sin_y cos_y] then
    dim x_y = point->x * cos_y + point->z * sin_y
    dim z_y = -point->x * sin_y + point->z * cos_y
    
    dim y_yx = point->y * cos_x - z_y * sin_x
    dim z_yx = point->y * sin_x + z_y * cos_x

    return [`x`: x_y, `y`: y_yx, `z`: z_yx]
end

func project [point] then
    dim distance = 5.0
    dim z_factor = distance / (distance - point->z)
    dim x = point->x * z_factor * (screen_width / 4) + (screen_width / 2)
    dim y = point->y * z_factor * (screen_height / 4) + (screen_height / 2)
    return [ `x`: x, `y`: y ]
end

set_title[`3D Cube Demo`]
resize_window[screen_width screen_height]

dim num_edges = len[edges]

loop [true] then
    angle_x = angle_x + 0.01
    angle_y = angle_y + 0.015

    dim sin_x = sin[angle_x]
    dim cos_x = cos[angle_x]
    dim sin_y = sin[angle_y]
    dim cos_y = cos[angle_y]

    clear[]
    set_color[255 255 255]
    set_line_width[2.0]

    dim i = 1
    loop [i <= num_edges] then
        dim edge = edges|i|
        dim v1_idx = edge|1|
        dim v2_idx = edge|2|

        dim v1 = vertices|v1_idx|
        dim v2 = vertices|v2_idx|

        dim rotated1 = rotate_point[v1 sin_x cos_x sin_y cos_y]
        dim p1 = project[rotated1]

        dim rotated2 = rotate_point[v2 sin_x cos_x sin_y cos_y]
        dim p2 = project[rotated2]

        if [p1 != nil and p2 != nil] then
            draw_line[p1->x p1->y p2->x p2->y]
        end
        i = i + 1
    end

    render[]

    sleep[16]
end