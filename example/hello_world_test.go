package example

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Sebelum unit Test")

	m.Run() // mengeksekusi semua unit test

	fmt.Println("Sesudah unit test")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Akbar")
	if result != "Hello Akbar" {
		//unit test failed
		t.Error("Result must be Hello Akbar") // Menggagalkan unit test tapi masih berlanjut untuk unit test berikutnya
		// panic("Test ERROR: Result isn't Hello Akbar")
	}
	fmt.Println("TestHelloWorld with If Else Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Akbar")
	// assert akan memanggil fail yang berarti eksekusi unit test akan tetap dilanjutkan
	assert.Equal(t, "Hello Akbar", result)
	fmt.Println("TestHelloWorld with Assert Done")
}
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Akbar")
	// require akan memanggil FailNow yang berarti eksekusi unit test tidak akan dilanjutkan
	require.Equal(t, "Hello Akbar", result)
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Hello World - Akbar",
			request:  "Akbar",
			expected: "Hello Akbar",
		},
		{
			name:     "Hello World - Akbar - Negative",
			request:  "Akbar",
			expected: "Hello Akbir",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}
