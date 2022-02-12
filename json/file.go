package json

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func (f *jsonFile) exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

func (f *jsonFile) GetPath() string {
	return f.path
}

func (f *jsonFile) GetData() *[]byte {
	return f.data
}

func (f *jsonFile) ValidatePath() error {
	if !strings.HasSuffix(f.path, ".json") {
		return errors.New("invalid path, file must be a .json file")
	}
	if !f.exists(f.path) {
		return fmt.Errorf("invalid path, file does not exist: %s", f.path)
	}
	return nil
}

func (f *jsonFile) Read() (*[]byte, error) {
	err := f.ValidatePath()
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(f.path)
	if err != nil {
		return nil, err
	}

	f.data = &data

	return f.data, nil
}
