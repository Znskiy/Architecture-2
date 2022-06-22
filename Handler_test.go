package lab2

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)


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
		assert.Equal(t, "error in input", err.Error())
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
		assert.Equal(t, "error in input", err.Error())
	} else {
		t.Errorf("Incorrect result")
	}
}
