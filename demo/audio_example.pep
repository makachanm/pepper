dim music_loaded = load_sound[`demo/res/untitled.ogg`]
dim channel = play_sound[music_loaded 3]

print[`sound playing in channel: `]
println[channel]

dim volume = 0
dim volume_direction = 5

loop [true] then
    volume = volume + volume_direction
    if [volume > 128 or volume < 0] then
        volume_direction = 0 - volume_direction
    end
    set_sfx_volume[channel volume]
    sleep[50]
end
