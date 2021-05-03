package mock

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func UseIoWriterMock(t *testing.T, f func(writer *os.File)) io.Writer {
	file, err := ioutil.TempFile("/tmp", "output.*.txt")
	defer os.Remove(file.Name())
	defer file.Close()
	if err != nil {
		t.Skip("could not create mock for writer")
	}
	f(file)
	return file
}

func UseInvalidIoWriterMock(t *testing.T, f func(writer *os.File)) io.Writer {
	file, err := ioutil.TempFile("/tmp", "output.*.txt")
	file.Close()
	defer os.Remove(file.Name())
	if err != nil {
		t.Skip("could not create mock for writer")
	}
	f(file)
	return file
}