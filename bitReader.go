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
	str := newStruct()
	// Now open the file for reading
	str.file, err = os.Open(file)
	if err != nil {
		return BitReader{}, err
	}
	// This will make us grab the first byte on the first read
	str.numBits = 8
	return BitReader{str}, err
}

// Creates a new BitReader on the passed file descriptor, instead of from a
// file name like NewReader. This allows for reading certain parts of a file
// by bit, and parts the normal way.
func NewReaderOnFile(file *os.File) (b BitReader, err error) {
	str := newStruct()
	str.file = file
	str.numBits = 8
	return BitReader{str}, nil
}

// Returns the next bit on the file stream. Will always be 0 or 1. Will
// return a non-nil err iff the read failed, or on EOF
func (b *BitReader) ReadBit() (bit byte, err error) {
	if b.numBits == 8 {
		// we need the next byte!
		err = b.nextByte()
	}
	bit = (b.bits[0] & (1 << 7)) >> 7 // get the highest-order bit
	b.bits[0] = b.bits[0] * 2         // get rid of the highest-order bit
	b.numBits++
	return bit, err
}

// Closes the reader, closing its associated file descriptor
func (b *BitReader) Close() (err error) {
	return b.file.Close()
}

func (b *BitReader) nextByte() (err error) {
	n, err := b.file.Read(b.bits)
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.New("Couldn't read from file")
	}
	b.numBits = 0
	return nil
}

func (b *BitReader) CloseAndReturnFile() (f *os.File, err error) {
	file := b.file
	b.file = nil
	return file, nil
}