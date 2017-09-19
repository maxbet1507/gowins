package SingleInstanceApplication

import (
	"unsafe"

	"github.com/maxbet1507/gowins/w32api"
)

func Check(name string) (error, func()) {
	sd := w32api.SECURITY_DESCRIPTOR{}
	if r, err := w32api.InitializeSecurityDescriptor(&sd, w32api.SECURITY_DESCRIPTOR_REVISION); !r {
		return err, func() {}
	}
	if r, err := w32api.SetSecurityDescriptorDacl(&sd, true, nil, false); !r {
		return err, func() {}
	}

	sa := w32api.SECURITY_ATTRIBUTES{}
	sa.Length = w32api.DWORD(unsafe.Sizeof(sa))
	sa.SecurityDescriptor = &sd
	sa.InheritHandle = w32api.TRUE

	handle, err := w32api.CreateMutex(&sa, false, name)
	if handle == uintptr(0) {
		return err, func() {}
	}

	if w32api.ErrorCode(err) == w32api.ERROR_SUCCESS {
		err = nil
	}

	return err, func() {
		w32api.CloseHandle(handle)
	}
}
