package utils

import (
	"encoding/base64"
	"testing"
)

func TestEncodeStringBase64(t *testing.T) {
	tests := []struct {
		name            string
		stringToEncode  string
		stringExtpected string
	}{
		{
			name:            "Encode a valid string",
			stringToEncode:  "JustAMessageToEncode",
			stringExtpected: base64.StdEncoding.EncodeToString([]byte("JustAMessageToEncode")),
		},
		{
			name:            "Encode empty string",
			stringToEncode:  "",
			stringExtpected: base64.StdEncoding.EncodeToString([]byte("")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EncodeStringBase64(tt.stringToEncode)
			if result != tt.stringExtpected {
				t.Errorf("EncodeStringBase64(%s) = %s; want %s", tt.stringToEncode, result, tt.stringExtpected)
			}
		})
	}
}

func TestDecodeStringBase64(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  string
		expectErr bool
	}{
		{
			name:      "Decode a valid base64 string",
			input:     base64.StdEncoding.EncodeToString([]byte("JustAMessageToEncode")),
			expected:  "JustAMessageToEncode",
			expectErr: false,
		},
		{
			name:      "Decode empty string",
			input:     base64.StdEncoding.EncodeToString([]byte("")),
			expected:  "",
			expectErr: false,
		},
		{
			name:      "Invalid base64 string",
			input:     "InvalidBase64String",
			expected:  "",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := DecodeStringBase64(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("DecodeStringBase64() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if result != tt.expected {
				t.Errorf("DecodeStringBase64() = %s; want %s", result, tt.expected)
			}
		})
	}
}
