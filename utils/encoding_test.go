package utils_test

import (
	"testing"

	"blops-me/utils"
)

func TestBase62Encoding(t *testing.T) {
	encoded := utils.EncodeBase62("Hello, World!")
	t.Logf("Original: Hello, World!")
	t.Logf("Encoded: %s", encoded)

	decoded := utils.DecodeBase62(encoded)
	if decoded != "Hello, World!" {
		t.Errorf("Expected hello, got %s", decoded)
	}

	encoded = utils.EncodeBase62("안녕! 123")
	t.Logf("Original: 안녕! 123")
	t.Logf("Encoded: %s", encoded)

	decoded = utils.DecodeBase62(encoded)
	if decoded != "안녕! 123" {
		t.Errorf("Expected 안녕! 123, got %s", decoded)
	}
}
