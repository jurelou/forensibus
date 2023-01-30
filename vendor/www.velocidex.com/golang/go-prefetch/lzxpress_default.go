// +build !windows

package prefetch

// Non windows systems fall back to build in decompression.
func LZXpressHuffmanDecompressWithFallback(input []byte, output_size int) ([]byte, error) {
	return LZXpressHuffmanDecompress(input, output_size)
}
