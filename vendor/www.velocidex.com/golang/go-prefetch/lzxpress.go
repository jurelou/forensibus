package prefetch

// Decompression algorithm reimplemented from Microsoft Reference here:
// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-frs2/8cb5bae9-edf3-4833-9f0a-9d7e24218d3d

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"sort"
)

type BitStream struct {
	// The input buffer.
	source []byte

	// The current byte offset into the source. As bits are drawn
	// from the stream we backfill them at this position.
	index int
	mask  uint32
	bits  uint32
}

func (self *BitStream) Lookup(n uint32) uint32 {
	if n == 0 {
		return 0
	}

	return self.mask >> (32 - n)
}

func (self *BitStream) Skip(n uint32) error {
	self.mask <<= n
	self.bits -= n
	if self.bits < 16 {
		if int(self.index+2) > len(self.source) {
			return io.EOF
		}
		self.mask += uint32(binary.LittleEndian.Uint16(
			self.source[self.index:self.index+2])) << (16 - self.bits)
		self.index += 2
		self.bits += 16
	}

	return nil
}

func NewBitStream(in []byte, in_pos int) *BitStream {
	return &BitStream{
		source: in,
		mask: uint32(binary.LittleEndian.Uint16(in[in_pos:]))<<16 +
			uint32(binary.LittleEndian.Uint16(in[in_pos+2:])),
		index: in_pos + 4,
		bits:  32,
	}
}

type PREFIX_CODE_NODE struct {
	id     uint32
	symbol uint32
	leaf   bool
	child  [2]*PREFIX_CODE_NODE
}

func (self PREFIX_CODE_NODE) String() string {
	return fmt.Sprintf("Node %d: symbol %v leaf %v\n",
		self.id, self.symbol, self.leaf)
}

type PREFIX_CODE_SYMBOL struct {
	id     uint32
	symbol uint32
	length uint32
}

func (self PREFIX_CODE_SYMBOL) String() string {
	return fmt.Sprintf("Symbol %d: symbol %v length %v\n",
		self.id, self.symbol, self.length)
}

/*
inout PREFIX_CODE_NODE treeNodes[1024]: A 1024 element

	PREFIX_CODE_NODE array that contains the Huffman prefix code
	tree's nodes.

in ULONG leafIndex: The index in treeNodes of the node to link into the tree.

in ULONG mask: The symbol's prefix code.

in ULONG bits: The number of bits in the symbol's prefix code.

# Return Value

Returns the index in treeNodes of the next node to be processed.
*/
func PrefixCodeTreeAddLeaf(
	treeNodes []PREFIX_CODE_NODE,
	leafIndex uint32,
	mask uint32,
	bits uint32) uint32 {

	node := &treeNodes[0]
	i := leafIndex + 1
	childIndex := uint32(0)

	for bits > 1 {
		bits = bits - 1
		childIndex = (mask >> bits) & 1
		if node.child[childIndex] == nil {
			node.child[childIndex] = &treeNodes[i]
			treeNodes[i].leaf = false
			i++
		}
		node = node.child[childIndex]
	}

	node.child[mask&1] = &treeNodes[leafIndex]

	return i
}

func PrefixCodeTreeRebuild(input []byte) *PREFIX_CODE_NODE {
	treeNodes := make([]PREFIX_CODE_NODE, 1024)
	symbolInfo := make([]PREFIX_CODE_SYMBOL, 512)

	for i := 0; i < 256; i++ {
		value := input[i]

		symbolInfo[2*i].id = uint32(2 * i)
		symbolInfo[2*i].symbol = uint32(2 * i)
		symbolInfo[2*i].length = uint32(value & 0xf)

		value >>= 4

		symbolInfo[2*i+1].id = uint32(2*i + 1)
		symbolInfo[2*i+1].symbol = uint32(2*i + 1)
		symbolInfo[2*i+1].length = uint32(value & 0xf)
	}

	sort.SliceStable(symbolInfo, func(i, j int) bool {
		a := symbolInfo[i]
		b := symbolInfo[j]
		if a.length < b.length {
			return true
		}

		if a.symbol < b.symbol {
			return true
		}

		return false
	})

	i := 0
	for i < 512 && symbolInfo[i].length == 0 {
		i++
	}

	mask := uint32(0)
	bits := uint32(1)

	root := &treeNodes[0]
	root.leaf = false

	j := uint32(1)
	for ; i < 512; i++ {
		treeNodes[j].id = uint32(j)
		treeNodes[j].symbol = symbolInfo[i].symbol
		treeNodes[j].leaf = true
		mask = mask << (symbolInfo[i].length - bits)
		bits = symbolInfo[i].length
		j = PrefixCodeTreeAddLeaf(treeNodes, j, mask, bits)
		mask++
	}

	return root
}

func PrefixCodeTreeDecodeSymbol(bstr *BitStream, root *PREFIX_CODE_NODE) (
	uint32, error) {
	node := root

	for {
		bit := bstr.Lookup(1)
		err := bstr.Skip(1)
		if err != nil {
			return 0, err
		}

		node = node.child[bit]
		if node == nil {
			return 0, errors.New("Corruption detected")
		}

		if node.leaf {
			break
		}
	}

	return node.symbol, nil
}

func LZXpressHuffmanDecompressChunk(
	in_idx int, // Current cursor in the input buffer.
	input []byte, // Input buffer.
	out_idx int, // Cursor cursor on the output buffer.
	output []byte, // The output buffer.
	chunk_size int, // The required size of uncompressed buffer
) (int, int, error) {

	// There must be at least this many bytes available to read.
	if in_idx+256 > len(input) {
		return 0, 0, io.EOF
	}

	root := PrefixCodeTreeRebuild(input[in_idx:])
	bstr := NewBitStream(input, in_idx+256)

	i := out_idx
	for i < out_idx+chunk_size {
		symbol, err := PrefixCodeTreeDecodeSymbol(bstr, root)
		if err != nil {
			if err == io.EOF {
				err = nil
			}

			return int(bstr.index), i, err
		}
		if symbol < 256 {
			output[i] = byte(symbol)
			i++

		} else {
			symbol -= 256
			length := uint32(symbol & 15)
			symbol >>= 4

			offset := int32(0)
			if symbol != 0 {
				offset = int32(bstr.Lookup(symbol))
			}
			offset |= 1 << symbol
			offset = -offset

			if length == 15 {
				length = uint32(bstr.source[bstr.index]) + 15
				bstr.index = bstr.index + 1

				if length == 270 {
					length = uint32(binary.LittleEndian.Uint16(
						bstr.source[bstr.index:]))
					bstr.index = bstr.index + 2
				}
			}

			err := bstr.Skip(symbol)
			if err != nil {
				return int(bstr.index), i, err
			}

			length = length + 3
			for {
				if i+int(offset) < 0 {
					return int(bstr.index),
						i, errors.New("Decompression error")
				}

				output[i] = output[i+int(offset)]
				i++
				length -= 1
				if length == 0 {
					break
				}
			}
		}

	}

	return int(bstr.index), i, nil
}

// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-frs2/8cb5bae9-edf3-4833-9f0a-9d7e24218d3d
func LZXpressHuffmanDecompress(input []byte, output_size int) ([]byte, error) {
	output := make([]byte, output_size)
	var err error

	// Index into the input buffer.
	in_idx := 0

	// Index into the output buffer.
	out_idx := 0

	for {
		// How much data belongs in the current chunk. Chunks
		// are split into maximum 65536 bytes.
		chunk_size := output_size - out_idx
		if chunk_size > 65536 {
			chunk_size = 65536
		}

		in_idx, out_idx, err = LZXpressHuffmanDecompressChunk(
			in_idx, input, out_idx, output, chunk_size)
		if err != nil {
			return output, err
		}

		// We are done.
		if out_idx >= len(output) || in_idx >= len(input) {
			break
		}
	}

	return output, nil
}
