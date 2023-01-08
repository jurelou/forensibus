// +build windows

package prefetch

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	ntoskrnl              = windows.NewLazySystemDLL("ntdll.dll")
	RtlDecompressBufferEx = ntoskrnl.NewProc("RtlDecompressBufferEx")

	COMPRESSION_FORMAT_XPRESS_HUFF = uint16(0x0004)
)

func LZXpressHuffmanDecompressWithFallback(input []byte, output_size int) ([]byte, error) {

	// For older windows, we fall back to the build in decompression.
	err := RtlDecompressBufferEx.Find()
	if err != nil {
		return LZXpressHuffmanDecompress(input, output_size)
	}

	result := make([]byte, output_size)
	final_size := uint32(output_size)

	workspace := make([]byte, output_size*2)

	ret, _, _ := RtlDecompressBufferEx.Call(
		uintptr(COMPRESSION_FORMAT_XPRESS_HUFF),
		uintptr(unsafe.Pointer(&result[0])),
		uintptr(output_size),
		uintptr(unsafe.Pointer(&input[0])),
		uintptr(len(input)),
		uintptr(unsafe.Pointer(&final_size)),
		uintptr(unsafe.Pointer(&workspace[0])))
	if ret != 0 {
		// Fallback to the built in implementation on error.
		return LZXpressHuffmanDecompress(input, output_size)
	}

	// Return the decompressed buffer from the API.
	return result[:final_size], nil
}
