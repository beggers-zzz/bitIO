// Ben Eggers
// GNU GPL'd

package bitIO

import (
	"errors"
	"os"
)

type BitWriter struct {
	bitIOStruct
}

// Set up and return a BitWriter on the passed file. Truncates the file if
// it already exists, be careful!
func NewWriter(file string) (b BitWriter, err error) {
	str := newStruct()

	// Now open the file for writing
	str.file, err = os.Create(file)
	return BitWriter{str}, err
}

// Creates a new BitWriter on the passed file descriptor instead of a filename
// like NewWriter. This allows for writing bits to only part of the file, and doing
// file IO the normal way on other parts.
func NewWriterOnFile(file *os.File) (b BitWriter, err error) {
	str := newStruct()
	str.file = file
	return BitWriter{str}, nil
}

// Writes one bit. If the passed int8 is 1, writes a one. If it's 0,
// writes a 0. Else, returns a non-nil error.
func (b *BitWriter) WriteBit(bit byte) (err error) {
	if bit != 0 && bit != 1 {
		return errors.New("Invalid bit to write.")
	}

	if b.numBits == 8 {
		err = b.flush()
		if err != nil {
			return err
		}
	}

	b.bits[0] += bit << (7 - b.numBits)
	b.numBits++
	return nil
}

// Flushes the current byte out to disk, padding with 0s if necessary.
func (b *BitWriter) flush() (err error) {
	for b.numBits < 8 {
		// Pad with 0s
		b.WriteBit(0)
	}
	_, err = b.file.Write(b.bits)
	b.numBits = 0
	b.bits[0] = 0
	return err
}

// Closes the BitReader, flushing final bits to disk if need be and closing
// the file descriptor.
func (b *BitWriter) Close() (err error) {
	err = b.flush()
	if err != nil {
		return err
	}
	return b.file.Close()
}

func (b *BitWriter) CloseAndReturnFile() (f *os.File, err error) {
	return b.file, b.flush()
}