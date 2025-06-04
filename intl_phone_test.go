package phone

import (
	"fmt"
	"testing"
)

func TestIntlPhone(t *testing.T) {
	testData := []struct {
		input    string
		expected []string
	}{
		{"Contact me at +1 555 123 4567 or email me", []string{"15551234567"}},
		{"Call UK support at +44 20 1234 5678", []string{"442012345678"}},
		{"German office: +49 30 12345678", []string{"493012345678"}},
		{"French contact: +33.1.23.45.67.89", []string{"33123456789"}},
		{"Japan branch: +81-3-1234-5678", []string{"81312345678"}},
		{"China office: +86 10 1234 5678 and India: +91 98765 43210", []string{"861012345678", "919876543210"}},
		{"Multiple numbers: +1 555 123 4567, 0987654321, and +44 20 1234 5678", []string{"15551234567", "0987654321", "442012345678"}},
		{"No valid phone numbers here", []string{}},
	}

	for _, tc := range testData {
		phones := GetPhone(tc.input)
		fmt.Printf("Input: %s\nDetected: %v\nExpected: %v\n\n", tc.input, phones, tc.expected)

		// Kiểm tra số lượng số điện thoại tìm được
		if len(phones) != len(tc.expected) {
			t.Errorf("Expected %d phones, got %d for input: %s", len(tc.expected), len(phones), tc.input)
			continue
		}

		// Kiểm tra từng số điện thoại
		for i, phone := range phones {
			if i >= len(tc.expected) || phone != tc.expected[i] {
				t.Errorf("Expected phone %s, got %s for input: %s", tc.expected[i], phone, tc.input)
			}
		}
	}
}

func TestIntlHiddenPhone(t *testing.T) {
	testData := []struct {
		input    string
		contains string
	}{
		{"Contact me at +1 555 123 4567 or email me", "Contact me at ************ or email me"},
		{"Call UK support at +44 20 1234 5678", "Call UK support at *************** or"},
		{"German office: +49 30 12345678", "German office: ************** or"},
		{"Multiple numbers: +1 555 123 4567, 0987654321, and +44 20 1234 5678", "Multiple numbers: ************, **********, and ***************"},
	}

	for _, tc := range testData {
		hidden := HiddenPhone(tc.input)
		fmt.Printf("Input: %s\nHidden: %s\n\n", tc.input, hidden)

		// Kiểm tra kết quả ẩn số điện thoại
		if hidden == tc.input {
			t.Errorf("Expected phone numbers to be hidden, but got the same string for input: %s", tc.input)
		}
	}
}
