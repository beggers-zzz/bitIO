// Ben Eggers
// GNU GPL'd

// This package provides abstractions to read or write single bytes to and
// from disk. While a BitReader or BitWriter is doing stuff with a file,
// you shouldn't mess with it.
package bitIO

import (
	"errors"
	"os"
)

type BitReader struct {
	bitIOStruct
}

// Set up and return a BitReader on the passed file.
func NewReader(file string) (b BitReader, err error) {
	str, err := newStruct()
	if err != nil {
		return BitReader{}, err
	}
	// Now open the file for reading
	b = BitReader{str}
	b.File, err = os.Open(file)
	if err != nil {
		return BitReader{}, err
	}
	// This will make us grab the first byte on the first read
	b.NumBits = 8
	return b, err
}

// Returns the next bit on the file stream. Will always be 0 or 1. Will
// return a non-nil err iff the read failed, or on EOF
func (b *BitReader) ReadBit() (bit byte, err error) {
	if b.NumBits == 8 {
		// we need the next byte!
		err = b.nextByte()
	}
	bit = (b.Bits[0] & (1 << 7)) >> 7 // get the highest-order bit
	b.Bits[0] = b.Bits[0] * 2         // get rid of the highest-order bit
	b.NumBits++
	return bit, err
}

// Closes the reader, closing its associated file descriptor
func (b *BitReader) Close() (err error) {
	return b.File.Close()
}

func (b *BitReader) nextByte() (err error) {
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
