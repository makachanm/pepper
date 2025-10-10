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

set_color[255 255 255] /* White */
draw_rect[0 0 800 600]

/* Set font and draw title */
set_font[`sans-serif`]

set_color[0 0 0] /* Black */

set_font_size[24]
draw_text[50 100 `Todo Item:`]

set_font_size[18]
draw_text[50 150 title]

/* Draw completed status indicator */
if [completed == true] then
    set_color[0 255 0] /* Green */
    draw_text[50 250 `Status: Completed`]
    draw_rect[50 300 200 50]
else
    set_color[255 0 0] /* Red */
    draw_text[50 250 `Status: Not Completed`]
    draw_rect[50 300 200 50]
end

/* Save the result */

/* Finish execution */
render[]