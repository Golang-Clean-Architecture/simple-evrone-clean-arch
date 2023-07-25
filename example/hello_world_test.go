package example

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Akbar")
	if result != "Hello Akbar" {
		//unit test failed
		t.Fail() // Menggagalkan unit test tapi masih berlanjut untuk unit test berikutnya
		// panic("Test ERROR: Result isn't Hello Akbar")
	}
}
