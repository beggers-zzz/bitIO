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
	str.File, err = os.Create(file)
	return BitWriter{str}, err
}

// Writes one bit. If the passed int8 is 1, writes a one. If it's 0,
// writes a 0. Else, returns a non-nil error.
func (b *BitWriter) WriteBit(bit byte) (err error) {
	if bit != 0 && bit != 1 {
		return errors.New("Invalid bit to write.")
	}

	if b.NumBits == 8 {
		err = b.flush()
		if err != nil {
			return err
		}
	}

	b.Bits[0] += bit << (7 - b.NumBits)
	b.NumBits++
	return nil
}

// Flushes the current byte out to disk, padding with 0s if necessary.
func (b *BitWriter) flush() (err error) {
	for b.NumBits != 8 {
		b.WriteBit(0)
	}
	_, err = b.File.Write(b.Bits)
	return err
}

// Closes the BitReader, flushing final bits to disk if need be and closing
// the file descriptor.
func (b *BitWriter) Close() (err error) {
	err = b.flush()
	if err != nil {
		return err
	}
	return b.File.Close()
}
