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
	err := ioutil.WriteFile(filename, make([]byte, 1), 0644)
	defer os.Remove(filename)
	if err != nil {
		t.Error(err)
	}
	b, err := NewReader(filename)
	if err != nil {
		t.Error(err)
	}
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
	bytes := make([]byte, 1)
	bytes[0] = 255 // all 1s

	err := ioutil.WriteFile(filename, bytes, 0644)
	defer os.Remove(filename)
	if err != nil {
		t.Error(err)
	}

	br, err := NewReader(filename)
	if err != nil {
		t.Error(err)
	}

	// We should get all 1s
	for i := 0; i < 8; i++ {
		bit, err := br.ReadBit()
		if err != nil {
			t.Error(err)
		}

		if bit != 1 {
			t.Error("Wanted 1, got:", bit)
		}
	}
}

func TestFileReadOffEndOfFile(t *testing.T) {
	bytes := make([]byte, 1)
	err := ioutil.WriteFile(filename, bytes, 0644)
	defer os.Remove(filename)
	if err != nil {
		t.Error(err)
	}

	br, err := NewReader(filename)
	if err != nil {
		t.Error(err)
	}

	// get rid of those 8 bits
	for i := 0; i < 8; i++ {
		_, err := br.ReadBit()
		if err != nil {
			// Something went wrong before we even got to what we're testing
			t.Error(err)
		}
	}

	// Now we should get an error
	_, err = br.ReadBit()
	if err == nil {
		t.Error("Should have gotten EOF error, got nil instead")
	}
}

func TestMultiByteFileRead(t *testing.T) {
	bytes := make([]byte, 4)
	bytes[0] = 255 // all 1s

	err := ioutil.WriteFile(filename, bytes, 0644)
	defer os.Remove(filename)
	if err != nil {
		t.Error(err)
	}

	br, err := NewReader(filename)
	if err != nil {
		t.Error(err)
	}

	// We should get all 1s
	for i := 0; i < 8; i++ {
		bit, err := br.ReadBit()
		if err != nil {
			t.Error(err)
		}

		if bit != 1 {
			t.Error("Wanted 1, got:", bit)
		}
	}
}
