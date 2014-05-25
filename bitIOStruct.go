// Ben Eggers
// GNU GPL'd

package bitIO

import "os"

// A struct useful for both bitReader and bitWriter
type bitIOStruct struct {
	bits    []byte // buffer, should ALWAYS have length 1
	numBits uint8
	file    *os.File
}

// Make a bitIOStruct on the passed File descriptor
func newStruct() (b bitIOStruct) {
	b.bits = make([]byte, 1)
	b.numBits = 0
	return b
}
