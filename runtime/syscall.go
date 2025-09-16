package runtime

func doSyscall(v VM, code int64) {
	switch {
	case code >= 0 && code <= 1:
		doSyscallIO(v, code)
	case code >= 2 && code <= 8:
		doSyscallGfx(v, code)
	case code >= 9 && code <= 11:
		doSyscallString(v, code)
	case code >= 12 && code <= 14:
		doSyscallHttp(v, code)
	default:
		panic("Unknown syscall code")
	}
}