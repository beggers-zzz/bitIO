// Ben Eggers
// GNU GPL'd

// Tests the BitWriter

package bitIO

import (
	"io/ioutil"
	"os"
	"testing"
)

////////////////////////////////////////////////////////////////////////////////
// NewWriter tests
////////////////////////////////////////////////////////////////////////////////

func TestNewWriter(t *testing.T) {
	bw, err := NewWriter(filename)
	defer os.Remove(filename) // we don't want it to stick around after the test
	if err != nil {
		t.Error(err)
	}
	_, err = os.Open(filename) // make sure the file is there
	if err != nil {
		t.Error(err)
	}
	err = bw.Close()
	if err != nil {
		t.Error(err)
	}
}

////////////////////////////////////////////////////////////////////////////////
// WriteBit tests
////////////////////////////////////////////////////////////////////////////////

func TestWriteBasic(t *testing.T) {
	bw, err := NewWriter(filename)
	defer os.Remove(filename)
	if err != nil {
		t.Error(err)
	}

	// Let's write some bits
	for i := 0; i < 8; i++ {
		bw.WriteBit(1)
	}

	// And close to make sure they're written out
	bw.Close()

	// Now let's check the file is there	 there
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		// Something went badly on opening? Weird
		t.Error(err)
	}

	if len(b) != 1 {
		// We read more bytes than we should have, error
		t.Error("Got more bytes than we should have:", len(b))
	}

	if b[0] != 255 {
		t.Error("Should have gotten all 1's (255), got", b[0])
	}
}

func TestWritePadsWithZeroes(t *testing.T) {
	bw, err := NewWriter(filename)
	defer os.Remove(filename)
	if err != nil {
		t.Error(err)
	}

	// Let's write one bit
	bw.WriteBit(1)

	// And close to make sure it's written out. It should be padded with 0's
	bw.Close()

	// Now let's check the file is there	 there
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		// Something went badly on opening? Weird
		t.Error(err)
	}

	if len(b) != 1 {
		// We read more bytes than we should have, error
		t.Error("Got more bytes than we should have:", len(b))
	}

	if b[0] != 128 {
		t.Error("Should have gotten all 0b1000000 (128), got", b[0])
	}
}