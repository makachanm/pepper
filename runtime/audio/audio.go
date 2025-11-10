package audio

import (
	"errors"
	"sync"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

type Sound struct {
	Chunk *mix.Chunk
}

type PepperAudio struct {
	sounds      map[int]*Sound
	music       *mix.Music
	nextSoundID int
	audioM      sync.Mutex
}

func NewPepperAudio() *PepperAudio {
	if err := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		panic(err)
	}
	return &PepperAudio{
		sounds:      make(map[int]*Sound),
		nextSoundID: 0,
	}
}

func (pa *PepperAudio) LoadSound(filename string) (int, error) {
	pa.audioM.Lock()
	defer pa.audioM.Unlock()
	chunk, err := mix.LoadWAV(filename)
	if err != nil {
		return -1, err
	}
	id := pa.nextSoundID
	pa.sounds[id] = &Sound{Chunk: chunk}
	pa.nextSoundID++
	return id, nil
}

func (pa *PepperAudio) PlaySound(id int, loops int) (int, error) {
	pa.audioM.Lock()
	defer pa.audioM.Unlock()
	sound, ok := pa.sounds[id]
	if !ok {
		return -1, errors.New("invalid sound ID")
	}
	return sound.Chunk.Play(-1, loops)
}

func (pa *PepperAudio) HaltChannel(channel int) {
	pa.audioM.Lock()
	defer pa.audioM.Unlock()
	mix.HaltChannel(channel)
}

func (pa *PepperAudio) DestroySound(id int) {
	pa.audioM.Lock()
	defer pa.audioM.Unlock()
	sound, ok := pa.sounds[id]
	if ok {
		sound.Chunk.Free()
		delete(pa.sounds, id)
	}
}

func (pa *PepperAudio) LoadMusic(file string) error {
	pa.audioM.Lock()
	defer pa.audioM.Unlock()
	var err error
	if pa.music != nil {
		pa.music.Free()
	}
	pa.music, err = mix.LoadMUS(file)
	if err != nil {
		return err
	}
	return nil
}

func (pa *PepperAudio) PlayMusic(loops int) error {
	pa.audioM.Lock()
	defer pa.audioM.Unlock()
	if pa.music == nil {
		return errors.New("no music loaded")
	}
	return pa.music.Play(loops)
}

func (pa *PepperAudio) HaltMusic() {
	pa.audioM.Lock()
	defer pa.audioM.Unlock()
	mix.HaltMusic()
}

func (pa *PepperAudio) SetVolume(channel int, volume int) {
	pa.audioM.Lock()
	defer pa.audioM.Unlock()
	mix.Volume(channel, volume)
}

func (pa *PepperAudio) SetMusicVolume(volume int) {
	pa.audioM.Lock()
	defer pa.audioM.Unlock()
	mix.VolumeMusic(volume)
}

func (pa *PepperAudio) PauseChannel(channel int) {
	pa.audioM.Lock()
	defer pa.audioM.Unlock()
	mix.Pause(channel)
}

func (pa *PepperAudio) ResumeChannel(channel int) {
	pa.audioM.Lock()
	defer pa.audioM.Unlock()
	mix.Resume(channel)
}

func (pa *PepperAudio) Close() {
	pa.audioM.Lock()
	defer pa.audioM.Unlock()
	for _, sound := range pa.sounds {
		sound.Chunk.Free()
	}
	if pa.music != nil {
		pa.music.Free()
	}
	pa.sounds = make(map[int]*Sound)
	mix.CloseAudio()
	sdl.Quit()
}
