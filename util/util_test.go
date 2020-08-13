package util

import (
	"testing"
)

// Test case for the function Palindrome
func TestIsPalindrome(t *testing.T) {
	input, expected := "eeeeieeee", true
	result := IsPalindrome(input)
	if result != expected {
		t.Errorf("Result: %t, Expected: %t", result, expected)
	}

}
