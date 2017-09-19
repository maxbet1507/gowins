// Code generated to call native win32. DO NOT EDIT.

package w32api

type HANDLE uintptr
type BOOL int32
type DWORD uint32
type UCHAR uint8
type USHORT uint16

type LOGFONT struct {
	Height         int32
	Width          int32
	Escapement     int32
	Orientation    int32
	Weight         int32
	Italic         byte
	Underline      byte
	StrikeOut      byte
	CharSet        byte
	OutPrecision   byte
	ClipPrecision  byte
	Quality        byte
	PitchAndFamily byte
	FaceName       [LF_FACESIZE]uint16
}

type NONCLIENTMETRICS struct {
	Size            uint32
	BorderWidth     int32
	ScrollWidth     int32
	ScrollHeight    int32
	CaptionWidth    int32
	CaptionHeight   int32
	CaptionFont     LOGFONT
	SmCaptionWidth  int32
	SmCaptionHeight int32
	SmCaptionFont   LOGFONT
	MenuWidth       int32
	MenuHeight      int32
	MenuFont        LOGFONT
	StatusFont      LOGFONT
	MessageFont     LOGFONT
}

type PSID uintptr

type ACL struct {
	AclRevision UCHAR
	Sbz1        UCHAR
	AclSize     USHORT
	AceCount    USHORT
	Sbz2        USHORT
}
type PACL uintptr

type SECURITY_DESCRIPTOR_CONTROL uint16

type SECURITY_DESCRIPTOR struct {
	Revision UCHAR
	Sbz1     UCHAR
	Control  SECURITY_DESCRIPTOR_CONTROL
	Owner    PSID
	Group    PSID
	Sacl     PACL
	Dacl     PACL
}
type PSECURITY_DESCRIPTOR uintptr

type SECURITY_ATTRIBUTES struct {
	Length             DWORD
	SecurityDescriptor *SECURITY_DESCRIPTOR
	InheritHandle      BOOL
}
