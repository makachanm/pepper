/*
    Pepper Language Demo
    - Fetches a TODO item from a public JSON API.
    - Displays the information using the graphics library with custom fonts.
*/

/* Fetch TODO item from JSONPlaceholder */
dim todo = http_get_json[`https://jsonplaceholder.typicode.com/todos/1`]

/* Get title and completed status from the pack */
dim title = todo->title
dim completed = todo->completed

gfx_set_source_rgb[255 255 255] /* White */
gfx_draw_rect[0 0 800 600]

/* Set font and draw title */
gfx_set_font_face[`sans-serif`]

gfx_set_source_rgb[0 0 0] /* Black */

gfx_set_font_size[24]
gfx_draw_text[50 100 `Todo Item:`]

gfx_set_font_size[18]
gfx_draw_text[50 150 title]

/* Draw completed status indicator */
if [completed == true] then
    gfx_set_source_rgb[0 255 0] /* Green */
    gfx_draw_text[50 250 `Status: Completed`]
    gfx_draw_rect[50 300 200 50]
else
    gfx_set_source_rgb[255 0 0] /* Red */
    gfx_draw_text[50 250 `Status: Not Completed`]
    gfx_draw_rect[50 300 200 50]
end

/* Save the result */

/* Finish execution */
gfx_finish[]