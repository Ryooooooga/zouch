package file

import (
	"io"
	"os"
	"time"
)

func UpdateTimestamp(destination string, time time.Time) error {
	accessed_time := time
	modified_time := time

	if err := os.Chtimes(destination, accessed_time, modified_time); err != nil {
		return err
	}

	return nil
}

func Copy(source string, destination string) error {
	input, err := os.Open(source)
	if err != nil {
		return err
	}
	defer input.Close()

	output, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer output.Close()

	if _, err := io.Copy(output, input); err != nil {
		return err
	}

	return nil
}
