package commands

import (
	"fmt"
	"io/ioutil"
	"os"
)

func (cmd *Command) List(files []string) error {
	fileEntries, err := ioutil.ReadDir(cmd.rootDir)
	if os.IsNotExist(err) {
		// Template directory does not exist
		// Nothing to do
		return nil
	} else if err != nil {
		return err
	}

	for _, fileEntry := range fileEntries {
		if !fileEntry.IsDir() {
			fmt.Println(fileEntry.Name())
		}
	}

	return nil
}
