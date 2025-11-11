package runtime

func doSyscall(v *VM, code int64) {
	switch {
	case code >= 0 && code < 100:
		doSyscallIO(v, code)
	case code >= 100 && code < 200:
		doSyscallMath(v, code)
	case code >= 200 && code < 300:
		if code == 203 || code == 204 { // Route old JSON codes to the new handler
			doSyscallJson(v, code)
		} else {
			doSyscallString(v, code)
		}
	case code >= 300 && code < 400:
		doSyscallGfx(v, code)
	case code >= 400 && code < 500:
		doSyscallHttp(v, code)
	case code >= 500 && code < 600:
		doSyscallJson(v, code)
	case code >= 600 && code < 700:
		doSyscallTime(v, code)
	case code >= 700 && code < 800:
		doSyscallAudio(v, code)
	default:
		panic("Unknown or unimplemented syscall code")
	}
}
