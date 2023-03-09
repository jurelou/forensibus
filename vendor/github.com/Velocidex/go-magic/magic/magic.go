package magic

/*
 #cgo CFLAGS: -DHAVE_CONFIG_H
 #include <magic.h>
 #include <stdlib.h>
*/
import "C"
import (
	"sync"
	"unsafe"
)

const (
	MAGIC_NONE              = C.MAGIC_NONE
	MAGIC_DEBUG             = C.MAGIC_DEBUG
	MAGIC_SYMLINK           = C.MAGIC_SYMLINK
	MAGIC_COMPRESS          = C.MAGIC_COMPRESS
	MAGIC_DEVICES           = C.MAGIC_DEVICES
	MAGIC_MIME_TYPE         = C.MAGIC_MIME_TYPE
	MAGIC_CONTINUE          = C.MAGIC_CONTINUE
	MAGIC_CHECK             = C.MAGIC_CHECK
	MAGIC_PRESERVE_ATIME    = C.MAGIC_PRESERVE_ATIME
	MAGIC_RAW               = C.MAGIC_RAW
	MAGIC_ERROR             = C.MAGIC_ERROR
	MAGIC_MIME_ENCODING     = C.MAGIC_MIME_ENCODING
	MAGIC_MIME              = C.MAGIC_MIME
	MAGIC_EXTENSION         = C.MAGIC_EXTENSION
	MAGIC_APPLE             = C.MAGIC_APPLE
	MAGIC_NO_CHECK_COMPRESS = C.MAGIC_NO_CHECK_COMPRESS
	MAGIC_NO_CHECK_TAR      = C.MAGIC_NO_CHECK_TAR
	MAGIC_NO_CHECK_SOFT     = C.MAGIC_NO_CHECK_SOFT
	MAGIC_NO_CHECK_APPTYPE  = C.MAGIC_NO_CHECK_APPTYPE
	MAGIC_NO_CHECK_ELF      = C.MAGIC_NO_CHECK_ELF
	MAGIC_NO_CHECK_TEXT     = C.MAGIC_NO_CHECK_TEXT
	MAGIC_NO_CHECK_CDF      = C.MAGIC_NO_CHECK_CDF
	MAGIC_NO_CHECK_TOKENS   = C.MAGIC_NO_CHECK_TOKENS
	MAGIC_NO_CHECK_ENCODING = C.MAGIC_NO_CHECK_ENCODING
	MAGIC_NO_CHECK_ASCII    = C.MAGIC_NO_CHECK_ASCII
	MAGIC_NO_CHECK_FORTRAN  = C.MAGIC_NO_CHECK_FORTRAN
	MAGIC_NO_CHECK_TROFF    = C.MAGIC_NO_CHECK_TROFF
)
const (
	MAGIC_NO_CHECK_BUILTIN = MAGIC_NO_CHECK_COMPRESS |
		MAGIC_NO_CHECK_TAR |
		MAGIC_NO_CHECK_APPTYPE |
		MAGIC_NO_CHECK_ELF |
		MAGIC_NO_CHECK_TEXT |
		MAGIC_NO_CHECK_CDF |
		MAGIC_NO_CHECK_TOKENS |
		MAGIC_NO_CHECK_ENCODING
)

// A handle to manage the library.
type Magic struct {
	// Libmagic is not thread safe so we need to serialize access to
	// it.
	mu     sync.Mutex
	cookie C.magic_t
}

func NewMagicHandle(flags int) *Magic {
	cookie := C.magic_open(C.int(flags))
	result := &Magic{
		cookie: cookie,
	}
	return result
}

func (self *Magic) Close() {
	self.mu.Lock()
	defer self.mu.Unlock()

	C.magic_close((C.magic_t)(self.cookie))
}

func (self *Magic) GetError() string {
	self.mu.Lock()
	defer self.mu.Unlock()

	s := (C.magic_error((C.magic_t)(self.cookie)))
	ret := C.GoString(s)
	(C.magic_error_clear((C.magic_t)(self.cookie)))
	return ret
}

func (self *Magic) Errno() int {
	self.mu.Lock()
	defer self.mu.Unlock()

	return (int)(C.magic_errno((C.magic_t)(self.cookie)))
}

func (self *Magic) File(filename string) string {
	self.mu.Lock()
	defer self.mu.Unlock()

	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))
	return C.GoString(C.magic_file((C.magic_t)(self.cookie), cfilename))
}

func (self *Magic) Buffer(b []byte) string {
	self.mu.Lock()
	defer self.mu.Unlock()

	length := C.size_t(len(b))
	return C.GoString(C.magic_buffer(
		(C.magic_t)(self.cookie), unsafe.Pointer(&b[0]), length))
}

func (self *Magic) SetFlags(flags int) int {
	self.mu.Lock()
	defer self.mu.Unlock()

	return (int)(C.magic_setflags((C.magic_t)(self.cookie), C.int(flags)))
}

func (self *Magic) Check(filename string) int {
	self.mu.Lock()
	defer self.mu.Unlock()

	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))
	return (int)(C.magic_check((C.magic_t)(self.cookie), cfilename))
}

func (self *Magic) Compile(filename string) int {
	self.mu.Lock()
	defer self.mu.Unlock()

	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))
	return (int)(C.magic_compile((C.magic_t)(self.cookie), cfilename))
}

func (self *Magic) Load(filename string) int {
	self.mu.Lock()
	defer self.mu.Unlock()

	if filename == "" {
		return (int)(C.magic_load((C.magic_t)(self.cookie), nil))
	}
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))
	return (int)(C.magic_load((C.magic_t)(self.cookie), cfilename))
}

// Load definitions from this buffer.
func (self *Magic) LoadBuffer(data string) int {
	self.mu.Lock()
	defer self.mu.Unlock()

	cdata := C.CString(data)
	defer C.free(unsafe.Pointer(cdata))
	return (int)(C.magic_load_from_buffer((C.magic_t)(self.cookie), cdata))
}
