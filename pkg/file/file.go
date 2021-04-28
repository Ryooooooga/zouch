package file

import (
	"fmt"
	"os"
)

func IsFile(filename string) (bool, error) {
	if stat, err := os.Stat(filename); os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else if stat.IsDir() {
		return false, fmt.Errorf("%s is a directory", filename)
	} else {
		return true, nil
	}
}
