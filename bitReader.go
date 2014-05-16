// Ben Eggers
// GNU GPL'd

// This package provides abstractions to read or write single bytes to and
// from disk. While a BitReader or BitWriter is doing stuff with a file,
// you shouldn't mess with it.
package bitIO

import (
	"errors"
	"fmt"
)

type BitReader struct {
	bitIOStruct
}

// Set up and return a BitReader on the passed file.
func NewReader(file string) (b BitReader, err error) {
	str, err := newStruct(file)
	return BitReader{str}, err
}

// Returns the next bit on the file stream. Will always be 0 or 1. Will
// return a non-nil err iff the read failed, or on EOF
func (b BitReader) ReadBit() (bit byte, err error) {
	fmt.Println("bits:", b.Bits[0], "numBits", b.NumBits)
	bit = b.Bits[0] & (1 << 7) >> 7 // get the highest-order bit
	b.Bits[0] = b.Bits[0] * 2
	b.NumBits = b.NumBits + 1
	if b.NumBits == 8 {
		// we need the next byte!
		n, err := b.File.Read(b.Bits)
		if err != nil {
			return 0, err
		}
		if n == 0 {
			return 0, errors.New("Couldn't read from file")
		}
		b.NumBits = 0
	}
	return bit, nil
}

// Closes the reader, closing its associated file descriptor
func (b BitReader) Close() (err error) {
	return b.File.Close()
}
