package unittest

import (
	"testing"
)

func TestHelloWorld(t *testing.T)  {
	result := HelloWorld("jono1")

	if result != "haloe jono" {
		t.Fail()
		// panic("kosong")
	}
}