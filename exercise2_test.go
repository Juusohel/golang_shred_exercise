package exercise2

import (
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestShred(t *testing.T) {
	file, err := os.Create("test.txt")
	if err != nil {
		t.Fatalf("Unable to create test file")
	}

	// Creating random data
	data := make([]byte, 100)
	// new seed
	rand.Seed(time.Now().UnixNano())
	_, err = rand.Read(data)
	if err != nil {
		t.Fatalf("Unable to fill test file")
	}
	// write to test file
	_, err = file.WriteAt(data, 0)
	if err != nil {
		t.Fatalf("Unable to write test file")
	}

	err = file.Close()
	if err != nil {
		t.Fatalf("unable close file")
	}

	// Testing deletion of file
	err = Shred("test.txt")
	if err != nil {
		t.Fatalf("Shred failed %s", err)
	}

}
