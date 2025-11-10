package runtime

import (
	"fmt"
	"pepper/runtime/audio"
)

var Audio *audio.PepperAudio

func AudioNew() {
	Audio = audio.NewPepperAudio()
}

func AudioLoadSound(stack *OperandStack) {
	path := stack.Pop().Value.(string)
	id, err := Audio.LoadSound(path)
	if err != nil {
		fmt.Println("Error loading audio:", err)
		stack.Push(VMDataObject{Type: INTGER, Value: int64(-1)})
		return
	}
	stack.Push(VMDataObject{Type: INTGER, Value: int64(id)})
}

func AudioPlaySound(stack *OperandStack) {
	loops := int(stack.Pop().Value.(int64))
	id := int(stack.Pop().Value.(int64))
	channel, err := Audio.PlaySound(id, loops)
	if err != nil {
		fmt.Println("Error playing audio:", err)
		stack.Push(VMDataObject{Type: INTGER, Value: int64(-1)})
		return
	}
	stack.Push(VMDataObject{Type: INTGER, Value: int64(channel)})
}

func AudioHaltChannel(stack *OperandStack) {
	channel := int(stack.Pop().Value.(int64))
	Audio.HaltChannel(channel)
}

func AudioDestroySound(stack *OperandStack) {
	id := int(stack.Pop().Value.(int64))
	Audio.DestroySound(id)
}

func AudioLoadMusic(stack *OperandStack) {
	path := stack.Pop().Value.(string)
	err := Audio.LoadMusic(path)
	if err != nil {
		fmt.Println("Error loading music:", err)
		stack.Push(VMDataObject{Type: BOOLEAN, Value: false})
		return
	}
	stack.Push(VMDataObject{Type: BOOLEAN, Value: true})
}

func AudioPlayMusic(stack *OperandStack) {
	loops := int(stack.Pop().Value.(int64))
	err := Audio.PlayMusic(loops)
	if err != nil {
		fmt.Println("Error playing music:", err)
	}
}

func AudioHaltMusic(stack *OperandStack) {
	Audio.HaltMusic()
}

func AudioSetVolume(stack *OperandStack) {
	volume := int(stack.Pop().Value.(int64))
	channel := int(stack.Pop().Value.(int64))
	Audio.SetVolume(channel, volume)
}

func AudioSetMusicVolume(stack *OperandStack) {
	volume := int(stack.Pop().Value.(int64))
	Audio.SetMusicVolume(volume)
}

func AudioPauseChannel(stack *OperandStack) {
	channel := int(stack.Pop().Value.(int64))
	Audio.PauseChannel(channel)
}

func AudioResumeChannel(stack *OperandStack) {
	channel := int(stack.Pop().Value.(int64))
	Audio.ResumeChannel(channel)
}
