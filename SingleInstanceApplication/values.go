package SingleInstanceApplication

import (
	"unsafe"

	"github.com/maxbet1507/gowins/w32api"
)

func Check(name string) (bool, func()) {
	sd := w32api.SECURITY_DESCRIPTOR{}
	if r, err := w32api.InitializeSecurityDescriptor(&sd, w32api.SECURITY_DESCRIPTOR_REVISION); !r {
		panic(err)
	}
	if r, err := w32api.SetSecurityDescriptorDacl(&sd, true, nil, false); !r {
		panic(err)
	}

	sa := w32api.SECURITY_ATTRIBUTES{}
	sa.Length = w32api.DWORD(unsafe.Sizeof(sa))
	sa.SecurityDescriptor = &sd
	sa.InheritHandle = w32api.TRUE

	handle, err := w32api.CreateMutex(&sa, false, name)
	if handle == uintptr(0) {
		panic(err)
	}

	closer := func() {
		w32api.CloseHandle(handle)
	}

	ret := w32api.ErrorCode(err) == w32api.ERROR_SUCCESS
	return ret, closer
}
