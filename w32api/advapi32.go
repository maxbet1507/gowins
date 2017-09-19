// Code generated to call native win32. DO NOT EDIT.

package w32api

import (
	"syscall"
	"unsafe"
)

var (
	advapi32 = syscall.NewLazyDLL("advapi32.dll")

	procInitializeSecurityDescriptor = advapi32.NewProc("InitializeSecurityDescriptor")
	procSetSecurityDescriptorDacl    = advapi32.NewProc("SetSecurityDescriptorDacl")
)

func InitializeSecurityDescriptor(pSecurityDescriptor *SECURITY_DESCRIPTOR, dwRevision DWORD) (bool, error) {
	v1 := unsafe.Pointer(pSecurityDescriptor)

	ret, _, err := procInitializeSecurityDescriptor.Call(uintptr(v1), uintptr(dwRevision))
	return ret != 0, err
}

func SetSecurityDescriptorDacl(pSecurityDescriptor *SECURITY_DESCRIPTOR, bDaclPresent bool, pDacl *ACL, bDaclDefaulted bool) (bool, error) {
	v1 := unsafe.Pointer(pSecurityDescriptor)
	v2 := FALSE
	v3 := unsafe.Pointer(pDacl)
	v4 := FALSE

	if bDaclPresent {
		v2 = TRUE
	}

	if bDaclDefaulted {
		v4 = TRUE
	}

	ret, _, err := procSetSecurityDescriptorDacl.Call(uintptr(v1), uintptr(v2), uintptr(v3), uintptr(v4))
	return ret != 0, err
}
