package testdata

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
)

func TestPath(t *testing.T) {
	_, currentFile, _, _ := runtime.Caller(0)
	fmt.Printf("currentFile : %v\n", currentFile)

	basepath = filepath.Dir(currentFile)
	fmt.Printf("bashpath : %v\n", basepath)

}
