package runtime

func doSyscallAudio(vm *VM, code int64) {
	switch code {
	case 700: // audio_load
		AudioLoadSound(vm.OperandStack)
	case 701: // audio_play
		AudioPlaySound(vm.OperandStack)
	case 702: // audio_halt
		AudioHaltChannel(vm.OperandStack)
	case 703: // audio_free
		AudioDestroySound(vm.OperandStack)
	case 704: // audio_load_music
		AudioLoadMusic(vm.OperandStack)
	case 705: // audio_play_music
		AudioPlayMusic(vm.OperandStack)
	case 706: // audio_halt_music
		AudioHaltMusic(vm.OperandStack)
	case 707: // audio_set_volume
		AudioSetVolume(vm.OperandStack)
	case 708: // audio_set_music_volume
		AudioSetMusicVolume(vm.OperandStack)
	case 709: // audio_pause
		AudioPauseChannel(vm.OperandStack)
	case 710: // audio_resume
		AudioResumeChannel(vm.OperandStack)
	}
}
