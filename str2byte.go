package box

import (
	"reflect"
	"unsafe"
)

// S2bNew converts string to a byte slice without memory allocation.
// need g1.20 or more
func S2bNew(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// S2bOld converts string to a byte slice without memory allocation.
//
// Note it may break if string and/or slice header will change
// in the future go versions.
func S2bOld(s string) (b []byte) {
	/* #nosec G103 */
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	/* #nosec G103 */
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return b
}
