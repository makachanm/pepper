/* Pepper Eye Exmaple */

func draw_eye [eye_cx eye_cy eye_r pupil_r mouse_x mouse_y] then
    gfx_set_source_rgb[255 255 255]
    gfx_draw_circle[eye_cx eye_cy eye_r]

    dim dx = mouse_x - eye_cx
    dim dy = mouse_y - eye_cy
    dim distance = sqrt[dx*dx + dy*dy]
    
    dim max_pupil_offset = eye_r - pupil_r

    if [distance > max_pupil_offset] then
        dx = dx / distance * max_pupil_offset
        dy = dy / distance * max_pupil_offset
    end

    dim pupil_x = eye_cx + dx
    dim pupil_y = eye_cy + dy

    gfx_set_source_rgb[0 0 0]
    gfx_draw_circle[pupil_x pupil_y pupil_r]
  
end

gfx_set_window_title[`Eyes Demo`]

loop [true] then
    dim event = gfx_wait_event[]
    dim event_type = event->type

    if [event_type == `quit`] then
        break
    end

    if [event_type == `mouse_motion`] then
        gfx_clear[]
        draw_eye[420 240 100 30 event->x event->y]
        draw_eye[220 240 100 30 event->x event->y]
        gfx_finish[]
    end
end