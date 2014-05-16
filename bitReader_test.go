// Ben Eggers
// GNU GPL'd

package bitIO

// Tests the BitReader

import (
	"io/ioutil"
	"os"
	"testing"
)

// globals
var filename = ".test"

////////////////////////////////////////////////////////////////////////////////
// NewReader tests
////////////////////////////////////////////////////////////////////////////////

func TestNewReader(t *testing.T) {
	err := makeBasicFile()
	defer os.Remove(filename)
	if err != nil {
		t.Error(err)
	}
	b, err := NewReader(filename)
	b.Close()
}

func TestNewReaderReturnsErrorOnNonExistentFile(t *testing.T) {
	_, err := NewReader(filename)
	if err == nil {
		t.Error("err should have been non-nil but was nil")
	}
}

////////////////////////////////////////////////////////////////////////////////
// ReadBit tests
////////////////////////////////////////////////////////////////////////////////

func TestBasicFileRead(t *testing.T) {
	err := makeBasicFile()
	//defer os.Remove(filename)
	if err != nil {
		t.Error(err)
	}

	br, err := NewReader(filename)
	if err != nil {
		t.Error(err)
	}

	// We should get 7 0's
	for i := 0; i < 8; i++ {
		bit, err := br.ReadBit()
		if err != nil {
			t.Error(err)
		}
		if bit != 0 {
			t.Error("Got incorrect bit. Should be 0, got", bit, "for bit", i)
		}
	}

	// Then a 1
	bit, err := br.ReadBit()
	if err != nil {
		t.Error(err)
	}
	if bit != 1 {
		t.Error("Got incorrect bit. Should be 1, got", bit)
	}
}

////////////////////////////////////////////////////////////////////////////////
// Helper functions
////////////////////////////////////////////////////////////////////////////////

// Makes a basic file to be used by the test suite. File bytes will be
// ascending powers of two.
func makeBasicFile() (err error) {
	bytes := make([]byte, 9)
	for i := 0; i < len(bytes)-1; i++ {
		bytes[i] = 1 << uint(i)
	}
	err = ioutil.WriteFile(filename, bytes, 0644)
	return err
}
