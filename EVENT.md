# Pepper Event Handling

Pepper supports an event-driven programming model, which is useful for creating interactive applications and games. Events from the mouse and keyboard are placed in an event queue, which can be accessed from your Pepper script.

### `wait_event[]`

Waits for the next event. If the event queue is empty, this function will pause the script's execution until an event occurs.

**Returns:**
- An event pack.

## Event Packs

The event functions return packs that contain information about the event. All event packs have a `type` field, which is a string that identifies the type of the event.

### Quit Event

-   `type`: `"quit"`
-   Triggered when the user closes the graphics window.

Example:
```
{
  type: "quit"
}
```

### Mouse Motion Event

-   `type`: `"mouse_motion"`
-   `x`: The new x-coordinate of the mouse cursor.
-   `y`: The new y-coordinate of the mouse cursor.

Example:
```
{
  type: "mouse_motion",
  x: 320,
  y: 240
}
```

### Mouse Button Events

-   `type`: `"mouse_button_down"` or `"mouse_button_up"`
-   `x`: The x-coordinate of the mouse cursor when the button was pressed/released.
-   `y`: The y-coordinate of the mouse cursor when the button was pressed/released.
-   `button`: The mouse button that was pressed/released.
    -   `1`: Left button
    -   `2`: Middle button
    -   `3`: Right button

Example:
```
{
  type: "mouse_button_down",
  x: 100,
  y: 150,
  button: 1
}
```

### Keyboard Events

-   `type`: `"key_down"` or `"key_up"`
-   `key_name`: A string representing the name of the key (e.g., `"A"`, `"Space"`, `"Return"`).

Example:
```
{
  type: "key_down",
  key_name: "W"
}
```

## Example Usage

```pepper
loop [true] then
    dim event = wait_event[]

    if [event->type == "quit"] then
        break
    end

    if [event->type == "mouse_motion"] then
        println["Mouse moved to " + event->x + ", " + event->y]
    end

    if [event->type == "key_down"] then
        println[event->key_name + " key was pressed."]
    end
end
```
