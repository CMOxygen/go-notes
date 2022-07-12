package filemanager

import "os"

func ReadFile(path string) (r []byte, err error) {

	r, err = os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return r, err
}
