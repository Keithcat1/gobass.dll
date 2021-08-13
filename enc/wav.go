package enc
import (
	"unsafe"
	"github.com/keithcat1/gobass"
)

/*
#include "bassenc.h"
#include "stdlib.h"
*/
import "C"
func NewWAVEncoderFile(channel bass.Channel, flags bass.Flags, file string) (Encoder, error) {
	cfile := C.CString(file)
	defer C.free(unsafe.Pointer(cfile))
	return Encoder(C.BASS_Encode_Start(C.DWORD(channel), cfile, EncodePCM, nil, nil)).ToError()
}