package lab2

import (
	"bytes"
	"github.com/stretchr/testify"
	"io/ioutil"
	"os"
	"testing"
)

func TestConsoleCorrectInput(t *testing.T) {
	stdout := os.Stdout
	read, write, _ := os.Pipe()
	os.Stdout = write
	handler := ComputeHandler{
		Input:  bytes.NewBufferString("2 2 +"),
		Output: write,
	}
	err := handler.Compute()
	write.Close()
	out, _ := ioutil.ReadAll(read)
	os.Stdout = stdout
	if assert.Nil(t, err) {
		assert.Equal(t, "2 + 2\n", string(out[:]))
	} else {
		t.Errorf("Incorrect result")
	}
}

func TestConsoleRandomChars(t *testing.T) {
	stdout := os.Stdout
	_, write, _ := os.Pipe()
	os.Stdout = write
	handler := ComputeHandler{
		Input:  bytes.NewBufferString("Banana guchi"),
		Output: write,
	}
	err := handler.Compute()
	write.Close()
	os.Stdout = stdout
	if assert.NotNil(t, err) {
		assert.Equal(t, "Convertation Error", err.Error())
	} else {
		t.Errorf("Incorrect result")
	}
}

func TestConsoleOnlyNumbers(t *testing.T) {
	stdout := os.Stdout
	_, write, _ := os.Pipe()
	os.Stdout = write
	handler := ComputeHandler{
		Input:  bytes.NewBufferString("3 5 "),
		Output: write,
	}
	err := handler.Compute()
	write.Close()
	os.Stdout = stdout
	if assert.NotNil(t, err) {
		assert.Equal(t, "Convertation Error", err.Error())
	} else {
		t.Errorf("Incorrect result")
	}
}
