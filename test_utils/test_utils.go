package test_utils

import (
	"fmt"
	"os"
	"testing"
)

type TestUtils struct {
}

func NewTestUtils() *TestUtils {
	return &TestUtils{}
}

func (this *TestUtils) CreateFolder(name string, t *testing.T) {
	if err := os.Mkdir(name, 0777); err != nil {
		t.Fatalf("Can't create %q folder, %v", name, err)
	}
}

func (this *TestUtils) CreateFiles(filePathFormat string, countFiles int, t *testing.T){
	for i := 1; i<=countFiles; i++ {
		file_name := fmt.Sprintf(filePathFormat, i)
		f, err := os.Create(file_name);
		if err != nil{
			t.Fatalf("Can't create test js file %s", file_name)
		}
		defer f.Close()
		
	}
}

func (this *TestUtils) RemoveAll(path string){
	if _, err := os.Stat(path); err == nil {
		defer os.RemoveAll(path)
	}
}
