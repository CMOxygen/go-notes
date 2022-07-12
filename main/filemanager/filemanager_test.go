package filemanager

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	t.Log("testing read file")
	{
		r, err := ReadFile("/etc/server/c/r")
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("read result = %v", string(r))
	}
}
