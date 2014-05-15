// Ben Eggers
// GNU GPL'd

package bitIO

// Tests the BitReader

import (
	"testing"
	"io/ioutil"
	"os"
)

// globals
var filename = ".test"

// NewReader tests

func TestNewReader(t *testing.T) {
	err := makeBasicFile()
	if err != nil {
		t.Error("Couldn't create file: ", err)
	}
	b, err := NewReader(filename)
	b.Close()
}

func TestNewReaderReturnsErrorOnNonExistentFile(t *testing.T) {

}

// ReadBit tests

func TestBasicFileRead(t *testing.T) {
	// do nothing
}

// Makes a basic file to be used by the test suite. File bytes will be
// ascending powers of two.
func makeBasicFile() (err error) {
	bytes := make([]byte, 10)
	for i := 0; i < len(bytes); i++ {
		bytes[i] = 1 << uint8(i)
 	}
 	err = ioutil.WriteFile(filename, bytes, 0644)
 	defer os.Remove(filename)
 	return err
}