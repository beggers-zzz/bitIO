// Ben Eggers
// GNU GPL'd

package bitIO

import "os"

// A struct useful for both bitReader and bitWriter
type bitIOStruct struct {
	Bits    []byte // should ALWAYS have length 1
	NumBits uint8
	File    *os.File
}

// Make a bitIOStruct on the passed File descriptor
func newStruct(file string) (b bitIOStruct, err error) {
	b.File, err = os.Open(file)
	if err != nil {
		return bitIOStruct{}, err
	}
	b.Bits = make([]byte, 1)
	b.NumBits = 0
	return b, nil
}
