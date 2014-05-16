// Ben Eggers
// GNU GPL'd

// This package provides abstractions to read or write single bytes to and
// from disk. While a BitReader or BitWriter is doing stuff with a file,
// you shouldn't mess with it.
package bitIO

import (
	"errors"
)

type BitReader struct {
	bitIOStruct
}

// Set up and return a BitReader on the passed file.
func NewReader(file string) (b BitReader, err error) {
	str, err := newStruct(file)
	if err != nil {
		return BitReader{}, err
	}
	b = BitReader{str}
	// We need to initialize it on a byte!
	err = b.nextByte()
	return b, err
}

// Returns the next bit on the file stream. Will always be 0 or 1. Will
// return a non-nil err iff the read failed, or on EOF
func (b BitReader) ReadBit() (bit byte, err error) {
	bit = (b.Bits[0] & (1 << 7)) >> 7 // get the highest-order bit
	b.Bits[0] = b.Bits[0] * 2         // get rid of the highest-order bit
	b.NumBits++
	if b.NumBits == 8 {
		// we need the next byte!
		b.nextByte()
	}
	return bit, nil
}

// Closes the reader, closing its associated file descriptor
func (b BitReader) Close() (err error) {
	return b.File.Close()
}

func (b BitReader) nextByte() (err error) {
	n, err := b.File.Read(b.Bits)
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.New("Couldn't read from file")
	}
	b.NumBits = 0
	return nil
}
