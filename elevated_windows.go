// +build windows

package elevated

import (
	"syscall"
	"unsafe"
)

// TokenElevated translated from winnt.h.
type TokenElevated struct {
	TokenIsElevated uint32
}

func IsElevated() bool {
	var token syscall.Token
	var handle syscall.Handle

	handle, err := syscall.GetCurrentProcess()
	if err != nil {
		panic(err)
	}

	err = syscall.OpenProcessToken(handle, syscall.TOKEN_QUERY, &token)
	if err != nil {
		panic(err)
	}
	defer token.Close()

	var bytesRequired uint32
	var tokenElevated TokenElevated
	tokenElevatedSize := uint32(unsafe.Sizeof(tokenElevated))
	err = syscall.GetTokenInformation(token,
		syscall.TokenElevation,
		(*byte)(unsafe.Pointer(&tokenElevated)),
		tokenElevatedSize,
		&bytesRequired)
	if err != nil {
		panic(err)
	}

	if tokenElevated.TokenIsElevated == 1 {
		return true
	}

	return false
}