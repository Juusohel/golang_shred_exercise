package exercise2

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Shred overwrites the given file 3 times with random data before deleting it.
func Shred(path string) error {
	// Checking if file exists and opening with write flag
	file, err := os.OpenFile(path, os.O_WRONLY, 0)
	if err != nil {
		return fmt.Errorf("error shredding file: %s", err)
	}

	// Checking file size to overwrite all the data
	stats, err := file.Stat()
	if err != nil {
		return err
	}
	fileSize := stats.Size()

	// Overwriting file 3 times with new generated random data each time
	for i := 0; i < 3; i++ {
		// Creating random data
		data := make([]byte, fileSize)
		// new seed
		rand.Seed(time.Now().UnixNano())
		_, err := rand.Read(data)
		if err != nil {
			return err
		}
		// write to file
		_, err = file.WriteAt(data, 0)
		if err != nil {
			return err
		}
	}
	file.Close()

	// Deleting file
	err = os.Remove(path)
	if err != nil {
		return err
	}
	return err
}
